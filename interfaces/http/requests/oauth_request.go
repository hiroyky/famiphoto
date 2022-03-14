package requests

type OauthGrantTokenRequest struct {
	GrantType    string `form:"grant_type" validator:"required,oneof=client_credentials authorization_code refresh_token"`
	ClientID     string `form:"client_id" validator:"required"`
	ClientSecret string `form:"client_secret" validator:"required"`
	Scope        string `json:"scope" form:"scope" validator:"required"`
	Code         string `json:"code" form:"code"`
	RedirectURL  string `form:"redirect_url"`
	RefreshToken string `form:"refresh_token"`
}
