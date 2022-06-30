package entities

import (
	"fmt"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/utils"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Photo struct {
	PhotoID      int64
	Name         string
	ImportedAt   time.Time
	GroupID      string
	OwnerID      string
	FileNameHash string
	Files        PhotoFileList
}

func (e Photo) HasJpeg() bool {
	return e.Files.FindFileByFileType(e.PhotoID, PhotoFileTypeJPEG) != nil
}

func (e Photo) PreviewURL() string {
	if !e.HasJpeg() {
		return ""
	}
	return fmt.Sprintf("%s/thumbnails/%s/%s/%d-%s.jpg", utils.RemoveTrailingSlash(config.Env.AssetBaseURL), e.GroupID, e.OwnerID, e.PhotoID, config.AssetPreviewImageName)
}

func (e Photo) ThumbnailURL() string {
	if !e.HasJpeg() {
		return ""
	}
	return fmt.Sprintf("%s/thumbnails/%s/%s/%d-%s.jpg", utils.RemoveTrailingSlash(config.Env.AssetBaseURL), e.GroupID, e.OwnerID, e.PhotoID, config.AssetThumbnailImageName)
}

type PhotoList []*Photo

func (l PhotoList) PhotoIDs() []int64 {
	idList := make([]int64, len(l))
	for _, p := range l {
		idList = append(idList, p.PhotoID)
	}
	return idList
}

type PhotoFile struct {
	PhotoFileID int64
	PhotoID     int64
	FilePath    string
	ImportedAt  time.Time
	GroupID     string
	OwnerID     string
	FileHash    string
}

type PhotoFileList []*PhotoFile

func (list PhotoFileList) FindFileTypesByPhotoID(photoID int64) []PhotoFileType {
	types := make([]PhotoFileType, 0)

	for _, item := range list {
		if item.PhotoID != photoID {
			continue
		}
		types = append(types, item.FileType())
	}

	return types
}

func (list PhotoFileList) FindFileByFileType(photoID int64, fileType PhotoFileType) *PhotoFile {
	for _, item := range list {
		if item.PhotoID != photoID {
			continue
		}
		if item.FileType() != fileType {
			continue
		}
		return item
	}
	return nil
}

func (f PhotoFile) FileType() PhotoFileType {
	ext := filepath.Ext(f.FilePath)
	switch strings.ToLower(ext) {
	case ".jpg":
		return PhotoFileTypeJPEG
	case ".arw":
		return PhotoFileTypeRAW
	}
	return PhotoFileTypeUnknown
}

type PhotoFileType string

func (t PhotoFileType) ToString() string {
	return string(t)
}

const (
	PhotoFileTypeJPEG    = "jpeg"
	PhotoFileTypeRAW     = "raw"
	PhotoFileTypeUnknown = "unknown"
)

type PhotoMetaItem struct {
	PhotoMetaItemID int64
	TagID           int64
	TagName         string
	TagType         string
	ValueString     string
}

func (i PhotoMetaItem) sortOrder() int64 {
	return i.TagID
}

type PhotoMeta []*PhotoMetaItem

func (m PhotoMeta) Sort() PhotoMeta {
	sort.Slice(m, func(i, j int) bool {
		return m[i].sortOrder() < m[j].sortOrder()
	})
	return m
}
