package model

import (
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/utils/array"
	"github.com/hiroyky/famiphoto/utils/gql"
	"time"
)

const PhotoFileName = "PhotoFile"

func NewPhotoFiles(l entities.PhotoFileList) []*PhotoFile {
	return array.Map(l, NewPhotoFile)
}

func NewPhotoFile(e *entities.PhotoFile) *PhotoFile {
	return &PhotoFile{
		ID:          gql.EncodeIntID(PhotoFileName, e.PhotoFileID),
		PhotoID:     gql.EncodeIntID(PhotoName, e.PhotoID),
		FileType:    e.FileType().ToString(),
		DownloadURL: "",
		ImportedAt:  e.ImportedAt.Format(time.RFC3339),
		GroupID:     gql.EncodeStrID(GroupName, e.GroupID),
		OwnerID:     gql.EncodeStrID(UserName, e.OwnerID),
		FileHash:    e.FileHash,
	}
}
