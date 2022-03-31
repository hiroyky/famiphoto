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
	AuthClientSecret(ctx context.Context, clientID, clientSecret string) (*entities.OauthClient, error)
	ValidateToAuthorizeUser(ctx context.Context, clientID, redirectURL, scope string) (*entities.OauthClient, error)
	Authorize(ctx context.Context, userID, password, clientID, redirectURL, scope string) (string, error)
	Oauth2ClientCredential(ctx context.Context, client *entities.OauthClient) (*entities.Oauth2ClientCredential, error)
	Oauth2AuthorizationCode(ctx context.Context, client *entities.OauthClient, code, redirectURL string, now time.Time) (*entities.Oauth2AuthorizationCode, error)
	Oauth2RefreshToken(ctx context.Context, client *entities.OauthClient, refreshToken string) (*entities.Oauth2AuthorizationCode, error)
	AuthAccessToken(ctx context.Context, accessToken string) (*entities.OauthSession, error)
}

func NewOauthUseCase(
	oauthClientAdapter OauthClientAdapter,
	oauthClientURLAdapter OauthClientRedirectURLAdapter,
	oauthAccessTokenAdapter OauthAccessTokenAdapter,
	oauthCodeAdapter OauthCodeAdapter,
	userService UserService,
	passwordService PasswordService,
	randomService RandomService,
) OauthUseCase {
	return &oauthUseCase{
		oauthClientAdapter:      oauthClientAdapter,
		oauthClientURLAdapter:   oauthClientURLAdapter,
		oauthAccessTokenAdapter: oauthAccessTokenAdapter,
		oauthCodeAdapter:        oauthCodeAdapter,
		userService:             userService,
		passwordService:         passwordService,
		randomService:           randomService,
	}
}

type oauthUseCase struct {
	oauthClientAdapter      OauthClientAdapter
	oauthClientURLAdapter   OauthClientRedirectURLAdapter
	oauthAccessTokenAdapter OauthAccessTokenAdapter
	oauthCodeAdapter        OauthCodeAdapter
	userAuthAdapter         UserAuthAdapter
	userService             UserService
	passwordService         PasswordService
	randomService           RandomService
}

