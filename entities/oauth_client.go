package entities

type OauthClient struct {
	OauthClientID      string
	ClientSecretHashed string
	Name               string
	Scope              OauthScope
	ClientType         OauthClientType
	RedirectURLs       []string
}

type OauthScope string

const (
	OauthScopeUnknown OauthScope = "Unknown"
	OauthScopeGeneral OauthScope = "General"
	OauthScopeAdmin   OauthScope = "Admin"
)

type OauthClientType int

const (
	OauthClientTypeUnknown          OauthClientType = 0
	OauthClientTypeUserClient       OauthClientType = 1
	OauthClientTypeClientCredential OauthClientType = 2
)

func (t OauthClientType) String() string {
	switch t {
	case OauthClientTypeClientCredential:
		return "client_credentials"
	}
	return ""
}

type Oauth2ClientCredential struct {
	AccessToken string
	TokenType   OauthClientType
	ExpireIn    int
}
