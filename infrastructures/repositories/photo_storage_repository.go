package repositories

import (
	"fmt"
	"github.com/dsoprea/go-exif"
	log "github.com/dsoprea/go-logging"
	"github.com/hiroyky/famiphoto/drivers/storage"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/models"
	"os"
	"path"
)

type PhotoStorageRepository interface {
	SaveContent(groupID, userID, fileName string, data []byte) (os.FileInfo, string, error)
	ReadDir(dirPath string) ([]os.FileInfo, error)
	LoadContent(path string) ([]byte, error)
	ParsePhotoMeta(path string) ([]models.IfdEntry, error)
	CreateGroupUserDir(groupID, userID string) error
}

func NewPhotoStorageRepository(driver storage.Driver) PhotoStorageRepository {
	return &photoStorageRepository{driver: driver}
}

type photoStorageRepository struct {
	driver storage.Driver
}

func (r *photoStorageRepository) SaveContent(groupID, userID, fileName string, data []byte) (os.FileInfo, string, error) {
	p := path.Join(groupID, userID, fileName)
	if exist := r.driver.Exist(p); exist {
		return nil, "", errors.New(errors.FileAlreadyExistError, nil)
	}

	if err := r.driver.CreateFile(p, data); err != nil {
		return nil, "", err
	}

	info, err := r.driver.Stat(p)
	if err != nil {
		return nil, "", err
	}
	return info, p, nil
}

func (r *photoStorageRepository) CreateGroupUserDir(groupID, userID string) error {
	if err := r.createDirIfNotExist(groupID); err != nil {
		return err
	}
	if err := r.createDirIfNotExist(path.Join(groupID, userID)); err != nil {
		return err
	}
	return nil
}

func (r *photoStorageRepository) createDirIfNotExist(p string) error {
	stat, err := r.driver.Stat(p)
	if err != nil && !errors.IsErrCode(err, errors.FileNotFoundError) {
		return err
	}
	if stat == nil {
		return r.driver.CreateDir(p, os.ModePerm)
	}
	if !stat.IsDir() {
		return errors.New(errors.UnExpectedFileAlreadyExistError, nil)
	}
	return nil
}

func (r *photoStorageRepository) ReadDir(dirPath string) ([]os.FileInfo, error) {
	return r.driver.ReadDir(dirPath)
}

func (r *photoStorageRepository) LoadContent(path string) ([]byte, error) {
	if exist := r.driver.Exist(path); !exist {
		return nil, errors.New(errors.FileNotFoundError, fmt.Errorf(path))
	}
	return r.driver.ReadFile(path)
}

func (r *photoStorageRepository) ParsePhotoMeta(path string) ([]models.IfdEntry, error) {
	data, err := r.LoadContent(path)
	if err != nil {
		return nil, err
	}

	ifdList, err := r.parseExif(data)
	if err != nil {
		return nil, err
	}

	return ifdList, nil
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
