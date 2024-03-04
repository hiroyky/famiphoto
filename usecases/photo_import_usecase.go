package usecases

import (
	"context"
	"fmt"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/infrastructures"
	"github.com/hiroyky/famiphoto/services"
	"github.com/hiroyky/famiphoto/utils"
	"os/exec"
	"time"
)

type PhotoImportUseCase interface {
	GenerateUploadURL(ctx context.Context, userID string, now time.Time) (*entities.PhotoUploadSign, error)
	UploadPhoto(ctx context.Context, signToken, fileName string, body []byte) error
	IndexingPhotos(ctx context.Context, rootPath string, extensions []string, fast bool) error
	ExecuteBatch(ctx context.Context, fast bool) error
}

func NewPhotoImportUseCase(
	photoService services.PhotoService,
	imageProcessService services.ImageProcessService,
	photoAdapter infrastructures.PhotoAdapter,
	storage infrastructures.PhotoStorageAdapter,
	searchAdapter infrastructures.SearchAdapter,
	userAdapter infrastructures.UserAdapter,
) PhotoImportUseCase {
	return &photoImportUseCase{
		photoService:        photoService,
		imageProcessService: imageProcessService,
		photoAdapter:        photoAdapter,
		storageAdapter:      storage,
		searchAdapter:       searchAdapter,
		userAdapter:         userAdapter,
		appendBulkUnit:      20,
		parseExifItemFunc:   utils.ParseExifItem,
		nowFunc:             time.Now,
	}
}

type photoImportUseCase struct {
	photoService        services.PhotoService
	imageProcessService services.ImageProcessService
	photoAdapter        infrastructures.PhotoAdapter
	storageAdapter      infrastructures.PhotoStorageAdapter
	searchAdapter       infrastructures.SearchAdapter
	userAdapter         infrastructures.UserAdapter
	appendBulkUnit      int
	parseExifItemFunc   func(data []byte, tagID int) (*utils.ExifItem, error)
	nowFunc             func() time.Time
}

func (u *photoImportUseCase) GenerateUploadURL(ctx context.Context, userID string, now time.Time) (*entities.PhotoUploadSign, error) {
	token, err := u.storageAdapter.GenerateSignToSavePhoto(ctx, userID, config.PhotoUploadSignExpireInSec)
	if err != nil {
		return nil, err
	}
	return &entities.PhotoUploadSign{
		SignToken: token,
		ExpireAt:  config.PhotoUploadSignExpireInSec + int(now.Unix()),
	}, nil
}

func (u *photoImportUseCase) UploadPhoto(ctx context.Context, signToken, fileName string, body []byte) error {
	info, err := u.storageAdapter.VerifySignToken(ctx, signToken)
	if err != nil {
		return err
	}

	dateTimeOriginal := utils.MustLocalTime(time.Now(), config.Env.ExifTimezone)
	exifDateTimeOriginal, err := u.parseExifItemFunc(body, config.ExifTagIDDateTimeOriginal)
	if err == nil {
		dt, err := utils.ParseDatetime(exifDateTimeOriginal.ValueString, utils.MustLoadLocation(config.Env.ExifTimezone))
		if err == nil {
			dateTimeOriginal = dt
		}
	}

	dstFile, err := u.storageAdapter.SavePhotoFile(ctx, info.UserID, fileName, dateTimeOriginal, body)
	if err != nil {
		return err
	}
	photo, err := u.registerPhoto(ctx, dstFile, false)
	if err != nil {
		return err
	}

	return u.searchAdapter.InsertPhoto(ctx, photo)
}

func (u *photoImportUseCase) IndexingPhotos(ctx context.Context, rootPath string, extensions []string, fast bool) error {
	return u.importDirRecursive(ctx, rootPath, extensions, fast)
}

func (u *photoImportUseCase) importDirRecursive(ctx context.Context, targetDirPath string, extensions []string, fast bool) error {
	contents, err := u.storageAdapter.FindDirContents(targetDirPath)
	if err != nil {
		return err
	}

	files := make([]*entities.StorageFileInfo, 0)
	for _, c := range contents {
		if c.IsDir {
			if err := u.importDirRecursive(ctx, c.Path, extensions, fast); err != nil {
				return err
			}
		} else if c.IsMatchExt(extensions) {
			files = append(files, c)
		}
	}

	photoList := make(entities.PhotoList, 0)
	for _, file := range files {
		photo, err := u.registerPhoto(ctx, file, fast)
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

func (u *photoImportUseCase) registerPhoto(ctx context.Context, file *entities.StorageFileInfo, fast bool) (*entities.Photo, error) {
	data, err := u.storageAdapter.LoadContent(file.Path)
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

	photo, err := u.photoService.RegisterPhoto(ctx, file.Path, data.FileHash())
	if err != nil {
		return nil, err
	}
	var orientation = config.ExifOrientationNone
	orientationMeta, err := u.photoAdapter.GetPhotoMetaItemByPhotoIDTagID(ctx, photo.PhotoID, config.ExifTagOrientation)
	if err == nil {
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

func (u *photoImportUseCase) ExecuteBatch(ctx context.Context, fast bool) error {
	fastArg := "false"
	if fast {
		fastArg = "true"
	}

	cmd := exec.Command("dst/indexing_photos", "--fast", fastArg)
	return cmd.Start()
}
