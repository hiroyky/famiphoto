package repositories

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/models"
	"github.com/hiroyky/famiphoto/usecases"
	"time"
)

func NewOauthAccessTokenRepository(db RedisAdapter) usecases.OauthAccessTokenAdapter {
	return &oauthAccessTokenRepository{
		db:     db,
		prefix: config.Env.AccessTokenHashedPrefix,
	}
}

type oauthAccessTokenRepository struct {
	db     RedisAdapter
	prefix string
}

func (r *oauthAccessTokenRepository) SetClientCredentialAccessToken(ctx context.Context, clientID, accessToken string, expireIn int64) error {
	val, err := (&models.OauthAccessToken{
		ClientID:   clientID,
		ClientType: models.OauthClientTypeClientCredential,
		Scope:      models.OauthScopeAdmin,
	}).String()
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
