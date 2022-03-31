package services

import (
	"context"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/usecases"
	"time"
)

func NewAuthService(
	passwordService usecases.PasswordService,
	randomService usecases.RandomService,
	oauthAccessTokenAdapter usecases.OauthAccessTokenAdapter,
	userAuthAdapter usecases.UserAuthAdapter,
	oauthCodeAdapter usecases.OauthCodeAdapter,
	oauthClientAdapter usecases.OauthClientAdapter,
	oauthClientURLAdapter usecases.OauthClientRedirectURLAdapter,
) usecases.AuthService {
	return &authService{
		passwordService:         passwordService,
		randomService:           randomService,
		oauthAccessTokenAdapter: oauthAccessTokenAdapter,
		userAuthAdapter:         userAuthAdapter,
		oauthCodeAdapter:        oauthCodeAdapter,
		oauthClientAdapter:      oauthClientAdapter,
		oauthClientURLAdapter:   oauthClientURLAdapter,
	}
}

type authService struct {
	passwordService         usecases.PasswordService
	randomService           usecases.RandomService
	oauthAccessTokenAdapter usecases.OauthAccessTokenAdapter
	userAuthAdapter         usecases.UserAuthAdapter
	oauthCodeAdapter        usecases.OauthCodeAdapter
	oauthClientAdapter      usecases.OauthClientAdapter
	oauthClientURLAdapter   usecases.OauthClientRedirectURLAdapter
}

func (s *authService) PublishUserAccessToken(ctx context.Context, client *entities.OauthClient, userID string) (string, int64, error) {
	accessToken, err := s.passwordService.GeneratePassword(config.AccessTokenLength)
	if err != nil {
		return "", 0, err
	}
	expireIn := config.Env.CCAccessTokenExpireInSec
	if err := s.oauthAccessTokenAdapter.SetUserAccessToken(
		ctx,
		client.OauthClientID,
		userID,
		accessToken,
		client.Scope,
		expireIn,
	); err != nil {
		return "", 0, err
	}
	return accessToken, expireIn, nil
}

func (s *authService) PublishCCAccessToken(ctx context.Context, client *entities.OauthClient) (string, int64, error) {
	accessToken, err := s.passwordService.GeneratePassword(config.AccessTokenLength)
	if err != nil {
		return "", 0, err
	}

	expireIn := config.Env.CCAccessTokenExpireInSec

	if err := s.oauthAccessTokenAdapter.SetClientCredentialAccessToken(
		ctx,
		client.OauthClientID,
		accessToken,
		expireIn,
	); err != nil {
		return "", 0, err
	}
	return accessToken, expireIn, nil
}

func (s *authService) GetSession(ctx context.Context, accessToken string) (*entities.OauthSession, error) {
	return s.oauthAccessTokenAdapter.GetSession(ctx, accessToken)
}

func (s *authService) AuthByRefreshToken(ctx context.Context, clientID, refreshToken string) (*entities.UserAuth, error) {
	ua, err := s.userAuthAdapter.GetUserAuthByRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	if ua.OAuthClientID != clientID {
		return nil, errors.New(errors.UserUnauthorizedError, nil)
	}
	return ua, nil
}

func (s *authService) UpsertUserAuth(ctx context.Context, clientID, userID string, now time.Time) (string, error) {
	refreshToken, err := s.passwordService.GeneratePassword(config.RefreshTokenLength)
	if err != nil {
		return "", err
	}

	ua := &entities.UserAuth{
		UserID:                  userID,
		OAuthClientID:           clientID,
		RefreshToken:            refreshToken,
		RefreshTokenPublishedAt: now.Unix(),
	}

	if _, err := s.userAuthAdapter.UpsertUserAuth(ctx, ua); err != nil {
		return "", err
	}
	return refreshToken, nil
}

func (s *authService) AuthCode(ctx context.Context, client *entities.OauthClient, code, redirectURL string) (*entities.OAuthCode, error) {
	oauthCode, err := s.oauthCodeAdapter.GetCode(ctx, code)
	if err != nil {
		return nil, err
	}
	if client.OauthClientID != oauthCode.ClientID {
		return nil, errors.New(errors.OAuthClientUnauthorizedError, nil)
	}
	if oauthCode.RedirectURL != redirectURL {
		return nil, errors.New(errors.OAuthClientUnauthorizedError, nil)
	}
	return oauthCode, nil
}

func (s *authService) PublishAuthCode(ctx context.Context, clientID, userID, redirectURL string, scope entities.OauthScope) (string, error) {
	code := s.randomService.GenerateRandomString(30)
	if err := s.oauthCodeAdapter.SetCode(ctx, &entities.OAuthCode{
		Code:        code,
		ClientID:    clientID,
		UserID:      userID,
		Scope:       scope,
		RedirectURL: redirectURL,
	}); err != nil {
		return "", err
	}
	return code, nil
}

func (s *authService) AuthClient(ctx context.Context, clientID, clientSecret string) (*entities.OauthClient, error) {
	client, err := s.oauthClientAdapter.GetByOauthClientID(ctx, clientID)
	if err != nil {
		return nil, err
	}

	if match, err := s.passwordService.MatchPassword(clientSecret, client.ClientSecretHashed); err != nil {
		return nil, err
	} else if !match {
		return nil, errors.New(errors.OAuthClientUnauthorizedError, nil)
	}
	return client, nil
}

func (s *authService) CreateClient(ctx context.Context, client *entities.OauthClient) (*entities.OauthClient, string, error) {
	clientSecret, err := s.passwordService.GeneratePassword(50)
	if err != nil {
		return nil, "", err
	}
	hashedClientSecret, err := s.passwordService.HashPassword(clientSecret)
	if err != nil {
		return nil, "", err
	}

	dst, err := s.oauthClientAdapter.CreateOAuthClient(ctx, client, hashedClientSecret)
	if err != nil {
		return nil, "", err
	}
	return dst, clientSecret, nil
}

func (s *authService) ValidateToCreateClient(ctx context.Context, client *entities.OauthClient) error {
	if exist, err := s.oauthClientAdapter.ExistOauthClient(ctx, client.OauthClientID); err != nil {
		return err
	} else if exist {
		return errors.New(errors.OAuthClientAlreadyExist, nil)
	}
	return nil
}

func (s *authService) GetUserClient(ctx context.Context, clientID string) (*entities.OauthClient, error) {
	client, err := s.oauthClientAdapter.GetByOauthClientID(ctx, clientID)
	if err != nil {
		return nil, err
	}
	if client.ClientType != entities.OauthClientTypeUserClient {
		return nil, errors.New(errors.OAuthClientNotFoundError, nil)
	}
	return client, nil
}

func (s *authService) ValidateRedirectURL(ctx context.Context, clientID, redirectURL string) error {
	urls, err := s.oauthClientURLAdapter.GetOAuthClientRedirectURLsByOAuthClientID(ctx, clientID)
	if err != nil {
		return err
	}
	if !urls.IsMatchURL(redirectURL) {
		return errors.New(errors.OAuthClientInvalidRedirectURLError, nil)
	}
	return nil
}
