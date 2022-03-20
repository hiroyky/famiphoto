package models

import "encoding/json"

type OauthAccessToken struct {
	ClientID   string `json:"client_id" validate:"required"`
	ClientType string `json:"client_type" validate:"required"`
	UserID     string `json:"user_id,omitempty"`
}

const OauthClientTypeClientCredential = "client_credential"
const OauthClientTypeUserCredential = "user_credential"

func (t *OauthAccessToken) String() (string, error) {
	s, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(s), nil
}
