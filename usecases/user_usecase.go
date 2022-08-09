package usecases

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures"
	"github.com/hiroyky/famiphoto/infrastructures/filters"
	"github.com/hiroyky/famiphoto/services"
	"time"
)

type UserUseCase interface {
	ValidateToCreateUser(ctx context.Context, userID, name string, password string) error
	CreateUser(ctx context.Context, userID, name string, password string, now time.Time) (*entities.User, error)
	GetUser(ctx context.Context, userID string) (*entities.User, error)
	GetUsers(ctx context.Context, userID *string, limit, offset int) (entities.UserList, int, error)
	GetUserPassword(ctx context.Context, userID string) (*entities.UserPassword, error)
	GetUsersBelongingGroup(ctx context.Context, groupID string, limit, offset int) (entities.UserList, int, error)
}

func NewUserUseCase(
	userAdapter infrastructures.UserAdapter,
	passwordService services.PasswordService,
) UserUseCase {
	return &userUseCase{
		userAdapter:     userAdapter,
		passwordService: passwordService,
	}
}

type userUseCase struct {
	userAdapter     infrastructures.UserAdapter
	passwordService services.PasswordService
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
	filter := &filters.UserFilter{UserID: userID}
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

func (u *userUseCase) GetUserPassword(ctx context.Context, userID string) (*entities.UserPassword, error) {
	p, err := u.userAdapter.GetUserPassword(ctx, userID)
	if err != nil {
		return nil, err
	}
	p.Password = ""
	return p, nil
}

func (u *userUseCase) GetUsersBelongingGroup(ctx context.Context, groupID string, limit, offset int) (entities.UserList, int, error) {
	users, err := u.userAdapter.GetUsersBelongingGroup(ctx, groupID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	total, err := u.userAdapter.CountUsersBelongingGroup(ctx, groupID)
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}
