package services

import (
	"context"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures"
)

type UserService interface {
	AuthUserPassword(ctx context.Context, userID, password string) error
}

func NewUserService(
	userAdapter infrastructures.UserAdapter,
	passwordService PasswordService,
) UserService {
	return &userService{
		userAdapter:     userAdapter,
		passwordService: passwordService,
	}
}

type userService struct {
	userAdapter     infrastructures.UserAdapter
	passwordService PasswordService
}

func (s *userService) AuthUserPassword(ctx context.Context, userID, password string) error {
	if exist, err := s.userAdapter.ExistUser(ctx, userID); err != nil {
		return err
	} else if !exist {
		return errors.New(errors.UserUnauthorizedError, nil)
	}
	pass, err := s.userAdapter.GetUserPassword(ctx, userID)
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
