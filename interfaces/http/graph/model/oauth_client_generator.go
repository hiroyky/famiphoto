package model

import (
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/utils/gql"
)

const OauthClientName = "OauthClient"

func NewOauthClientWithSecret(oauth *entities.OauthClient, secret string) *OauthClient {
	var scope OauthClientScope
	switch oauth.Scope {
	case entities.OauthScopeGeneral:
		scope = OauthClientScopeGeneral
	case entities.OauthScopeAdmin:
		scope = OauthClientScopeAdmin
	}

	var clientType OauthClientType
	switch oauth.ClientType {
	case entities.OauthClientTypeUserClient:
		clientType = OauthClientTypeUserClient
	case entities.OauthClientTypeClientCredential:
		clientType = OauthClientTypeClientCredential
	}

	return &OauthClient{
		ID:           gql.EncodeStrID(OauthClientName, oauth.OauthClientID),
		Name:         oauth.Name,
		Scope:        scope,
		ClientType:   clientType,
		ClientSecret: &secret,
	}
}
