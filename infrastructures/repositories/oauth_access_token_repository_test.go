package repositories

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/hiroyky/famiphoto/infrastructures/models"
	mock_repositories "github.com/hiroyky/famiphoto/testing/mocks/infrastructures/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestOauthAccessTokenRepository_SetClientCredentialAccessToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	clientID := "client_id1"
	accessToken := "access_token"
	expireIn := int64(1000)
	content, _ := (&models.OauthAccessToken{
		ClientID:   clientID,
		ClientType: models.OauthClientTypeClientCredential,
		Scope:      models.OauthScopeAdmin,
	}).String()

	redisAdapter := mock_repositories.NewMockRedisAdapter(ctrl)
	redisAdapter.EXPECT().SetEx(gomock.Any(), gomock.Any(), content, time.Duration(expireIn)*time.Second)

	repo := oauthAccessTokenRepository{
		db:     redisAdapter,
		prefix: "prefix_",
	}

	err := repo.SetClientCredentialAccessToken(context.Background(), clientID, accessToken, expireIn)
	assert.NoError(t, err)
}
