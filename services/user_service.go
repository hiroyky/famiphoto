package services

import (
	"context"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/usecases"
)

func NewUserService(
	userAdapter usecases.UserAdapter,
	userPasswordAdapter usecases.UserPasswordAdapter,
	passwordService usecases.PasswordService,
) usecases.UserService {
	return &userService{
		userAdapter:         userAdapter,
		userPasswordAdapter: userPasswordAdapter,
		passwordService:     passwordService,
	}
}

type userService struct {
	userAdapter         usecases.UserAdapter
	userPasswordAdapter usecases.UserPasswordAdapter
	passwordService     usecases.PasswordService
}

func (s *userService) AuthUserPassword(ctx context.Context, userID, password string) error {
	if exist, err := s.userAdapter.ExistUser(ctx, userID); err != nil {
		return err
	} else if !exist {
		return errors.New(errors.UserUnauthorizedError, nil)
	}
	pass, err := s.userPasswordAdapter.GetUserPassword(ctx, userID)
	if err != nil {
		return err
	}
	if correct, err := s.passwordService.MatchPassword(password, pass.Password); err != nil {
		return err
	} else if !correct {
		return errors.New(errors.UserUnauthorizedError, nil)
	}
	return nil
}
