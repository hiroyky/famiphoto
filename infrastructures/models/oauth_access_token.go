package models

import (
	"encoding/json"
	"github.com/hiroyky/famiphoto/entities"
)

type OauthAccessToken struct {
	ClientID   string                `json:"client_id" validate:"required"`
	ClientType OauthAccessClientType `json:"client_type" validate:"required"`
	Scope      OauthAccessTokenScope `json:"scope" validate:"required"`
	UserID     string                `json:"user_id,omitempty"`
}

type OauthAccessClientType string

const (
	OauthClientTypeClientCredential OauthAccessClientType = "client_credential"
	OauthClientTypeUserCredential   OauthAccessClientType = "user_credential"
)

func (s OauthAccessClientType) Entity() entities.OauthClientType {
	switch s {
	case OauthClientTypeClientCredential:
		return entities.OauthClientTypeClientCredential
	case OauthClientTypeUserCredential:
		return entities.OauthClientTypeUserClient
	}
	panic("invalid client type: " + s)
}

type OauthAccessTokenScope string

const (
	OauthScopeAdmin OauthAccessTokenScope = "admin"
	OauthScopeUser  OauthAccessTokenScope = "user"
)

func (s OauthAccessTokenScope) Entity() entities.OauthScope {
	switch s {
	case OauthScopeAdmin:
		return entities.OauthScopeAdmin
	case OauthScopeUser:
		return entities.OauthScopeUser
	}
	panic("invalid scope: " + s)
}

func (t *OauthAccessToken) String() (string, error) {
	s, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(s), nil
}
