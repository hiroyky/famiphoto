package usecases

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"time"
)

type RandomService interface {
	GenerateRandomString(length int) string
}

type UserAdapter interface {
	GetUser(ctx context.Context, userID string) (*entities.User, error)
	GetUsers(ctx context.Context, filter *UserFilter, limit, offset int) (entities.UserList, error)
	CountUsers(ctx context.Context, filter *UserFilter) (int, error)
	ExistUser(ctx context.Context, userID string) (bool, error)
	CreateUser(ctx context.Context, user *entities.User, password string, isInitializedPassword bool, now time.Time) (*entities.User, error)
}

type UserFilter struct {
	UserID *string
}

type UserService interface {
	AuthUserPassword(ctx context.Context, userID, password string) error
}

type UserPasswordAdapter interface {
	GetUserPassword(ctx context.Context, userID string) (*entities.UserPassword, error)
}

type PasswordService interface {
	HashPassword(password string) (string, error)
	MatchPassword(password string, hash string) (bool, error)
	GeneratePassword(length int) (string, error)
}

type OauthClientAdapter interface {
	GetByOauthClientID(ctx context.Context, id string) (*entities.OauthClient, error)
	CreateOAuthClient(ctx context.Context, client *entities.OauthClient, clientSecret string) (*entities.OauthClient, error)
	ExistOauthClient(ctx context.Context, id string) (bool, error)
}

type OauthClientRedirectURLAdapter interface {
	GetOAuthClientRedirectURLsByOAuthClientID(ctx context.Context, oauthClientID string) (entities.OAuthClientRedirectURLList, error)
	CreateOAuthClientRedirectURL(ctx context.Context, url *entities.OAuthClientRedirectURL) (*entities.OAuthClientRedirectURL, error)
}

type OauthAccessTokenAdapter interface {
	SetClientCredentialAccessToken(ctx context.Context, clientID, accessToken string, expireAt int64) error
	GetSession(ctx context.Context, accessToken string) (*entities.OauthSession, error)
}

type OauthCodeAdapter interface {
	SetCode(ctx context.Context, code *entities.OAuthCode) error
	GetCode(ctx context.Context, code string) (*entities.OAuthCode, error)
}
