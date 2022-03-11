package usecases

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"time"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, userID, name string, password string, now time.Time) (*entities.User, error)
}

func NewUserUseCase(userAdapter UserAdapter) UserUseCase {
	return &userUseCase{userAdapter: userAdapter}
}

type userUseCase struct {
	userAdapter UserAdapter
}

func (u *userUseCase) CreateUser(ctx context.Context, userID, name string, password string, now time.Time) (*entities.User, error) {
	user := &entities.User{
		UserID: userID,
		Name:   name,
		Status: entities.UserStatusActive,
	}

	createdUser, err := u.userAdapter.CreateUser(ctx, user, password, true, now)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
