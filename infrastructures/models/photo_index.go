package models

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type PhotoIndex struct {
	PhotoID          int64    `json:"photo_id"`
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

type PhotoResult struct {
	Total  int64
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
