package model

import "github.com/hiroyky/famiphoto/entities"

type Photo struct {
	ID               string   `json:"id"`
	OwnerID          string   `json:"ownerId"`
	GroupID          string   `json:"groupId"`
	Name             string   `json:"name"`
	ImportedAt       string   `json:"importedAt"`
	DateTimeOriginal string   `json:"dateTimeOriginal"`
	PreviewURL       string   `json:"previewUrl"`
	ThumbnailURL     string   `json:"thumbnailUrl"`
	FileTypes        []string `json:"fileTypes"`
}

func (Photo) IsNode() {}

type PhotoFile struct {
	ID         string `json:"id"`
	PhotoID    string `json:"photoId"`
	FileType   string `json:"fileType"`
	ImportedAt string `json:"importedAt"`
	GroupID    string `json:"groupId"`
	OwnerID    string `json:"ownerId"`
	FileHash   string `json:"fileHash"`
}

func (PhotoFile) IsNode() {}

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
