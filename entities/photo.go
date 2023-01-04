package entities

import (
	"fmt"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/utils"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Photo struct {
	PhotoID      int
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

func (l PhotoList) PhotoIDs() []int {
	idList := make([]int, len(l))
	for _, p := range l {
		idList = append(idList, p.PhotoID)
	}
	return idList
}

type PhotoFile struct {
	PhotoFileID int
	PhotoID     int
	FilePath    string
	ImportedAt  time.Time
	GroupID     string
	OwnerID     string
	FileHash    string
}

func (f PhotoFile) FileType() PhotoFileType {
	ext := filepath.Ext(f.FilePath)
	switch strings.ToLower(ext) {
	case ".jpeg", ".jpg":
		return PhotoFileTypeJPEG
	case ".arw":
		return PhotoFileTypeRAW
	}
	return PhotoFileTypeUnknown
}

func (f PhotoFile) MimeType() string {
	switch f.FileType() {
	case PhotoFileTypeJPEG:
		return "image/jpeg"
	case PhotoFileTypeRAW:
		return "image/x-dcraw"
	default:
		return "application/octet-stream"
	}
}

type PhotoFileList []*PhotoFile

func (list PhotoFileList) FindFileTypesByPhotoID(photoID int) []PhotoFileType {
	types := make([]PhotoFileType, 0)

	for _, item := range list {
		if item.PhotoID != photoID {
			continue
		}
		types = append(types, item.FileType())
	}

	return types
}

func (list PhotoFileList) FindFileByFileType(photoID int, fileType PhotoFileType) *PhotoFile {
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

func (list PhotoFileList) FileTypes() []PhotoFileType {
	dst := make([]PhotoFileType, len(list))
	for i, v := range list {
		dst[i] = v.FileType()
	}
	return dst
}

type PhotoFileType string

func (t PhotoFileType) ToString() string {
	return string(t)
}

type PhotoFileTypeList []PhotoFileType

func (l PhotoFileTypeList) ToStrings() []string {
	dst := make([]string, len(l))
	for i, v := range l {
		dst[i] = v.ToString()
	}
	return dst
}

const (
	PhotoFileTypeJPEG    PhotoFileType = "jpeg"
	PhotoFileTypeRAW     PhotoFileType = "raw"
	PhotoFileTypeUnknown PhotoFileType = "unknown"
)

type PhotoMetaItem struct {
	PhotoMetaItemID int
	TagID           int
	TagName         string
	TagType         string
	ValueString     string
}

func (i PhotoMetaItem) ValueInt() int {
	val, err := strconv.Atoi(i.ValueString)
	if err != nil {
		return 0
	}
	return val
}

func (i PhotoMetaItem) sortOrder() int {
	return i.TagID
}

type PhotoMeta []*PhotoMetaItem

func (m PhotoMeta) Sort() PhotoMeta {
	sort.Slice(m, func(i, j int) bool {
		return m[i].sortOrder() < m[j].sortOrder()
	})
	return m
}

type PhotoSearchResultItem struct {
	PhotoID          int
	OwnerID          string
	GroupID          string
	FileTypes        []string
	Name             string
	ImportedAt       time.Time
	DateTimeOriginal time.Time
	PreviewURL       string
	ThumbnailURL     string
}

type PhotoSearchResult struct {
	Items []*PhotoSearchResultItem
	Total int
}
