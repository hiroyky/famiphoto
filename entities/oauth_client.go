package entities

type OauthClient struct {
	OauthClientID string
	Name          string
	Scope         OauthScope
	ClientType    OauthClientType
	RedirectURLs  []string
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
