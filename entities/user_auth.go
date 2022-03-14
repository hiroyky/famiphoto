package entities

type UserAuth struct {
	UserID                  string
	OAuthClientID           string
	RefreshToken            string
	RefreshTokenPublishedAt int64
}
