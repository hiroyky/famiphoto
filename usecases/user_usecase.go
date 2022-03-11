package usecases

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"time"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, userID, name string, password string, now time.Time) (*entities.User, error)
}

func NewUserUseCase(userAdapter UserAdapter, passwordService PasswordService) UserUseCase {
	return &userUseCase{
		userAdapter:     userAdapter,
		passwordService: passwordService,
	}
}

type userUseCase struct {
	userAdapter     UserAdapter
	passwordService PasswordService
}

func (u *userUseCase) CreateUser(ctx context.Context, userID, name string, password string, now time.Time) (*entities.User, error) {
	user := &entities.User{
		UserID: userID,
		Name:   name,
		Status: entities.UserStatusActive,
	}

	encPassword, err := u.passwordService.HashPassword(password)
	if err != nil {
		return nil, err
	}

	createdUser, err := u.userAdapter.CreateUser(ctx, user, encPassword, true, now)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
