package usecases

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"time"
)

type UserUseCase interface {
	ValidateToCreateUser(ctx context.Context, userID, name string, password string) error
	CreateUser(ctx context.Context, userID, name string, password string, now time.Time) (*entities.User, error)
	GetUser(ctx context.Context, userID string) (*entities.User, error)
	GetUsers(ctx context.Context, userID *string, limit, offset int) (entities.UserList, int, error)
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

func (u *userUseCase) ValidateToCreateUser(ctx context.Context, userID, name string, password string) error {
	if exists, err := u.userAdapter.ExistUser(ctx, userID); err != nil {
		return err
	} else if exists {
		return errors.New(errors.UserAlreadyExists, nil)
	}

	if len(password) < 5 {
		return errors.New(errors.PasswordWeakError, nil)
	}

	return nil
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

func (u *userUseCase) GetUser(ctx context.Context, userID string) (*entities.User, error) {
	return u.userAdapter.GetUser(ctx, userID)
}

func (u *userUseCase) GetUsers(ctx context.Context, userID *string, limit, offset int) (entities.UserList, int, error) {
	filter := &UserFilter{UserID: userID}
	users, err := u.userAdapter.GetUsers(ctx, filter, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	total, err := u.userAdapter.CountUsers(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}
