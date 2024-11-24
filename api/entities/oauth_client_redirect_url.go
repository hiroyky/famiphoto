package entities

type OAuthClientRedirectURL struct {
	OAuthClientRedirectUrlID int
	OauthClientID            string
	RedirectURL              string
}

type OAuthClientRedirectURLList []*OAuthClientRedirectURL

func (l OAuthClientRedirectURLList) IsMatchURL(u string) bool {
	for _, v := range l {
		if u == v.RedirectURL {
			return true
		}
	}
	return false
}

type OAuthCode struct {
	Code        string
	ClientID    string
	UserID      string
	RedirectURL string
}
