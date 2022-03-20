package usecases

import (
	"context"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"time"
)

type OauthUseCase interface {
	CreateOauthClient(ctx context.Context, client *entities.OauthClient) (*entities.OauthClient, string, error)
	GetOauthClientRedirectURLs(ctx context.Context, oauthClientID string) ([]*entities.OAuthClientRedirectURL, error)
	Oauth2ClientCredential(ctx context.Context, clientID, clientSecret string, now time.Time) (*entities.Oauth2ClientCredential, error)
}

func NewOauthUseCase(
	oauthClientAdapter OauthClientAdapter,
	oauthClientURLAdapter OauthClientRedirectURLAdapter,
	oOauthAccessTokenAdapter OauthAccessTokenAdapter,
	passwordService PasswordService,
) OauthUseCase {
	return &oauthUseCase{
		oauthClientAdapter: oauthClientAdapter,
		passwordService:    passwordService,
	}
}

type oauthUseCase struct {
	oauthClientAdapter       OauthClientAdapter
	oauthClientURLAdapter    OauthClientRedirectURLAdapter
	oOauthAccessTokenAdapter OauthAccessTokenAdapter
	passwordService          PasswordService
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

func (u *oauthUseCase) GetOauthClientRedirectURLs(ctx context.Context, oauthClientID string) ([]*entities.OAuthClientRedirectURL, error) {
	return u.oauthClientURLAdapter.GetOAuthClientRedirectURLsByOAuthClientID(ctx, oauthClientID)
}

func (u *oauthUseCase) Oauth2ClientCredential(ctx context.Context, clientID, clientSecret string, now time.Time) (*entities.Oauth2ClientCredential, error) {
	client, err := u.oauthClientAdapter.GetByOauthClientID(ctx, clientID)
	if err != nil {
		return nil, err
	}

	if match, err := u.passwordService.MatchPassword(clientSecret, client.ClientSecretHashed); err != nil {
		return nil, err
	} else if !match {
		return nil, errors.New(errors.OAuthClientNotFoundError, nil)
	}

	accessToken, err := u.passwordService.GeneratePassword()
	if err != nil {
		return nil, err
	}

	expireIn := config.Env.CCAccessTokenExpireInSec

	if err := u.oOauthAccessTokenAdapter.SetClientCredentialAccessToken(
		ctx,
		client.OauthClientID,
		accessToken,
		expireIn,
	); err != nil {
		return nil, err
	}

	return &entities.Oauth2ClientCredential{
		AccessToken: accessToken,
		TokenType:   entities.OauthClientTypeClientCredential,
		ExpireIn:    int(expireIn),
	}, nil
}
