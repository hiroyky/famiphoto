package model

import "github.com/hiroyky/famiphoto/entities"

type Photo struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	FilePath   string `json:"filePath"`
	ImportedAt string `json:"importedAt"`
	GroupID    string `json:"groupId"`
	OwnerID    string `json:"ownerId"`
}

func (Photo) IsNode() {}

type User struct {
	ID     string     `json:"id"`
	Name   string     `json:"name"`
	Status UserStatus `json:"status"`
}

func (User) IsNode() {}

type Group struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (Group) IsNode() {}

type OauthClient struct {
	ID           string           `json:"id"`
	ClientID     string           `json:"clientId"`
	Name         string           `json:"name"`
	Scope        OauthClientScope `json:"scope"`
	ClientType   OauthClientType  `json:"clientType"`
	ClientSecret *string          `json:"clientSecret"`
}

func (OauthClient) IsNode() {}

func (e OauthClientScope) ToEntity() entities.OauthScope {
	switch e {
	case OauthClientScopeAdmin:
		return entities.OauthScopeAdmin
	case OauthClientScopeGeneral:
		return entities.OauthScopeUser
	}
	return entities.OauthScopeUnknown
}

func (e OauthClientType) ToEntity() entities.OauthClientType {
	switch e {
	case OauthClientTypeUserClient:
		return entities.OauthClientTypeUserClient
	case OauthClientTypeClientCredential:
		return entities.OauthClientTypeClientCredential
	}
	return entities.OauthClientTypeUnknown
}
