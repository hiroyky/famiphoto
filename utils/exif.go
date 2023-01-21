package utils

import (
	"fmt"
	"github.com/dsoprea/go-exif"
	log "github.com/dsoprea/go-logging"
	"github.com/hiroyky/famiphoto/errors"
	"time"
)

func ParseDatetime(val string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation("2006:01:02 15:04:05", val, loc)
}

func ParseExifItemsAll(data []byte) (ExifItemList, error) {
	list := make([]*ExifItem, 0)

	rawExif, err := exif.SearchAndExtractExif(data)
	if err != nil {
		if errors.Is(err, exif.ErrNoExif) {
			return list, nil
		}
		return nil, err
	}

	im := exif.NewIfdMappingWithStandard()
	ti := exif.NewTagIndex()

	visitorFunc := func(fqIfdPath string, ifdIndex int, tagId uint16, tagType exif.TagType, valueContext exif.ValueContext) (err error) {
		item, err := parseVisitorFunc(im, ti, fqIfdPath, ifdIndex, tagId, tagType, valueContext)
		if err != nil {
			return err
		}
		if item == nil {
			return nil
		}
		list = append(list, item)
		return nil
	}

	if _, err := exif.Visit(exif.IfdStandard, im, ti, rawExif, visitorFunc); err != nil {
		return nil, err
	}

	return list, nil
}

func ParseExifItem(data []byte, exifTagID int) (*ExifItem, error) {
	rawExif, err := exif.SearchAndExtractExif(data)
	if err != nil {
		if errors.Is(err, exif.ErrNoExif) {
			return nil, errors.New(errors.NoExifError, err)
		}
		return nil, err
	}

	im := exif.NewIfdMappingWithStandard()
	ti := exif.NewTagIndex()

	var exifItem *ExifItem = nil
	var errExif error
	visitorFunc := func(fqIfdPath string, ifdIndex int, tagId uint16, tagType exif.TagType, valueContext exif.ValueContext) (err error) {
		if tagId != uint16(exifTagID) {
			return nil
		}
		exifItem, errExif = parseVisitorFunc(im, ti, fqIfdPath, ifdIndex, tagId, tagType, valueContext)
		return errExif
	}

	if _, err := exif.Visit(exif.IfdStandard, im, ti, rawExif, visitorFunc); err != nil {
		return nil, err
	}

	if exifItem == nil {
		return nil, errors.New(errors.NoExifError, nil)
	}
	return exifItem, nil
}

func parseVisitorFunc(im *exif.IfdMapping, ti *exif.TagIndex, fqIfdPath string, ifdIndex int, tagId uint16, tagType exif.TagType, valueContext exif.ValueContext) (*ExifItem, error) {
	ifdPath, err := im.StripPathPhraseIndices(fqIfdPath)
	if err != nil {
		return nil, err
	}

	it, err := ti.Get(ifdPath, tagId)
	if err != nil {
		if log.Is(err, exif.ErrTagNotFound) {
			return nil, nil
		}
		return nil, err
	}

	valueString := ""
	var value interface{}
	if tagType.Type() == exif.TypeUndefined {
		var err error
		value, err = valueContext.Undefined()
		if err != nil {
			if err == exif.ErrUnhandledUnknownTypedTag {
				value = nil
			} else {
				return nil, err
			}
		}

		valueString = fmt.Sprintf("%v", value)
	} else {
		valueString, err = valueContext.FormatFirst()
		//log.PanicIf(err)
		value = valueString
	}

	item := &ExifItem{
		IfdPath:     ifdPath,
		FqIfdPath:   fqIfdPath,
		IfdIndex:    ifdIndex,
		TagId:       tagId,
		TagName:     it.Name,
		TagTypeId:   tagType.Type(),
		TagTypeName: tagType.Name(),
		UnitCount:   valueContext.UnitCount(),
		Value:       value,
		ValueString: valueString,
	}
	return item, nil
}

type ExifItem struct {
	IfdPath     string
	FqIfdPath   string
	IfdIndex    int
	TagId       uint16
	TagName     string
	TagTypeId   exif.TagTypePrimitive
	TagTypeName string
	UnitCount   uint32
	Value       interface{}
	ValueString string
}

type ExifItemList []*ExifItem
