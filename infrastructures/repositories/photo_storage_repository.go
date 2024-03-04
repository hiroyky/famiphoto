package repositories

import (
	"fmt"
	"github.com/hiroyky/famiphoto/drivers/storage"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/models"
	"github.com/hiroyky/famiphoto/utils"
	"github.com/hiroyky/famiphoto/utils/array"
	"os"
	"path"
	"time"
)

type PhotoStorageRepository interface {
	SaveContent(userID, fileName string, dateTimeOriginal time.Time, data []byte) (os.FileInfo, string, error)
	ReadDir(dirPath string) ([]os.FileInfo, error)
	LoadContent(path string) ([]byte, error)
	ParsePhotoMetaFromFile(path string) ([]models.IfdEntry, error)
	CreateUserDir(userID string) error
}

func NewPhotoStorageRepository(driver storage.Driver) PhotoStorageRepository {
	return &photoStorageRepository{driver: driver}
}

type photoStorageRepository struct {
	driver storage.Driver
}

func (r *photoStorageRepository) SaveContent(userID, fileName string, dateTimeOriginal time.Time, data []byte) (os.FileInfo, string, error) {
	dateTimeOriginalName := fmt.Sprintf("%04d-%02d-%02d", dateTimeOriginal.Year(), dateTimeOriginal.Month(), dateTimeOriginal.Day())
	if err := r.createDirIfNotExist(path.Join(userID, dateTimeOriginalName)); err != nil {
		return nil, "", err
	}
	p := path.Join(userID, dateTimeOriginalName, fileName)
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

func (r *photoStorageRepository) CreateUserDir(userID string) error {
	if err := r.createDirIfNotExist(userID); err != nil {
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

func (r *photoStorageRepository) ParsePhotoMetaFromFile(path string) ([]models.IfdEntry, error) {
	data, err := r.LoadContent(path)
	if err != nil {
		return nil, err
	}

	list, err := utils.ParseExifItemsAll(data)
	if err != nil {
		return nil, err
	}

	ifdList := array.Map(list, func(i *utils.ExifItem) models.IfdEntry {
		return models.IfdEntry{
			IfdPath:     i.IfdPath,
			TagId:       i.TagId,
			TagName:     i.TagName,
			TagTypeId:   i.TagTypeId,
			TagTypeName: i.TagTypeName,
			UnitCount:   i.UnitCount,
			Value:       i.Value,
			ValueString: i.ValueString,
		}
	})

	return ifdList, nil
}
