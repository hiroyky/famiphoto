package usecases

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
)

type OauthUseClientCase interface {
	CreateOauthClient(ctx context.Context, client *entities.OauthClient) (*entities.OauthClient, string, error)
}

func NewOauthUseCase(
	oauthClientAdapter OauthClientAdapter,
	passwordService PasswordService,
) OauthUseClientCase {
	return &oauthUseCase{
		oauthClientAdapter: oauthClientAdapter,
		passwordService:    passwordService,
	}
}

type oauthUseCase struct {
	oauthClientAdapter OauthClientAdapter
	passwordService    PasswordService
}

func (u *oauthUseCase) CreateOauthClient(ctx context.Context, client *entities.OauthClient) (*entities.OauthClient, string, error) {
	if exist, err := u.oauthClientAdapter.ExistOauthClient(ctx, client.OauthClientID); err != nil {
		return nil, "", err
	} else if exist {
		return nil, "", errors.New(errors.OAuthClientAlreadyExist, nil)
	}

	clientSecret, err := u.passwordService.GeneratePassword()
	if err != nil {
		return nil, "", err
	}
	hashedClientSecret, err := u.passwordService.HashPassword(clientSecret)
	if err != nil {
		return nil, "", err
	}

	dst, err := u.oauthClientAdapter.CreateOAuthClient(ctx, client, hashedClientSecret)
	if err != nil {
		return nil, "", err
	}

	return dst, clientSecret, nil
}
