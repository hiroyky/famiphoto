package usecases

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures"
)

type DownloadUseCase interface {
	VerifyDownloadPermission(ctx context.Context, fileID int, userID string) error
	LoadPhotoFile(ctx context.Context, fileID int) (entities.StorageFileData, *entities.PhotoFile, error)
}

func NewDownloadUseCase(
	photoAdapter infrastructures.PhotoAdapter,
	groupAdapter infrastructures.GroupAdapter,
	photoStorageAdapter infrastructures.PhotoStorageAdapter,
) DownloadUseCase {
	return &downloadUseCase{
		photoAdapter:        photoAdapter,
		groupAdapter:        groupAdapter,
		photoStorageAdapter: photoStorageAdapter,
	}
}

type downloadUseCase struct {
	photoAdapter        infrastructures.PhotoAdapter
	groupAdapter        infrastructures.GroupAdapter
	photoStorageAdapter infrastructures.PhotoStorageAdapter
}

func (u *downloadUseCase) VerifyDownloadPermission(ctx context.Context, fileID int, userID string) error {
	photoFile, err := u.photoAdapter.GetPhotoFileByPhotoFileID(ctx, fileID)
	if err != nil {
		return err
	}

	if belong, err := u.groupAdapter.IsBelongGroupUser(ctx, photoFile.GroupID, userID); err != nil {
		return err
	} else if !belong {
		return errors.New(errors.FileNotFoundError, nil)
	}

	return nil
}

func (u *downloadUseCase) LoadPhotoFile(ctx context.Context, fileID int) (entities.StorageFileData, *entities.PhotoFile, error) {
	photoFile, err := u.photoAdapter.GetPhotoFileByPhotoFileID(ctx, fileID)
	if err != nil {
		return nil, nil, err
	}

	data, err := u.photoStorageAdapter.LoadContent(photoFile.FilePath)
	if err != nil {
		return nil, nil, err
	}
	return data, photoFile, nil
}
