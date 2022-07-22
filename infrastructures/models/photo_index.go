package models

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/hiroyky/famiphoto/entities"
	"time"
)

type PhotoIndex struct {
	PhotoID          int      `json:"photo_id"`
	OwnerID          string   `json:"owner_id"`
	GroupID          string   `json:"group_id"`
	FileTypes        []string `json:"file_types"`
	Name             string   `json:"name"`
	ImportedAt       int64    `json:"imported_at"`
	DateTimeOriginal int64    `json:"date_time_original"`
	PreviewURL       string   `json:"preview_url"`
	ThumbnailURL     string   `json:"thumbnail_url"`
}

func (m PhotoIndex) PhotoIndexID() string {
	id := fmt.Sprintf("Photo:%d", m.PhotoID)
	return base64.StdEncoding.EncodeToString([]byte(id))
}

func (m *PhotoIndex) ToEntityItem() *entities.PhotoSearchResultItem {
	return &entities.PhotoSearchResultItem{
		PhotoID:          m.PhotoID,
		OwnerID:          m.OwnerID,
		GroupID:          m.GroupID,
		FileTypes:        m.FileTypes,
		Name:             m.Name,
		ImportedAt:       time.Unix(m.ImportedAt, 0),
		DateTimeOriginal: time.Unix(m.DateTimeOriginal, 0),
		PreviewURL:       m.PreviewURL,
		ThumbnailURL:     m.ThumbnailURL,
	}
}

type PhotoResult struct {
	Total  int
	Photos []*PhotoIndex
}

func NewPhotoIndexFromMap(v map[string]any) (*PhotoIndex, error) {
	body, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	var p PhotoIndex
	if err := json.Unmarshal(body, &p); err != nil {
		return nil, err
	}
	return &p, nil
}
