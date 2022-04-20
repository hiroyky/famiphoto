package usecases

import (
	"fmt"
	"github.com/hiroyky/famiphoto/entities"
)

type PhotoImportUseCase interface {
	ImportPhotos(basePath string, extensions []string) error
}

func NewPhotoImportUseCase(photoService PhotoService, storage PhotoStorageAdapter) PhotoImportUseCase {
	return &photoImportUseCase{
		photoService: photoService,
		storage:      storage,
	}
}

type photoImportUseCase struct {
	photoService PhotoService
	storage      PhotoStorageAdapter
}

func (u *photoImportUseCase) ImportPhotos(basePath string, extensions []string) error {
	contents, err := u.storage.FindDirContents(basePath)
	if err != nil {
		return err
	}
	files := make([]*entities.StorageFile, 0)
	for _, c := range contents {
		if c.IsDir {
			if err := u.ImportPhotos(c.Path, extensions); err != nil {
				return err
			}
		} else if c.IsMatchExt(extensions) {
			files = append(files, c)
		}
	}

	for _, file := range files {
		fmt.Println(file)
		_, err := u.photoService.ParsePhotoExif(file.Path)
		if err != nil {
			return err
		}
	}

	return nil
}