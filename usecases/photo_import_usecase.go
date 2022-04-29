package usecases

import (
	"context"
	"fmt"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/utils"
)

type PhotoImportUseCase interface {
	ImportPhotos(ctx context.Context, basePath string, extensions []string) error
}

func NewPhotoImportUseCase(
	photoService PhotoService,
	imageProcessService ImageProcessService,
	storage PhotoStorageAdapter,
) PhotoImportUseCase {
	return &photoImportUseCase{
		photoService:        photoService,
		imageProcessService: imageProcessService,
		storage:             storage,
	}
}

type photoImportUseCase struct {
	photoService        PhotoService
	imageProcessService ImageProcessService
	storage             PhotoStorageAdapter
}

func (u *photoImportUseCase) ImportPhotos(ctx context.Context, basePath string, extensions []string) error {
	groupID, ownerID, err := u.parseBasePath(basePath)
	if err != nil {
		return err
	}

	contents, err := u.storage.FindDirContents(basePath)
	if err != nil {
		return err
	}
	files := make([]*entities.StorageFileInfo, 0)
	for _, c := range contents {
		if c.IsDir {
			if err := u.ImportPhotos(ctx, c.Path, extensions); err != nil {
				return err
			}
		} else if c.IsMatchExt(extensions) {
			files = append(files, c)
		}
	}

	for _, file := range files {
		fmt.Println(file)
		data, err := u.storage.LoadContent(file.Path)
		if err != nil {
			return err
		}

		photoFile, err := u.photoService.RegisterPhoto(ctx, file.Path, data.FileHash(), ownerID, groupID)
		if err != nil {
			return err
		}

		// JPEGでなければサムネ画像の作成処理は行わないので終了
		if photoFile.FileType() == entities.PhotoFileTypeJPEG {
			if err := u.imageProcessService.CreateThumbnails(ctx, photoFile, data); err != nil {
				return err
			}
		}

	}

	return nil
}

func (u *photoImportUseCase) parseBasePath(basePath string) (string, string, error) {
	directories := utils.SplitPath(basePath)
	fmt.Println(directories)
	if len(directories) < 3 {
		return "", "", errors.New(errors.InvalidFilePathFatal, fmt.Errorf(basePath))
	}
	groupID := directories[1]
	ownerID := directories[2]
	return groupID, ownerID, nil
}
