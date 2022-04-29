package repositories

import (
	"fmt"
	"github.com/dsoprea/go-exif"
	log "github.com/dsoprea/go-logging"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/models"
	"github.com/hiroyky/famiphoto/usecases"
	"github.com/hiroyky/famiphoto/utils/array"
	"path"
	"path/filepath"
)

func NewPhotoStorageRepository(driver StorageAdapter) usecases.PhotoStorageAdapter {
	return &photoStorageRepository{
		driver: driver,
	}
}

type photoStorageRepository struct {
	driver StorageAdapter
}

func (r *photoStorageRepository) FindDirContents(dirPath string) ([]*entities.StorageFileInfo, error) {
	list, err := r.driver.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	files := make([]*entities.StorageFileInfo, len(list))
	for i, v := range list {
		files[i] = &entities.StorageFileInfo{
			Name:  v.Name(),
			Path:  path.Join(dirPath, v.Name()),
			Ext:   filepath.Ext(v.Name()),
			IsDir: v.IsDir(),
		}
	}
	return files, nil
}

func (r *photoStorageRepository) LoadContent(path string) (entities.StorageFileData, error) {
	if exist := r.driver.Exist(path); !exist {
		return nil, errors.New(errors.FileNotFound, fmt.Errorf(path))
	}
	return r.driver.ReadFile(path)
}

func (r *photoStorageRepository) ParsePhotoMeta(path string) (entities.PhotoMeta, error) {
	data, err := r.LoadContent(path)
	if err != nil {
		return nil, err
	}

	ifdList, err := r.parseExif(data)
	if err != nil {
		return nil, err
	}

	photoMeta := array.Map(ifdList, func(t models.IfdEntry) *entities.PhotoMetaItem {
		return &entities.PhotoMetaItem{
			TagID:       int64(t.TagId),
			TagName:     t.TagName,
			TagType:     t.TagTypeName,
			ValueString: t.ValueString,
		}
	})

	return photoMeta, err
}

func (r *photoStorageRepository) parseExif(data []byte) ([]models.IfdEntry, error) {
	rawExif, err := exif.SearchAndExtractExif(data)
	if err != nil {
		if errors.Is(err, exif.ErrNoExif) {
			return make([]models.IfdEntry, 0), nil
		}
		return nil, err
	}

	im := exif.NewIfdMappingWithStandard()
	ti := exif.NewTagIndex()

	entries := make([]models.IfdEntry, 0)
	visitorFunc := func(fqIfdPath string, ifdIndex int, tagId uint16, tagType exif.TagType, valueContext exif.ValueContext) (err error) {
		ifdPath, err := im.StripPathPhraseIndices(fqIfdPath)
		if err != nil {
			return err
		}

		it, err := ti.Get(ifdPath, tagId)
		if err != nil {
			if log.Is(err, exif.ErrTagNotFound) {
				fmt.Printf("WARNING: Unknown tag: [%s] (%04x)\n", ifdPath, tagId)
				return nil
			}
			return err
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
					return err
				}
			}

			valueString = fmt.Sprintf("%v", value)
		} else {
			valueString, err = valueContext.FormatFirst()
			//log.PanicIf(err)

			value = valueString
		}

		entry := models.IfdEntry{
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

		entries = append(entries, entry)

		return nil
	}

	if _, err := exif.Visit(exif.IfdStandard, im, ti, rawExif, visitorFunc); err != nil {
		return nil, err
	}
	return entries, nil
}
