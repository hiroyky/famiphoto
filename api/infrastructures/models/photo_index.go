package models

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/utils"
	"github.com/hiroyky/famiphoto/utils/array"
	"io"
	"strings"
	"time"
)

func NewPhotoIndex(p *entities.Photo, dateTimeOriginalEpoc int64) *PhotoIndex {
	return &PhotoIndex{
		PhotoID: p.PhotoID,
		FileTypes: array.Map(p.Files.FindFileTypesByPhotoID(p.PhotoID), func(t entities.PhotoFileType) string {
			return t.ToString()
		}),
		Name:             p.Name,
		ImportedAt:       p.ImportedAt.Unix(),
		DateTimeOriginal: dateTimeOriginalEpoc,
		PreviewURL:       p.PreviewURL(),
		ThumbnailURL:     p.ThumbnailURL(),
	}
}

type PhotoIndex struct {
	PhotoID          int      `json:"photo_id"`
	FileTypes        []string `json:"file_types"`
	Name             string   `json:"name"`
	ImportedAt       int64    `json:"imported_at"`
	DateTimeOriginal int64    `json:"date_time_original"`
	PreviewURL       string   `json:"preview_url"`
	ThumbnailURL     string   `json:"thumbnail_url"`
}

func (m *PhotoIndex) BodyReader() (io.Reader, error) {
	jsonData, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return strings.NewReader(string(jsonData)), nil
}

func (m *PhotoIndex) PhotoIndexID() string {
	id := fmt.Sprintf("Photo:%d", m.PhotoID)
	return base64.StdEncoding.EncodeToString([]byte(id))
}

func (m *PhotoIndex) ToEntityItem() *entities.PhotoSearchResultItem {
	return &entities.PhotoSearchResultItem{
		PhotoID:          m.PhotoID,
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

type PhotoDateHistogram struct {
	EpochSec int64
	DocCount int
}

func (m *PhotoDateHistogram) ToEntity(tz string) *entities.PhotoDateTimeAggregationItem {
	if m.EpochSec == 0 {
		return &entities.PhotoDateTimeAggregationItem{
			Year:  0,
			Month: 0,
			Date:  0,
			Num:   m.DocCount,
		}
	}
	tm := utils.MustLocalTime(time.Unix(m.EpochSec, 0), tz)
	return &entities.PhotoDateTimeAggregationItem{
		Year:  tm.Year(),
		Month: int(tm.Month()),
		Date:  tm.Day(),
		Num:   m.DocCount,
	}
}
