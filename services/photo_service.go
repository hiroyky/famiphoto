package services

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/infrastructures"
	"github.com/hiroyky/famiphoto/utils"
	"time"
)

type PhotoService interface {
	RegisterPhoto(ctx context.Context, filePath, fileHash, ownerID, groupID string) (*entities.Photo, error)
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

func (s *photoService) RegisterPhoto(ctx context.Context, filePath, fileHash, ownerID, groupID string) (*entities.Photo, error) {
	now := s.nowFunc()

	dstPhoto, err := s.photoAdapter.UpsertPhotoByFilePath(ctx, &entities.Photo{
		Name:         utils.FileNameExceptExt(filePath),
		ImportedAt:   now,
		GroupID:      groupID,
		OwnerID:      ownerID,
		FileNameHash: utils.MD5(utils.FileNameExceptExt(filePath)),
		Files: []*entities.PhotoFile{
			{
				FilePath:   filePath,
				ImportedAt: now,
				GroupID:    groupID,
				OwnerID:    ownerID,
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
