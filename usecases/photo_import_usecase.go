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
	"os/exec"
	"path"
	"time"
)

type PhotoImportUseCase interface {
	GenerateUploadURL(ctx context.Context, userID, groupID string, now time.Time) (*entities.PhotoUploadSign, error)
	UploadPhoto(ctx context.Context, signToken, fileName string, body []byte) error
	IndexingPhotos(ctx context.Context, rootPath, groupID, userID string, extensions []string, fast bool) error
	ExecuteBatch(ctx context.Context, groupID, userID string, fast bool) error
}

func NewPhotoImportUseCase(
	photoService services.PhotoService,
	imageProcessService services.ImageProcessService,
	photoAdapter infrastructures.PhotoAdapter,
	storage infrastructures.PhotoStorageAdapter,
	searchAdapter infrastructures.SearchAdapter,
	userAdapter infrastructures.UserAdapter,
	groupAdapter infrastructures.GroupAdapter,
) PhotoImportUseCase {
	return &photoImportUseCase{
		photoService:        photoService,
		imageProcessService: imageProcessService,
		photoAdapter:        photoAdapter,
		storageAdapter:      storage,
		searchAdapter:       searchAdapter,
		groupAdapter:        groupAdapter,
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
	groupAdapter        infrastructures.GroupAdapter
	appendBulkUnit      int
	parseExifItemFunc   func(data []byte, tagID int) (*utils.ExifItem, error)
	nowFunc             func() time.Time
}

func (u *photoImportUseCase) GenerateUploadURL(ctx context.Context, userID, groupID string, now time.Time) (*entities.PhotoUploadSign, error) {
	if belonging, err := u.groupAdapter.IsBelongGroupUser(ctx, groupID, userID); err != nil {
		return nil, err
	} else if !belonging {
		return nil, errors.New(errors.ForbiddenError, nil)
	}

	token, err := u.storageAdapter.GenerateSignToSavePhoto(ctx, userID, groupID, config.PhotoUploadSignExpireInSec)
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

	if belonging, err := u.groupAdapter.IsBelongGroupUser(ctx, info.GroupID, info.UserID); err != nil {
		return err
	} else if !belonging {
		return errors.New(errors.ForbiddenError, nil)
	}

	dateTimeOriginal := utils.MustLocalTime(time.Now(), config.Env.ExifTimezone)
	exifDateTimeOriginal, err := u.parseExifItemFunc(body, config.ExifTagIDDateTimeOriginal)
	if err == nil {
		dt, err := utils.ParseDatetime(exifDateTimeOriginal.ValueString, utils.MustLoadLocation(config.Env.ExifTimezone))
		if err == nil {
			dateTimeOriginal = dt
		}
	}

	dstFile, err := u.storageAdapter.SavePhotoFile(ctx, info.UserID, info.GroupID, fileName, dateTimeOriginal, body)
	if err != nil {
		return err
	}
	photo, err := u.registerPhoto(ctx, dstFile, info.UserID, info.GroupID, false)
	if err != nil {
		return err
	}

	return u.searchAdapter.InsertPhoto(ctx, photo)
}

func (u *photoImportUseCase) IndexingPhotos(ctx context.Context, rootPath, groupID, userID string, extensions []string, fast bool) error {
	targetDirPath := path.Join(rootPath, groupID, userID)
	return u.importDirRecursive(ctx, targetDirPath, groupID, userID, extensions, fast)
}

func (u *photoImportUseCase) importDirRecursive(ctx context.Context, targetDirPath, groupID, userID string, extensions []string, fast bool) error {
	contents, err := u.storageAdapter.FindDirContents(targetDirPath)
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

	photo, err := u.photoService.RegisterPhoto(ctx, file.Path, data.FileHash(), ownerID, groupID)
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

func (u *photoImportUseCase) ExecuteBatch(ctx context.Context, groupID, userID string, fast bool) error {
	if belong, err := u.groupAdapter.IsBelongGroupUser(ctx, groupID, userID); err != nil {
		return err
	} else if !belong {
		return errors.New(errors.ForbiddenError, nil)
	}

	fastArg := "false"
	if fast {
		fastArg = "true"
	}

	cmd := exec.Command("dst/indexing_photos", "--user-id", userID, "--group-id", groupID, "--fast", fastArg)
	return cmd.Start()
}
