package usecases

import (
	"context"
	"fmt"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures"
	"github.com/hiroyky/famiphoto/services"
	"github.com/hiroyky/famiphoto/utils"
)

type PhotoImportUseCase interface {
	ImportPhotos(ctx context.Context, basePath string, extensions []string) error
}

func NewPhotoImportUseCase(
	photoService services.PhotoService,
	imageProcessService services.ImageProcessService,
	photoAdapter infrastructures.PhotoAdapter,
	storage infrastructures.PhotoStorageAdapter,
) PhotoImportUseCase {
	return &photoImportUseCase{
		photoService:        photoService,
		imageProcessService: imageProcessService,
		photoAdapter:        photoAdapter,
		storage:             storage,
	}
}

type photoImportUseCase struct {
	photoService        services.PhotoService
	imageProcessService services.ImageProcessService
	photoAdapter        infrastructures.PhotoAdapter
	storage             infrastructures.PhotoStorageAdapter
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
		if err := u.registerPhoto(ctx, file, ownerID, groupID); err != nil {
			return err
		}
		fmt.Println(file.Path)
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

func (u *photoImportUseCase) registerPhoto(ctx context.Context, file *entities.StorageFileInfo, ownerID, groupID string) error {
	data, err := u.storage.LoadContent(file.Path)
	if err != nil {
		return err
	}

	photo, err := u.photoService.RegisterPhoto(ctx, file.Path, data.FileHash(), ownerID, groupID)
	if err != nil {
		return err
	}

	var orientation = config.ExifOrientationNone
	orientationMeta, err := u.photoAdapter.GetPhotoMetaItemByPhotoIDTagID(ctx, photo.PhotoID, config.ExifTagOrientation)
	if err == nil {
		fmt.Println(photo.Name, orientationMeta.ValueString)
		orientation = orientationMeta.ValueInt()
	}

	for _, photoFile := range photo.Files {
		if photoFile.FileType() == entities.PhotoFileTypeJPEG {
			if err := u.imageProcessService.CreateThumbnails(ctx, photoFile, data, orientation); err != nil {
				return err
			}
		}
	}
	return nil
}
