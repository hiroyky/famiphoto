package repositories

import (
	"context"
	"database/sql"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/dbmodels"
	"github.com/hiroyky/famiphoto/usecases"
	"github.com/hiroyky/famiphoto/utils/cast"
)

func NewUserPasswordRepository(db SQLExecutor) usecases.UserPasswordAdapter {
	return &userPasswordRepository{db: db}
}

type userPasswordRepository struct {
	db SQLExecutor
}

func (r *userPasswordRepository) GetUserPassword(ctx context.Context, userID string) (*entities.UserPassword, error) {
	up, err := dbmodels.FindUserPassword(ctx, r.db, userID)
	if err == sql.ErrNoRows {
		return nil, errors.New(errors.UserPasswordNotFoundError, err)
	}
	if err != nil {
		return nil, err
	}
	return &entities.UserPassword{
		UserId:         up.UserID,
		Password:       up.Password,
		LastModifiedAt: up.LastModifiedAt,
		IsInitialized:  cast.IntToBool(up.IsInitialized),
	}, nil
}
