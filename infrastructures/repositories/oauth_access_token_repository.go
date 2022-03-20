package repositories

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/hiroyky/famiphoto/config"
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
	}).String()
	if err != nil {
		return err
	}

	return r.db.SetEx(ctx, r.toHash(accessToken), val, time.Duration(expireIn)*time.Second)
}

func (r *oauthAccessTokenRepository) toHash(accessToken string) string {
	base := fmt.Sprintf("%s-%s", r.prefix, accessToken)

	hashed := sha256.Sum256([]byte(base))
	return base64.StdEncoding.EncodeToString(hashed[:])
}
