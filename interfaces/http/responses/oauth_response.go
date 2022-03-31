package responses

import (
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"net/url"
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

type OAuthAuthorizationCodeResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

func NewOAuthAuthorizationCodeResponse(code *entities.Oauth2AuthorizationCode) *OAuthAuthorizationCodeResponse {
	return &OAuthAuthorizationCodeResponse{
		AccessToken:  code.AccessToken,
		RefreshToken: code.RefreshToken,
		ExpiresIn:    code.ExpireIn,
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

func NewOAuthCodeRedirectURL(redirectURL, code, state string) (string, error) {
	u, err := url.Parse(redirectURL)
	if err != nil {
		return "", errors.New(errors.InvalidRequestError, err)
	}
	q := u.Query()
	q.Set("code", code)
	q.Set("state", state)
	u.RawQuery = q.Encode()
	return u.String(), nil
}
