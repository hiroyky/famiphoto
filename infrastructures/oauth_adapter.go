package infrastructures

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/infrastructures/dbmodels"
	"github.com/hiroyky/famiphoto/infrastructures/models"
	"github.com/hiroyky/famiphoto/infrastructures/repositories"
	"github.com/hiroyky/famiphoto/utils/array"
	"github.com/hiroyky/famiphoto/utils/cast"
)

type OAuthAdapter interface {
	GetByOauthClientID(ctx context.Context, id string) (*entities.OauthClient, error)
	GetOAuthClientRedirectURLsByOAuthClientID(ctx context.Context, oauthClientID string) (entities.OAuthClientRedirectURLList, error)
	CreateOAuthClient(ctx context.Context, client *entities.OauthClient, clientSecret string) (*entities.OauthClient, error)
	ExistOauthClient(ctx context.Context, id string) (bool, error)

	SetClientCredentialAccessToken(ctx context.Context, clientID, accessToken string, expireAt int64) error
	SetUserAccessToken(ctx context.Context, clientID, userID, accessToken string, scope entities.OauthScope, expireIn int64) error
	GetSession(ctx context.Context, accessToken string) (*entities.OauthSession, error)

	SetCode(ctx context.Context, code *entities.OAuthCode) error
	GetCode(ctx context.Context, code string) (*entities.OAuthCode, error)

	UpsertUserAuth(ctx context.Context, userAuth *entities.UserAuth) (*entities.UserAuth, error)
	GetUserAuth(ctx context.Context, userID, clientID string) (*entities.UserAuth, error)
	GetUserAuthByRefreshToken(ctx context.Context, refreshToken string) (*entities.UserAuth, error)
	DeleteUserAuth(ctx context.Context, userID, clientID string) error
	DeleteClientAllAuth(ctx context.Context, clientID string) error
}

func NewAuthAdapter(
	oauthClientRepo repositories.OAuthClientRepository,
	oauthAccessTokenRepo repositories.OauthAccessTokenRepository,
	oauthClientRedirectURLRepo repositories.OAuthClientRedirectURLRepository,
	oauthCodeRepo repositories.OauthCodeRepository,
	userAuthRepository repositories.UserAuthRepository,
) OAuthAdapter {
	return &oauthAdapter{
		oauthClientRepo:            oauthClientRepo,
		oauthAccessTokenRepo:       oauthAccessTokenRepo,
		oauthClientRedirectURLRepo: oauthClientRedirectURLRepo,
		oauthCodeRepo:              oauthCodeRepo,
		userAuthRepository:         userAuthRepository,
	}
}

type oauthAdapter struct {
	oauthClientRepo            repositories.OAuthClientRepository
	oauthAccessTokenRepo       repositories.OauthAccessTokenRepository
	oauthClientRedirectURLRepo repositories.OAuthClientRedirectURLRepository
	oauthCodeRepo              repositories.OauthCodeRepository
	userAuthRepository         repositories.UserAuthRepository
}

func (a *oauthAdapter) GetByOauthClientID(ctx context.Context, id string) (*entities.OauthClient, error) {
	oauthClient, err := a.oauthClientRepo.GetByOauthClientID(ctx, id)
	if err != nil {
		return nil, err
	}

	redirectURLs, err := a.oauthClientRedirectURLRepo.GetOAuthClientRedirectURLsByOAuthClientID(ctx, id)
	if err != nil {
		return nil, err
	}
	return a.toOAuthClientEntity(oauthClient, redirectURLs), nil
}

func (a *oauthAdapter) GetOAuthClientRedirectURLsByOAuthClientID(ctx context.Context, oauthClientID string) (entities.OAuthClientRedirectURLList, error) {
	urls, err := a.oauthClientRedirectURLRepo.GetOAuthClientRedirectURLsByOAuthClientID(ctx, oauthClientID)
	if err != nil {
		return nil, err
	}
	return array.Map(urls, a.toRedirectURLEntity), nil
}

func (a *oauthAdapter) CreateOAuthClient(ctx context.Context, client *entities.OauthClient, clientSecret string) (*entities.OauthClient, error) {
	dbClient := &dbmodels.OauthClient{
		OauthClientID: client.OauthClientID,
		Name:          client.Name,
		ClientSecret:  clientSecret,
		Scope:         client.Scope.String(),
		ClientType:    client.ClientType.Int(),
	}
	dbURLS := array.Map(client.RedirectURLs, func(t string) *dbmodels.OauthClientRedirectURL {
		return &dbmodels.OauthClientRedirectURL{
			OauthClientID: client.OauthClientID,
			RedirectURL:   t,
		}
	})

	dstClient, dstURLs, err := a.oauthClientRepo.CreateOAuthClient(ctx, dbClient, dbURLS)
	if err != nil {
		return nil, err
	}

	return a.toOAuthClientEntity(dstClient, dstURLs), nil
}

func (a *oauthAdapter) ExistOauthClient(ctx context.Context, id string) (bool, error) {
	return a.oauthClientRepo.ExistOauthClient(ctx, id)
}

