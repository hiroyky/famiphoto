package model

import (
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/utils/array"
	"github.com/hiroyky/famiphoto/utils/gql"
	"path"
	"time"
)

const PhotoFileName = "PhotoFile"

func NewPhotoFiles(l entities.PhotoFileList) []*PhotoFile {
	return array.Map(l, NewPhotoFile)
}

func NewPhotoFile(e *entities.PhotoFile) *PhotoFile {
	return &PhotoFile{
		ID:         gql.EncodeIntID(PhotoFileName, e.PhotoFileID),
		PhotoID:    gql.EncodeIntID(PhotoName, e.PhotoID),
		FileType:   e.FileType().ToString(),
		ImportedAt: e.ImportedAt.Format(time.RFC3339),
		FileHash:   e.FileHash,
		FileName:   path.Base(e.FilePath),
	}
}
