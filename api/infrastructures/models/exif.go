package models

import (
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	"github.com/hiroyky/famiphoto/entities"
)

type IfdEntry struct {
	IfdPath     string                      `json:"ifd_path"`
	TagId       uint16                      `json:"tag_id"`
	TagName     string                      `json:"tag_name"`
	TagTypeId   exifcommon.TagTypePrimitive `json:"tag_type_id"`
	TagTypeName string                      `json:"tag_type_name"`
	UnitCount   uint32                      `json:"unit_count"`
	Value       interface{}                 `json:"value"`
	ValueString string                      `json:"value_string"`
}

func (e IfdEntry) ToEntity() entities.PhotoMetaItem {
	return entities.PhotoMetaItem{}
}
