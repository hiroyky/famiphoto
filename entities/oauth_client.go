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
	OauthScopeUser    OauthScope = "General"
	OauthScopeAdmin   OauthScope = "Admin"
)

func (s OauthScope) String() string {
	return string(s)
}

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

type OauthSession struct {
	ClientType OauthClientType
	Scope      OauthScope
	ClientID   string
	UserID     string
}
