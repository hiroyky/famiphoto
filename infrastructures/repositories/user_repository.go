package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/dbmodels"
	"github.com/hiroyky/famiphoto/usecases"
	"github.com/hiroyky/famiphoto/utils/cast"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"time"
)

func NewUserRepository(db SQLExecutor) usecases.UserAdapter {
	return &userRepository{
		db: db,
	}
}

type userFilter struct {
	*usecases.UserFilter
}

func (f *userFilter) WhereMods() []qm.QueryMod {
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
	db SQLExecutor
}

func (r *userRepository) GetUser(ctx context.Context, userID string) (*entities.User, error) {
	user, err := dbmodels.FindUser(ctx, r.db, userID)
	if err != sql.ErrNoRows {
		return nil, errors.New(errors.UserNotFoundError, err)
	}
	return r.toUserEntity(user), nil
}

func (r *userRepository) GetUsers(ctx context.Context, filter *usecases.UserFilter, limit, offset int) (entities.UserList, error) {
	f := &userFilter{UserFilter: filter}
	mods := f.WhereMods()
	mods = append(mods, qm.Limit(limit), qm.Offset(offset))

	users, err := dbmodels.Users(mods...).All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return cast.Array(users, r.toUserEntity), err
}

func (r *userRepository) CountUsers(ctx context.Context, filter *usecases.UserFilter) (int, error) {
	f := &userFilter{UserFilter: filter}
	mods := f.WhereMods()
	total, err := dbmodels.Users(mods...).Count(ctx, r.db)
	if err != nil {
		return 0, err
	}
	return int(total), nil
}

func (r *userRepository) ExistUser(ctx context.Context, userID string) (bool, error) {
	return dbmodels.UserExists(ctx, r.db, userID)
}

func (r *userRepository) CreateUser(ctx context.Context, user *entities.User, password string, isInitializedPassword bool, now time.Time) (*entities.User, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.New(errors.TxnBeginFatal, err)
	}

	dbUser := &dbmodels.User{
		UserID: user.UserID,
		Name:   user.Name,
		Status: r.toDBUserStatus(user.Status),
	}
	dbPassword := &dbmodels.UserPassword{
		UserID:         user.UserID,
		Password:       password,
		LastModifiedAt: now,
		IsInitialized:  cast.BoolToInt8(isInitializedPassword),
	}
	if err := dbUser.Insert(ctx, tx, boil.Infer()); err != nil {
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
	return r.toUserEntity(dbUser), nil
}

func (r *userRepository) toDBUserStatus(s entities.UserStatus) int {
	return int(s)
}

func (r *userRepository) toEntityUserStatus(s int) entities.UserStatus {
	return entities.UserStatus(s)
}

func (r *userRepository) toUserEntity(user *dbmodels.User) *entities.User {
	return &entities.User{
		UserID: user.UserID,
		Name:   user.Name,
		Status: r.toEntityUserStatus(user.Status),
	}
}
