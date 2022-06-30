package repositories

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/drivers/redis"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/models"
	"time"
)

type OauthAccessTokenRepository interface {
	SetClientCredentialAccessToken(ctx context.Context, oauthAccessToken *models.OauthAccessToken, accessToken string, expireIn int64) error
	SetUserAccessToken(ctx context.Context, oauthAccessToken *models.OauthAccessToken, accessToken string, expireIn int64) error
	GetSession(ctx context.Context, accessToken string) (*entities.OauthSession, error)
}

func NewOauthAccessTokenRepository(db redis.Driver) OauthAccessTokenRepository {
	return &oauthAccessTokenRepository{
		db:     db,
		prefix: config.Env.AccessTokenHashedPrefix,
	}
}

type oauthAccessTokenRepository struct {
	db     redis.Driver
	prefix string
}

func (r *oauthAccessTokenRepository) SetClientCredentialAccessToken(ctx context.Context, oauthAccessToken *models.OauthAccessToken, accessToken string, expireIn int64) error {
	val, err := oauthAccessToken.String()
	if err != nil {
		return err
	}

	return r.db.SetEx(ctx, r.toHash(accessToken), val, time.Duration(expireIn)*time.Second)
}

func (r *oauthAccessTokenRepository) SetUserAccessToken(ctx context.Context, oauthAccessToken *models.OauthAccessToken, accessToken string, expireIn int64) error {
	val, err := oauthAccessToken.String()
	if err != nil {
		return err
	}
	return r.db.SetEx(ctx, r.toHash(accessToken), val, time.Duration(expireIn)*time.Second)
}

func (r *oauthAccessTokenRepository) GetSession(ctx context.Context, accessToken string) (*entities.OauthSession, error) {
	str, err := r.db.Get(ctx, r.toHash(accessToken))
	if err != nil {
		code := errors.GetFPErrorCode(err)
		if code == errors.RedisKeyNotFound {
			return nil, errors.New(errors.OAuthAccessTokenNotFoundError, nil)
		}
		return nil, err
	}
	var val models.OauthAccessToken
	if err := json.Unmarshal([]byte(str), &val); err != nil {
		return nil, err
	}

	dst := &entities.OauthSession{
		ClientType: val.ClientType.Entity(),
		ClientID:   val.ClientID,
		Scope:      val.Scope.Entity(),
		UserID:     val.UserID,
	}
	return dst, nil
}

func (r *oauthAccessTokenRepository) toHash(accessToken string) string {
	base := fmt.Sprintf("%s-%s", r.prefix, accessToken)

	hashed := sha256.Sum256([]byte(base))
	return base64.StdEncoding.EncodeToString(hashed[:])
}
