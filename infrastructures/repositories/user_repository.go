package repositories

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/models"
	"github.com/hiroyky/famiphoto/usecases"
	"github.com/hiroyky/famiphoto/utils/cast"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"time"
)

func NewUserRepository(db SQLExecutor) usecases.UserAdapter {
	return &userRepository{
		db: db,
	}
}

type userRepository struct {
	db SQLExecutor
}

func (r *userRepository) GetUser(ctx context.Context, userID string) (*entities.User, error) {
	return nil, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user *entities.User, password string, isInitializedPassword bool, now time.Time) (*entities.User, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.New(errors.TxnBeginFatal, err)
	}

	dbUser := &models.User{
		UserID: user.UserID,
		Name:   user.Name,
		Status: r.toDBUserStatus(user.Status),
	}
	dbPassword := &models.UserPassword{
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
	return &entities.User{
		UserID: dbUser.UserID,
		Name:   dbUser.Name,
		Status: r.toEntityUserStatus(dbUser.Status),
	}, nil
}

func (r *userRepository) toDBUserStatus(s entities.UserStatus) int {
	return int(s)
}

func (r *userRepository) toEntityUserStatus(s int) entities.UserStatus {
	return entities.UserStatus(s)
}
