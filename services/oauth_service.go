package services

import (
	"context"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures"
	"github.com/hiroyky/famiphoto/utils/random"
	"time"
)

type OAuthService interface {
	PublishUserAccessToken(ctx context.Context, client *entities.OauthClient, userID string) (string, int64, error)
	PublishCCAccessToken(ctx context.Context, client *entities.OauthClient) (string, int64, error)
	GetSession(ctx context.Context, accessToken string) (*entities.OauthSession, error)
	AuthByRefreshToken(ctx context.Context, clientID, refreshToken string) (*entities.UserAuth, error)
	UpsertUserAuth(ctx context.Context, clientID, userID string, now time.Time) (string, error)
	AuthCode(ctx context.Context, client *entities.OauthClient, code, redirectURL string) (*entities.OAuthCode, error)
	PublishAuthCode(ctx context.Context, clientID, userID, redirectURL string) (string, error)
	AuthClient(ctx context.Context, clientID, clientSecret string) (*entities.OauthClient, error)
	CreateClientWithClientSecret(ctx context.Context, client *entities.OauthClient, clientSecret string) error
	CreateClient(ctx context.Context, client *entities.OauthClient) (*entities.OauthClient, string, error)
	ValidateToCreateClient(ctx context.Context, client *entities.OauthClient) error
	GetUserClient(ctx context.Context, clientID string) (*entities.OauthClient, error)
	GetOAuthClientRedirectURLsByOAuthClientID(ctx context.Context, clientID string) (entities.OAuthClientRedirectURLList, error)
	ValidateRedirectURL(ctx context.Context, clientID, redirectURL string) error
}

func NewOAuthService(
	passwordService PasswordService,
	authAdapter infrastructures.OAuthAdapter,
) OAuthService {
	return &oauthService{
		passwordService:          passwordService,
		authAdapter:              authAdapter,
		generateRandomStringFunc: random.GenerateRandomString,
	}
}

type oauthService struct {
	passwordService          PasswordService
	authAdapter              infrastructures.OAuthAdapter
	generateRandomStringFunc func(length int) string
}

func (s *oauthService) PublishUserAccessToken(ctx context.Context, client *entities.OauthClient, userID string) (string, int64, error) {
	accessToken := s.generateRandomStringFunc(config.AccessTokenLength)
	expireIn := config.Env.CCAccessTokenExpireInSec
	if err := s.authAdapter.SetUserAccessToken(
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

func (s *oauthService) PublishCCAccessToken(ctx context.Context, client *entities.OauthClient) (string, int64, error) {
	accessToken := s.generateRandomStringFunc(config.AccessTokenLength)
	expireIn := config.Env.CCAccessTokenExpireInSec

	if err := s.authAdapter.SetClientCredentialAccessToken(
		ctx,
		client.OauthClientID,
		accessToken,
		expireIn,
	); err != nil {
		return "", 0, err
	}
	return accessToken, expireIn, nil
}

func (s *oauthService) GetSession(ctx context.Context, accessToken string) (*entities.OauthSession, error) {
	return s.authAdapter.GetSession(ctx, accessToken)
}

func (s *oauthService) AuthByRefreshToken(ctx context.Context, clientID, refreshToken string) (*entities.UserAuth, error) {
	ua, err := s.authAdapter.GetUserAuthByRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	if ua.OAuthClientID != clientID {
		return nil, errors.New(errors.UserUnauthorizedError, nil)
	}
	return ua, nil
}

func (s *oauthService) UpsertUserAuth(ctx context.Context, clientID, userID string, now time.Time) (string, error) {
	refreshToken := s.generateRandomStringFunc(config.RefreshTokenLength)
	ua := &entities.UserAuth{
		UserID:                  userID,
		OAuthClientID:           clientID,
		RefreshToken:            refreshToken,
		RefreshTokenPublishedAt: now.Unix(),
	}

	if _, err := s.authAdapter.UpsertUserAuth(ctx, ua); err != nil {
		return "", err
	}
	return refreshToken, nil
}

func (s *oauthService) AuthCode(ctx context.Context, client *entities.OauthClient, code, redirectURL string) (*entities.OAuthCode, error) {
	oauthCode, err := s.authAdapter.GetCode(ctx, code)
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

func (s *oauthService) PublishAuthCode(ctx context.Context, clientID, userID, redirectURL string) (string, error) {
	code := s.generateRandomStringFunc(30)
	if err := s.authAdapter.SetCode(ctx, &entities.OAuthCode{
		Code:        code,
		ClientID:    clientID,
		UserID:      userID,
		RedirectURL: redirectURL,
	}); err != nil {
		return "", err
	}
	return code, nil
}

func (s *oauthService) AuthClient(ctx context.Context, clientID, clientSecret string) (*entities.OauthClient, error) {
	client, err := s.authAdapter.GetByOauthClientID(ctx, clientID)
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

func (s *oauthService) CreateClientWithClientSecret(ctx context.Context, client *entities.OauthClient, clientSecret string) error {
	hashedClientSecret, err := s.passwordService.HashPassword(clientSecret)
	if err != nil {
		return err
	}
	if _, err := s.authAdapter.CreateOAuthClient(ctx, client, hashedClientSecret); err != nil {
		return err
	}
	return nil
}

func (s *oauthService) CreateClient(ctx context.Context, client *entities.OauthClient) (*entities.OauthClient, string, error) {
	clientSecret, err := s.passwordService.GeneratePassword(50)
	if err != nil {
		return nil, "", err
	}
	hashedClientSecret, err := s.passwordService.HashPassword(clientSecret)
	if err != nil {
		return nil, "", err
	}

	dst, err := s.authAdapter.CreateOAuthClient(ctx, client, hashedClientSecret)
	if err != nil {
		return nil, "", err
	}
	return dst, clientSecret, nil
}

func (s *oauthService) ValidateToCreateClient(ctx context.Context, client *entities.OauthClient) error {
	if exist, err := s.authAdapter.ExistOauthClient(ctx, client.OauthClientID); err != nil {
		return err
	} else if exist {
		return errors.New(errors.OAuthClientAlreadyExist, nil)
	}
	return nil
}

func (s *oauthService) GetUserClient(ctx context.Context, clientID string) (*entities.OauthClient, error) {
	client, err := s.authAdapter.GetByOauthClientID(ctx, clientID)
	if err != nil {
		return nil, err
	}
	if client.ClientType != entities.OauthClientTypeUserClient {
		return nil, errors.New(errors.OAuthClientNotFoundError, nil)
	}
	return client, nil
}

func (s *oauthService) GetOAuthClientRedirectURLsByOAuthClientID(ctx context.Context, clientID string) (entities.OAuthClientRedirectURLList, error) {
	return s.authAdapter.GetOAuthClientRedirectURLsByOAuthClientID(ctx, clientID)
}

func (s *oauthService) ValidateRedirectURL(ctx context.Context, clientID, redirectURL string) error {
	urls, err := s.authAdapter.GetOAuthClientRedirectURLsByOAuthClientID(ctx, clientID)
	if err != nil {
		return err
	}
	if !urls.IsMatchURL(redirectURL) {
		return errors.New(errors.OAuthClientInvalidRedirectURLError, nil)
	}
	return nil
}