func (a *oauthAdapter) toOAuthClientEntity(m *dbmodels.OauthClient, redirectURLs []*dbmodels.OauthClientRedirectURL) *entities.OauthClient {
	return &entities.OauthClient{
		OauthClientID:      m.OauthClientID,
		ClientSecretHashed: m.ClientSecret,
		Name:               m.Name,
		Scope:              entities.OauthScope(m.Scope),
		ClientType:         entities.OauthClientType(m.ClientType),
		RedirectURLs: cast.ArrayValues(redirectURLs, func(t *dbmodels.OauthClientRedirectURL) string {
			return t.RedirectURL
		}),
	}
}

func (a *oauthAdapter) SetClientCredentialAccessToken(ctx context.Context, clientID, accessToken string, expireAt int64) error {
	m := &models.OauthAccessToken{
		ClientID:   clientID,
		ClientType: models.OauthClientTypeClientCredential,
		Scope:      models.OauthScopeAdmin,
	}
	return a.oauthAccessTokenRepo.SetClientCredentialAccessToken(ctx, m, accessToken, expireAt)
}

func (a *oauthAdapter) SetUserAccessToken(ctx context.Context, clientID, userID, accessToken string, scope entities.OauthScope, expireIn int64) error {
	m := &models.OauthAccessToken{
		ClientID:   clientID,
		ClientType: models.OauthClientTypeUserCredential,
		Scope:      models.OauthAccessTokenFromEntity(scope),
		UserID:     userID,
	}
	return a.oauthAccessTokenRepo.SetClientCredentialAccessToken(ctx, m, accessToken, expireIn)
}
func (a *oauthAdapter) GetSession(ctx context.Context, accessToken string) (*entities.OauthSession, error) {
	return a.oauthAccessTokenRepo.GetSession(ctx, accessToken)
}

func (a *oauthAdapter) SetCode(ctx context.Context, code *entities.OAuthCode) error {
	return a.oauthCodeRepo.SetCode(ctx, code.Code, &models.OauthCode{
		ClientID:    code.ClientID,
		UserID:      code.UserID,
		RedirectURL: code.RedirectURL,
	})
}

func (a *oauthAdapter) GetCode(ctx context.Context, code string) (*entities.OAuthCode, error) {
	m, err := a.oauthCodeRepo.GetCode(ctx, code)
	if err != nil {
		return nil, err
	}
	return &entities.OAuthCode{
		Code:        code,
		ClientID:    m.ClientID,
		UserID:      m.UserID,
		RedirectURL: m.RedirectURL,
	}, nil
}

func (a *oauthAdapter) UpsertUserAuth(ctx context.Context, userAuth *entities.UserAuth) (*entities.UserAuth, error) {
	m := &dbmodels.UserAuth{
		UserID:                  userAuth.UserID,
		OauthClientID:           userAuth.OAuthClientID,
		RefreshToken:            "",
		RefreshTokenPublishedAt: userAuth.RefreshTokenPublishedAt,
	}
	dst, err := a.userAuthRepository.UpsertUserAuth(ctx, m, userAuth.RefreshToken)
	if err != nil {
		return nil, err
	}
	return a.toUserAuthEntity(dst), nil
}

func (a *oauthAdapter) GetUserAuth(ctx context.Context, userID, clientID string) (*entities.UserAuth, error) {
	m, err := a.userAuthRepository.GetUserAuth(ctx, userID, clientID)
	if err != nil {
		return nil, err
	}
	return a.toUserAuthEntity(m), nil
}

func (a *oauthAdapter) GetUserAuthByRefreshToken(ctx context.Context, refreshToken string) (*entities.UserAuth, error) {
	m, err := a.userAuthRepository.GetUserAuthByRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	return a.toUserAuthEntity(m), nil
}
func (a *oauthAdapter) DeleteUserAuth(ctx context.Context, userID, clientID string) error {
	return a.userAuthRepository.DeleteUserAuth(ctx, userID, clientID)
}
func (a *oauthAdapter) DeleteClientAllAuth(ctx context.Context, clientID string) error {
	return a.userAuthRepository.DeleteClientAllAuth(ctx, clientID)
}

func (a *oauthAdapter) toUserAuthEntity(m *dbmodels.UserAuth) *entities.UserAuth {
	return &entities.UserAuth{
		UserID:                  m.UserID,
		OAuthClientID:           m.OauthClientID,
		RefreshToken:            m.RefreshToken,
		RefreshTokenPublishedAt: m.RefreshTokenPublishedAt,
	}
}

func (a *oauthAdapter) toRedirectURLEntity(url *dbmodels.OauthClientRedirectURL) *entities.OAuthClientRedirectURL {
	return &entities.OAuthClientRedirectURL{
		OAuthClientRedirectUrlID: url.OauthClientRedirectURLID,
		OauthClientID:            url.OauthClientID,
		RedirectURL:              url.RedirectURL,
	}
}
