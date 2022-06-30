package usecases

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/services"
	"time"
)

type OauthUseCase interface {
	CreateOauthClient(ctx context.Context, client *entities.OauthClient) (*entities.OauthClient, string, error)
	GetOauthClientRedirectURLs(ctx context.Context, oauthClientID string) ([]*entities.OAuthClientRedirectURL, error)
	AuthClientSecret(ctx context.Context, clientID, clientSecret string) (*entities.OauthClient, error)
	ValidateToAuthorizeUser(ctx context.Context, clientID, redirectURL string) (*entities.OauthClient, error)
	Authorize(ctx context.Context, userID, password, clientID, redirectURL string) (string, error)
	Oauth2ClientCredential(ctx context.Context, client *entities.OauthClient) (*entities.Oauth2ClientCredential, error)
	Oauth2AuthorizationCode(ctx context.Context, client *entities.OauthClient, code, redirectURL string, now time.Time) (*entities.Oauth2AuthorizationCode, error)
	Oauth2RefreshToken(ctx context.Context, client *entities.OauthClient, refreshToken string) (*entities.Oauth2AuthorizationCode, error)
	AuthAccessToken(ctx context.Context, accessToken string) (*entities.OauthSession, error)
}

func NewOauthUseCase(
	authService services.AuthService,
	userService services.UserService,
) OauthUseCase {
	return &oauthUseCase{
		authService: authService,
		userService: userService,
	}
}

type oauthUseCase struct {
	authService services.AuthService
	userService services.UserService
}

func (u *oauthUseCase) CreateOauthClient(ctx context.Context, client *entities.OauthClient) (*entities.OauthClient, string, error) {
	if err := u.authService.ValidateToCreateClient(ctx, client); err != nil {
		return nil, "", nil
	}

	dst, clientSecret, err := u.authService.CreateClient(ctx, client)
	if err != nil {
		return nil, "", err
	}

	return dst, clientSecret, nil
}

func (u *oauthUseCase) GetOauthClientRedirectURLs(ctx context.Context, oauthClientID string) ([]*entities.OAuthClientRedirectURL, error) {
	return u.authService.GetOAuthClientRedirectURLsByOAuthClientID(ctx, oauthClientID)
}

func (u *oauthUseCase) AuthClientSecret(ctx context.Context, clientID, clientSecret string) (*entities.OauthClient, error) {
	return u.authService.AuthClient(ctx, clientID, clientSecret)
}

func (u *oauthUseCase) Oauth2ClientCredential(ctx context.Context, client *entities.OauthClient) (*entities.Oauth2ClientCredential, error) {
	accessToken, expireIn, err := u.authService.PublishCCAccessToken(ctx, client)
	if err != nil {
		return nil, err
	}

	return &entities.Oauth2ClientCredential{
		AccessToken: accessToken,
		TokenType:   entities.OauthClientTypeClientCredential,
		ExpireIn:    int(expireIn),
	}, nil
}

func (u *oauthUseCase) ValidateToAuthorizeUser(ctx context.Context, clientID, redirectURL string) (*entities.OauthClient, error) {
	client, err := u.authService.GetUserClient(ctx, clientID)
	if err != nil {
		return nil, err
	}

	if err := u.authService.ValidateRedirectURL(ctx, clientID, redirectURL); err != nil {
		return nil, err
	}

	// TODO: scopeの確認
	return client, nil
}

func (u *oauthUseCase) Authorize(ctx context.Context, userID, password, clientID, redirectURL string) (string, error) {
	if err := u.userService.AuthUserPassword(ctx, userID, password); err != nil {
		return "", err
	}
	// クライアントＩＤの認証の有無
	if _, err := u.authService.GetUserClient(ctx, clientID); err != nil {
		return "", err
	}

	if err := u.authService.ValidateRedirectURL(ctx, clientID, redirectURL); err != nil {
		return "", err
	}

	code, err := u.authService.PublishAuthCode(ctx, clientID, userID, redirectURL)
	if err != nil {
		return "", err
	}

	return code, nil
}

func (u *oauthUseCase) Oauth2AuthorizationCode(ctx context.Context, client *entities.OauthClient, code, redirectURL string, now time.Time) (*entities.Oauth2AuthorizationCode, error) {
	oauthCode, err := u.authService.AuthCode(ctx, client, code, redirectURL)
	if err != nil {
		return nil, err
	}

	accessToken, expireIn, err := u.authService.PublishUserAccessToken(ctx, client, oauthCode.UserID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := u.authService.UpsertUserAuth(ctx, client.OauthClientID, oauthCode.UserID, now)
	if err != nil {
		return nil, err
	}

	return &entities.Oauth2AuthorizationCode{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpireIn:     expireIn,
	}, nil
}

func (u *oauthUseCase) Oauth2RefreshToken(ctx context.Context, client *entities.OauthClient, refreshToken string) (*entities.Oauth2AuthorizationCode, error) {
	ua, err := u.authService.AuthByRefreshToken(ctx, client.OauthClientID, refreshToken)
	if err != nil {
		return nil, err
	}

	accessToken, expireIn, err := u.authService.PublishUserAccessToken(ctx, client, ua.UserID)
	if err != nil {
		return nil, err
	}

	return &entities.Oauth2AuthorizationCode{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpireIn:     expireIn,
	}, nil
}

func (u *oauthUseCase) AuthAccessToken(ctx context.Context, accessToken string) (*entities.OauthSession, error) {
	return u.authService.GetSession(ctx, accessToken)
}
