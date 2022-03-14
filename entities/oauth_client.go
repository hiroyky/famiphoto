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
	OauthScopeGeneral OauthScope = "General"
	OauthScopeAdmin   OauthScope = "Admin"
)

type OauthClientType int

const (
	OauthClientTypeUserClient OauthClientType = 1
	OauthClientTypeAdmin      OauthClientType = 2
)
