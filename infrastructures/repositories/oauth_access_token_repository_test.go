package repositories

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/hiroyky/famiphoto/infrastructures/models"
	mock_redis "github.com/hiroyky/famiphoto/testing/mocks/drivers/redis"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestOauthAccessTokenRepository_SetAccessToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	clientID := "client_id1"
	accessToken := "access_token"
	expireIn := int64(1000)
	token := &models.OauthAccessToken{
		ClientID:   clientID,
		ClientType: models.OauthClientTypeClientCredential,
		Scope:      models.OauthScopeAdmin,
	}
	content, _ := (token).String()

	redisAdapter := mock_redis.NewMockDriver(ctrl)
	redisAdapter.EXPECT().SetEx(gomock.Any(), gomock.Any(), content, time.Duration(expireIn)*time.Second)

	repo := oauthAccessTokenRepository{
		db:     redisAdapter,
		prefix: "prefix_",
	}

	err := repo.SetAccessToken(context.Background(), token, accessToken, expireIn)
	assert.NoError(t, err)
}
