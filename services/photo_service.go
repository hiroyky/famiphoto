package services

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/infrastructures"
	"github.com/hiroyky/famiphoto/utils"
	"time"
)

type PhotoService interface {
	ShouldSkipToRegisterPhoto(ctx context.Context, filePath, fileHash string) (bool, error)
	RegisterPhoto(ctx context.Context, filePath, fileHash string) (*entities.Photo, error)
}

func NewPhotoService(photoRepo infrastructures.PhotoAdapter, photoStorage infrastructures.PhotoStorageAdapter) PhotoService {
	return &photoService{
		photoAdapter:        photoRepo,
		photoStorageAdapter: photoStorage,
		nowFunc:             time.Now,
	}
}

type photoService struct {
	photoAdapter        infrastructures.PhotoAdapter
	photoStorageAdapter infrastructures.PhotoStorageAdapter
	nowFunc             func() time.Time
}

func (s *photoService) ShouldSkipToRegisterPhoto(ctx context.Context, filePath, fileHash string) (bool, error) {
	if exist, err := s.photoAdapter.ExistPhotoFileByFilePath(ctx, filePath); err != nil {
		return false, err
	} else if !exist {
		return false, nil
	}

	photoFile, err := s.photoAdapter.GetPhotoFileByFilePath(ctx, filePath)
	if err != nil {
		return false, err
	}

	return fileHash == photoFile.FileHash, nil
}
func (s *photoService) RegisterPhoto(ctx context.Context, filePath, fileHash string) (*entities.Photo, error) {
	now := s.nowFunc()

	dstPhoto, err := s.photoAdapter.UpsertPhotoByFilePath(ctx, &entities.Photo{
		Name:         utils.FileNameExceptExt(filePath),
		ImportedAt:   now,
		FileNameHash: utils.MD5(utils.FileNameExceptExt(filePath)),
		Files: []*entities.PhotoFile{
			{
				FilePath:   filePath,
				ImportedAt: now,
				FileHash:   fileHash,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	meta, err := s.photoStorageAdapter.ParsePhotoMeta(filePath)
	if err != nil {
		return nil, err
	}

	for _, item := range meta {
		if _, err := s.photoAdapter.UpsertPhotoMetaItemByPhotoTagID(ctx, dstPhoto.PhotoID, item); err != nil {
			return nil, err
		}
	}

	return dstPhoto, err
}
