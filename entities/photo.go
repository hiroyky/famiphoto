package entities

import (
	"github.com/hiroyky/famiphoto/utils"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Photo struct {
	PhotoID    int64
	Name       string
	ImportedAt time.Time
	GroupID    string
	OwnerID    string
	FilePath   string
}

func (e Photo) FileNameHash() string {
	p := path.Join(filepath.Dir(e.FilePath), utils.FileNameExceptExt(e.FilePath))
	return utils.MD5(p)
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
