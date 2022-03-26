package responses

import (
	"github.com/hiroyky/famiphoto/entities"
)

type OauthAccessTokenResponse struct {
	AccessToken  string `json:"access_token,omitempty"`
	Token        string `json:"token,omitempty"`
	ExpireIn     int    `json:"expire_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func NewOauthAccessTokenFromClientCredential(cc *entities.Oauth2ClientCredential) *OauthAccessTokenResponse {
	return &OauthAccessTokenResponse{
		AccessToken: cc.AccessToken,
		Token:       cc.TokenType.String(),
		ExpireIn:    cc.ExpireIn,
	}
}

func NewAuthorizePage(csrf, redirectURL, state, scope string, client *entities.OauthClient) map[string]interface{} {
	return map[string]interface{}{
		"csrf":         csrf,
		"client_id":    client.OauthClientID,
		"client_name":  client.Name,
		"redirect_url": redirectURL,
		"state":        state,
		"scope":        scope,
	}
}
