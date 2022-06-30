package repositories

import (
	"context"
	"database/sql"
	"github.com/hiroyky/famiphoto/drivers/mysql"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/dbmodels"
)

type UserPasswordRepository interface {
	GetUserPassword(ctx context.Context, userID string) (*dbmodels.UserPassword, error)
}

func NewUserPasswordRepository(db mysql.SQLExecutor) UserPasswordRepository {
	return &userPasswordRepository{db: db}
}

type userPasswordRepository struct {
	db mysql.SQLExecutor
}

func (r *userPasswordRepository) GetUserPassword(ctx context.Context, userID string) (*dbmodels.UserPassword, error) {
	up, err := dbmodels.FindUserPassword(ctx, r.db, userID)
	if err == sql.ErrNoRows {
		return nil, errors.New(errors.UserPasswordNotFoundError, err)
	}
	if err != nil {
		return nil, err
	}

	return up, err
}
