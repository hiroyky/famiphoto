package model

import (
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/utils/gql"
)

const PhotoExifName = "PhotoExif"

func NewPhotoExifData(data entities.PhotoMeta) []*PhotoExif {
	dst := make([]*PhotoExif, len(data))
	for i, v := range data {
		dst[i] = newPhotoExifItem(v)
	}
	return dst
}

func newPhotoExifItem(i *entities.PhotoMetaItem) *PhotoExif {
	return &PhotoExif{
		ID:          gql.EncodeIntID(PhotoExifName, i.PhotoMetaItemID),
		TagID:       i.TagID,
		TagType:     i.TagType,
		ValueString: i.ValueString,
	}
}
