package usecases

import (
	"context"
	"fmt"
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
	ExistUser(ctx context.Context, userID string) (bool, error)
	GetUserPassword(ctx context.Context, userID string) (*entities.UserPassword, error)
	GetUsersBelongingGroup(ctx context.Context, groupID string, limit, offset int) (entities.UserList, int, error)
	Login(ctx context.Context, client *entities.OauthClient, userID, password string, now time.Time) (*entities.Oauth2AuthorizationCode, error)
}

func NewUserUseCase(
	userAdapter infrastructures.UserAdapter,
	groupAdapter infrastructures.GroupAdapter,
	userService services.UserService,
	authService services.OAuthService,
	passwordService services.PasswordService,
) UserUseCase {
	return &userUseCase{
		userAdapter:     userAdapter,
		groupAdapter:    groupAdapter,
		userService:     userService,
		authService:     authService,
		passwordService: passwordService,
	}
}

type userUseCase struct {
	userAdapter     infrastructures.UserAdapter
	groupAdapter    infrastructures.GroupAdapter
	userService     services.UserService
	authService     services.OAuthService
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
	if exist, err := u.userAdapter.ExistUser(ctx, userID); err != nil {
		return nil, err
	} else if exist {
		return nil, errors.New(errors.UserAlreadyExists, nil)
	}
	if exist, err := u.groupAdapter.ExistGroup(ctx, userID); err != nil {
		return nil, err
	} else if exist {
		return nil, errors.New(errors.GroupAlreadyExistError, nil)
	}

	user := &entities.User{
		UserID: userID,
		Name:   name,
		Status: entities.UserStatusActive,
	}
	groupName := fmt.Sprintf("%s (%s)", name, userID)

	encPassword, err := u.passwordService.HashPassword(password)
	if err != nil {
		return nil, err
	}
	createdUser, err := u.userAdapter.CreateUser(ctx, user, groupName, encPassword, true, now)
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

func (u *userUseCase) ExistUser(ctx context.Context, userID string) (bool, error) {
	if _, err := u.userAdapter.GetUser(ctx, userID); err != nil {
		if errors.IsErrCode(err, errors.UserNotFoundError) {
			return false, nil
		}
		return false, err
	}
	return true, nil
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

func (u *userUseCase) Login(ctx context.Context, client *entities.OauthClient, userID, password string, now time.Time) (*entities.Oauth2AuthorizationCode, error) {
	if err := u.userService.AuthUserPassword(ctx, userID, password); err != nil {
		return nil, err
	}

	accessToken, expireIn, err := u.authService.PublishUserAccessToken(ctx, client, userID)
	if err != nil {
		return nil, err
	}
	refreshToken, err := u.authService.UpsertUserAuth(ctx, client.OauthClientID, userID, now)
	if err != nil {
		return nil, err
	}
	return &entities.Oauth2AuthorizationCode{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpireIn:     expireIn,
	}, nil
}
