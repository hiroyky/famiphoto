package entities

import (
	"sort"
	"time"
)

type Photo struct {
	PhotoID    int64
	Name       string
	FilePath   string
	ImportedAt time.Time
	GroupID    string
	OwnerID    string
}

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
