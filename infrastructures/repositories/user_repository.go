package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/hiroyky/famiphoto/drivers/mysql"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/dbmodels"
	"github.com/hiroyky/famiphoto/utils/cast"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"time"
)

type UserRepository interface {
	GetUser(ctx context.Context, userID string) (*dbmodels.User, error)
	GetUsers(ctx context.Context, filter *UserFilter, limit, offset int) ([]*dbmodels.User, error)
	CountUsers(ctx context.Context, filter *UserFilter) (int, error)
	ExistUser(ctx context.Context, userID string) (bool, error)
	CreateUser(ctx context.Context, user *dbmodels.User, password string, isInitializedPassword bool, now time.Time) (*dbmodels.User, error)
}

func NewUserRepository(db mysql.SQLExecutor) UserRepository {
	return &userRepository{
		db: db,
	}
}

type UserFilter struct {
	UserID *string
}

func (f *UserFilter) WhereMods() []qm.QueryMod {
	var filter []qm.QueryMod
	if f == nil {
		return filter
	}
	if f.UserID != nil {
		filter = append(filter, qm.Where(fmt.Sprintf("%s = ?", dbmodels.UserColumns.UserID), f.UserID))
	}
	return filter
}

type userRepository struct {
	db mysql.SQLExecutor
}

func (r *userRepository) GetUser(ctx context.Context, userID string) (*dbmodels.User, error) {
	user, err := dbmodels.FindUser(ctx, r.db, userID)
	if err == sql.ErrNoRows {
		return nil, errors.New(errors.UserNotFoundError, err)
	}
	return user, nil
}

func (r *userRepository) GetUsers(ctx context.Context, filter *UserFilter, limit, offset int) ([]*dbmodels.User, error) {
	mods := filter.WhereMods()
	mods = append(mods, qm.Limit(limit), qm.Offset(offset))

	users, err := dbmodels.Users(mods...).All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) CountUsers(ctx context.Context, filter *UserFilter) (int, error) {
	mods := filter.WhereMods()
	total, err := dbmodels.Users(mods...).Count(ctx, r.db)
	if err != nil {
		return 0, err
	}
	return int(total), nil
}

func (r *userRepository) ExistUser(ctx context.Context, userID string) (bool, error) {
	return dbmodels.UserExists(ctx, r.db, userID)
}

func (r *userRepository) CreateUser(ctx context.Context, user *dbmodels.User, password string, isInitializedPassword bool, now time.Time) (*dbmodels.User, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.New(errors.TxnBeginFatal, err)
	}

	dbPassword := &dbmodels.UserPassword{
		UserID:         user.UserID,
		Password:       password,
		LastModifiedAt: now,
		IsInitialized:  cast.BoolToInt8(isInitializedPassword),
	}
	if err := user.Insert(ctx, tx, boil.Infer()); err != nil {
		return nil, err
	}
	if err := dbPassword.Insert(ctx, tx, boil.Infer()); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, errors.New(errors.TxnRollbackFatal, err)
		}
		return nil, errors.New(errors.UserCreateFatal, err)
	}
	return user, nil
}

func (r *userRepository) toDBUserStatus(s entities.UserStatus) int {
	return int(s)
}
