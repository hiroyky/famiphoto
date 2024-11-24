package usecases

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/infrastructures"
)

type PhotoUseCase interface {
	GetPhotoMetaByPhotoID(ctx context.Context, photoID int) (entities.PhotoMeta, error)
	GetPhotoFileByPhotoFileID(ctx context.Context, photoFileID int) (*entities.PhotoFile, error)
	GetPhotoFilesByPhotoID(ctx context.Context, photoID int) (entities.PhotoFileList, error)
}

func NewPhotoUseCase(photoAdapter infrastructures.PhotoAdapter) PhotoUseCase {
	return &photoUseCase{
		photoAdapter: photoAdapter,
	}
}

type photoUseCase struct {
	photoAdapter infrastructures.PhotoAdapter
}

func (u *photoUseCase) GetPhotoMetaByPhotoID(ctx context.Context, photoID int) (entities.PhotoMeta, error) {
	return u.photoAdapter.GetPhotoMetaByPhotoID(ctx, photoID)
}

func (u *photoUseCase) GetPhotoFileByPhotoFileID(ctx context.Context, photoFileID int) (*entities.PhotoFile, error) {
	return u.photoAdapter.GetPhotoFileByPhotoFileID(ctx, photoFileID)
}

func (u *photoUseCase) GetPhotoFilesByPhotoID(ctx context.Context, photoID int) (entities.PhotoFileList, error) {
	return u.photoAdapter.GetPhotoFilesByPhotoID(ctx, photoID)
}
