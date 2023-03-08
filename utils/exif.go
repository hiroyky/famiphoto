package utils

import (
	"github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	"github.com/hiroyky/famiphoto/errors"
	"time"
)

func ParseDatetime(val string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation("2006:01:02 15:04:05", val, loc)
}

func ParseExifItemsAll(data []byte) (ExifItemList, error) {
	rawExif, err := exif.SearchAndExtractExif(data)
	if err != nil {
		if errors.Is(err, exif.ErrNoExif) {
			return make([]*ExifItem, 0), nil
		}
		return nil, err
	}

	entries, _, err := exif.GetFlatExifDataUniversalSearch(rawExif, nil, true)
	if err != nil {
		return nil, err
	}

	list := make([]*ExifItem, len(entries))
	for i, entry := range entries {
		list[i] = &ExifItem{
			IfdPath:     entry.IfdPath,
			TagId:       entry.TagId,
			TagName:     entry.TagName,
			TagTypeId:   entry.TagTypeId,
			TagTypeName: entry.TagTypeName,
			UnitCount:   entry.UnitCount,
			Value:       entry.Value,
			ValueString: entry.Formatted,
		}
	}

	return list, nil
}

func ParseExifItem(data []byte, exifTagID int) (*ExifItem, error) {
	list, err := ParseExifItemsAll(data)
	if err != nil {
		return nil, err
	}

	for _, item := range list {
		if item.TagId == uint16(exifTagID) {
			return item, nil
		}
	}

	return nil, errors.New(errors.NoExifError, nil)
}

type ExifItem struct {
	IfdPath     string
	TagId       uint16
	TagName     string
	TagTypeId   exifcommon.TagTypePrimitive
	TagTypeName string
	UnitCount   uint32
	Value       interface{}
	ValueString string
}

type ExifItemList []*ExifItem
