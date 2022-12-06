package usecases

import (
	"context"
	"fmt"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/infrastructures"
	"github.com/hiroyky/famiphoto/services"
	"path"
)

type PhotoImportUseCase interface {
	IndexingPhotos(ctx context.Context, rootPath, groupID, userID string, extensions []string, fast bool) error
}

func NewPhotoImportUseCase(
	photoService services.PhotoService,
	imageProcessService services.ImageProcessService,
	photoAdapter infrastructures.PhotoAdapter,
	storage infrastructures.PhotoStorageAdapter,
	searchAdapter infrastructures.SearchAdapter,
) PhotoImportUseCase {
	return &photoImportUseCase{
		photoService:        photoService,
		imageProcessService: imageProcessService,
		photoAdapter:        photoAdapter,
		storage:             storage,
		searchAdapter:       searchAdapter,
		appendBulkUnit:      20,
	}
}

type photoImportUseCase struct {
	photoService        services.PhotoService
	imageProcessService services.ImageProcessService
	photoAdapter        infrastructures.PhotoAdapter
	storage             infrastructures.PhotoStorageAdapter
	searchAdapter       infrastructures.SearchAdapter
	appendBulkUnit      int
}

func (u *photoImportUseCase) IndexingPhotos(ctx context.Context, rootPath, groupID, userID string, extensions []string, fast bool) error {
	targetDirPath := path.Join(rootPath, groupID, userID)
	return u.importDirRecursive(ctx, targetDirPath, groupID, userID, extensions, fast)
}

func (u *photoImportUseCase) importDirRecursive(ctx context.Context, targetDirPath, groupID, userID string, extensions []string, fast bool) error {
	contents, err := u.storage.FindDirContents(targetDirPath)
	if err != nil {
		return err
	}

	files := make([]*entities.StorageFileInfo, 0)
	for _, c := range contents {
		if c.IsDir {
			if err := u.importDirRecursive(ctx, c.Path, groupID, userID, extensions, fast); err != nil {
				return err
			}
		} else if c.IsMatchExt(extensions) {
			files = append(files, c)
		}
	}

	photoList := make(entities.PhotoList, 0)
	for _, file := range files {
		photo, err := u.registerPhoto(ctx, file, userID, groupID, fast)
		if err != nil {
			return err
		}
		if photo != nil {
			photoList = append(photoList, photo)
		}
		if len(photoList) > u.appendBulkUnit {
			if err := u.buildInsertSearchEngine(ctx, photoList); err != nil {
				return err
			}
			photoList = make(entities.PhotoList, 0)
		}
	}

	return u.buildInsertSearchEngine(ctx, photoList)
}

func (u *photoImportUseCase) buildInsertSearchEngine(ctx context.Context, photoList entities.PhotoList) error {
	if photoList == nil || len(photoList) == 0 {
		return nil
	}
	stats, err := u.searchAdapter.BulkInsertPhotos(ctx, photoList)
	if err != nil {
		return err
	}

	fmt.Printf(
		"NumAdded: %d, NumFlushed: %d, NumFailed: %d, NumIndex:%d, NumCreated: %d, NumUpdated:%d, NumDeleted: %d, NumRequest:%d\n",
		stats.NumAdded,
		stats.NumFlushed,
		stats.NumFailed,
		stats.NumIndexed,
		stats.NumCreated,
		stats.NumUpdated,
		stats.NumDeleted,
		stats.NumDeleted,
	)
	return nil
}

func (u *photoImportUseCase) registerPhoto(ctx context.Context, file *entities.StorageFileInfo, ownerID, groupID string, fast bool) (*entities.Photo, error) {
	data, err := u.storage.LoadContent(file.Path)
	if err != nil {
		return nil, err
	}

	if fast {
		if skip, err := u.photoService.ShouldSkipToRegisterPhoto(ctx, file.Path, data.FileHash()); err != nil {
			return nil, err
		} else if skip {
			return nil, nil
		}
	}

	photo, err := u.photoService.RegisterPhoto(ctx, file.Path, data.FileHash(), ownerID, groupID)
	if err != nil {
		return nil, err
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
				return nil, err
			}
		}
	}
	return photo, nil
}
