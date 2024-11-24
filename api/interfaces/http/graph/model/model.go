package model

import "github.com/hiroyky/famiphoto/entities"

type Photo struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	ImportedAt       string   `json:"importedAt"`
	DateTimeOriginal string   `json:"dateTimeOriginal"`
	PreviewURL       string   `json:"previewUrl"`
	ThumbnailURL     string   `json:"thumbnailUrl"`
	FileTypes        []string `json:"fileTypes"`
}

func (Photo) IsNode()         {}
func (m Photo) GetID() string { return m.GetID() }

type PhotoFile struct {
	ID         string `json:"id"`
	PhotoID    string `json:"photoId"`
	FileType   string `json:"fileType"`
	ImportedAt string `json:"importedAt"`
	FileHash   string `json:"fileHash"`
	FileName   string `json:"fileName"`
}

func (PhotoFile) IsNode()         {}
func (m PhotoFile) GetID() string { return m.GetID() }

type User struct {
	ID     string     `json:"id"`
	UserID string     `json:"userId"`
	Name   string     `json:"name"`
	Status UserStatus `json:"status"`
}

func (User) IsNode()         {}
func (m User) GetID() string { return m.GetID() }

type Group struct {
	ID      string `json:"id"`
	GroupID string `json:"groupId"`
	Name    string `json:"name"`
}

func (Group) IsNode()         {}
func (m Group) GetID() string { return m.GetID() }

type OauthClient struct {
	ID           string           `json:"id"`
	ClientID     string           `json:"clientId"`
	Name         string           `json:"name"`
	Scope        OauthClientScope `json:"scope"`
	ClientType   OauthClientType  `json:"clientType"`
	ClientSecret *string          `json:"clientSecret"`
}

func (OauthClient) IsNode()         {}
func (m OauthClient) GetID() string { return m.GetID() }

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
