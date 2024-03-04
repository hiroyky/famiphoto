package usecases

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/infrastructures"
)

type DownloadUseCase interface {
	LoadPhotoFile(ctx context.Context, fileID int) (entities.StorageFileData, *entities.PhotoFile, error)
}

func NewDownloadUseCase(
	photoAdapter infrastructures.PhotoAdapter,
	photoStorageAdapter infrastructures.PhotoStorageAdapter,
) DownloadUseCase {
	return &downloadUseCase{
		photoAdapter:        photoAdapter,
		photoStorageAdapter: photoStorageAdapter,
	}
}

type downloadUseCase struct {
	photoAdapter        infrastructures.PhotoAdapter
	photoStorageAdapter infrastructures.PhotoStorageAdapter
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