func (u *oauthUseCase) CreateOauthClient(ctx context.Context, client *entities.OauthClient) (*entities.OauthClient, string, error) {
	if exist, err := u.oauthClientAdapter.ExistOauthClient(ctx, client.OauthClientID); err != nil {
		return nil, "", err
	} else if exist {
		return nil, "", errors.New(errors.OAuthClientAlreadyExist, nil)
	}

	clientSecret, err := u.passwordService.GeneratePassword(50)
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

func (u *oauthUseCase) AuthClientSecret(ctx context.Context, clientID, clientSecret string) (*entities.OauthClient, error) {
	client, err := u.oauthClientAdapter.GetByOauthClientID(ctx, clientID)
	if err != nil {
		code := errors.GetFPErrorCode(err)
		if code == errors.OAuthClientNotFoundError {
			return nil, errors.New(errors.OAuthClientUnauthorizedError, nil)
		}
		return nil, err
	}

	if match, err := u.passwordService.MatchPassword(clientSecret, client.ClientSecretHashed); err != nil {
		return nil, err
	} else if !match {
		return nil, errors.New(errors.OAuthClientUnauthorizedError, nil)
	}
	return client, nil
}

func (u *oauthUseCase) Oauth2ClientCredential(ctx context.Context, client *entities.OauthClient) (*entities.Oauth2ClientCredential, error) {
	accessToken, err := u.passwordService.GeneratePassword(config.AccessTokenLength)
	if err != nil {
		return nil, err
	}

	expireIn := config.Env.CCAccessTokenExpireInSec

	if err := u.oauthAccessTokenAdapter.SetClientCredentialAccessToken(
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

func (u *oauthUseCase) ValidateToAuthorizeUser(ctx context.Context, clientID, redirectURL, scope string) (*entities.OauthClient, error) {
	client, err := u.oauthClientAdapter.GetByOauthClientID(ctx, clientID)
	if err != nil {
		return nil, err
	}
	if client.ClientType != entities.OauthClientTypeUserClient {
		return nil, errors.New(errors.OAuthClientNotFoundError, nil)
	}
	urls, err := u.oauthClientURLAdapter.GetOAuthClientRedirectURLsByOAuthClientID(ctx, clientID)
	if err != nil {
		return nil, err
	}
	if !urls.IsMatchURL(redirectURL) {
		return nil, errors.New(errors.OAuthClientInvalidRedirectURLError, nil)
	}
	// TODO: scopeの確認
	return client, nil
}

func (u *oauthUseCase) Authorize(ctx context.Context, userID, password, clientID, redirectURL, scope string) (string, error) {
	if err := u.userService.AuthUserPassword(ctx, userID, password); err != nil {
		return "", err
	}
	// クライアントＩＤの認証の有無
	_, err := u.oauthClientAdapter.GetByOauthClientID(ctx, clientID)
	if err != nil {
		return "", err
	}
	redirectURLs, err := u.oauthClientURLAdapter.GetOAuthClientRedirectURLsByOAuthClientID(ctx, clientID)
	if err != nil {
		return "", err
	}
	if !redirectURLs.IsMatchURL(redirectURL) {
		return "", errors.New(errors.OAuthClientInvalidRedirectURLError, nil)
	}

	// TODO: scopeの確認

	// 一時コードの発行
	code := u.randomService.GenerateRandomString(30)
	if err := u.oauthCodeAdapter.SetCode(ctx, &entities.OAuthCode{
		Code:        code,
		ClientID:    clientID,
		UserID:      userID,
		Scope:       entities.OauthScope(scope),
		RedirectURL: redirectURL,
	}); err != nil {
		return "", err
	}

	return code, nil
}

func (u *oauthUseCase) Oauth2AuthorizationCode(ctx context.Context, client *entities.OauthClient, code, redirectURL string, now time.Time) (*entities.Oauth2AuthorizationCode, error) {
	oauthCode, err := u.oauthCodeAdapter.GetCode(ctx, code)
	if err != nil {
		return nil, err
	}
	if client.OauthClientID != oauthCode.ClientID {
		return nil, errors.New(errors.OAuthClientUnauthorizedError, nil)
	}
	if oauthCode.RedirectURL != redirectURL {
		return nil, errors.New(errors.OAuthClientUnauthorizedError, nil)
	}

	accessToken, err := u.passwordService.GeneratePassword(config.AccessTokenLength)
	if err != nil {
		return nil, err
	}
	if err := u.oauthAccessTokenAdapter.SetUserAccessToken(
		ctx,
		oauthCode.ClientID,
		oauthCode.UserID,
		accessToken,
		oauthCode.Scope,
		config.Env.UserAccessTokenExpireInSec,
	); err != nil {
		return nil, err
	}

	refreshToken, err := u.passwordService.GeneratePassword(config.RefreshTokenLength)
	if err != nil {
		return nil, err
	}

	ua := &entities.UserAuth{
		UserID:                  oauthCode.UserID,
		OAuthClientID:           oauthCode.ClientID,
		RefreshToken:            refreshToken,
		RefreshTokenPublishedAt: now.Unix(),
	}

	if _, err := u.userAuthAdapter.UpsertUserAuth(ctx, ua); err != nil {
		return nil, err
	}

	return &entities.Oauth2AuthorizationCode{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpireIn:     config.Env.UserAccessTokenExpireInSec,
	}, nil
}

func (u *oauthUseCase) Oauth2RefreshToken(ctx context.Context, client *entities.OauthClient, refreshToken string) (*entities.Oauth2AuthorizationCode, error) {
	ua, err := u.userAuthAdapter.GetUserAuthByRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	if ua.OAuthClientID != client.OauthClientID {
		return nil, errors.New(errors.UserUnauthorizedError, nil)
	}

	accessToken, err := u.passwordService.GeneratePassword(config.AccessTokenLength)
	if err != nil {
		return nil, err
	}
	if err := u.oauthAccessTokenAdapter.SetUserAccessToken(
		ctx,
		ua.OAuthClientID,
		ua.UserID,
		accessToken,
		client.Scope,
		config.Env.UserAccessTokenExpireInSec,
	); err != nil {
		return nil, err
	}

	return &entities.Oauth2AuthorizationCode{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpireIn:     config.Env.UserAccessTokenExpireInSec,
	}, nil
}

func (u *oauthUseCase) AuthAccessToken(ctx context.Context, accessToken string) (*entities.OauthSession, error) {
	sess, err := u.oauthAccessTokenAdapter.GetSession(ctx, accessToken)
	if err != nil {
		code := errors.GetFPErrorCode(err)
		if code == errors.OAuthAccessTokenNotFoundError {
			return nil, errors.New(errors.UnauthorizedError, err)
		}
		return nil, err
	}
	return sess, nil
}
