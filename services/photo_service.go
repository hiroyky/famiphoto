package services

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/usecases"
	"path/filepath"
	"time"
)

func NewPhotoService(photoStorage usecases.PhotoStorageAdapter) usecases.PhotoService {
	return &photoService{
		photoStorage: photoStorage,
	}
}

type photoService struct {
	photoRepo    usecases.PhotoAdapter
	photoStorage usecases.PhotoStorageAdapter
	nowFunc      func() time.Time
}

func (s *photoService) RegisterPhoto(ctx context.Context, filePath, ownerID, groupID string) error {
	photo, err := s.insertPhotoIfNotExist(ctx, filePath, ownerID, groupID)
	if err != nil {
		return err
	}

	meta, err := s.photoStorage.ParsePhotoMeta(filePath)
	if err != nil {
		return err
	}

	for _, item := range meta {
		if err := s.upsertPhotoMetaItem(ctx, photo.PhotoID, item); err != nil {
			return err
		}
	}

	return err
}

func (s *photoService) insertPhotoIfNotExist(ctx context.Context, filePath, ownerID, groupID string) (*entities.Photo, error) {
	existPhoto, err := s.photoRepo.GetPhotoByFilePath(ctx, filePath)
	if err == nil && existPhoto != nil {
		return existPhoto, nil
	}
	return s.photoRepo.InsertPhoto(
		ctx,
		&entities.Photo{
			Name:     filepath.Base(filePath),
			FilePath: filePath,
			GroupID:  groupID,
			OwnerID:  ownerID,
		},
	)
}

func (s *photoService) upsertPhotoMetaItem(ctx context.Context, photoID int64, metaItem *entities.PhotoMetaItem) error {
	existTag, err := s.photoRepo.GetPhotoMetaItemByTagID(ctx, photoID, metaItem.TagID)
	if err == nil && existTag != nil {
		if _, err := s.photoRepo.UpdatePhotoMetaItem(ctx, photoID, metaItem); err != nil {
			return err
		}
	}

	if _, err := s.photoRepo.InsertPhotoMetaItem(ctx, photoID, metaItem); err != nil {
		return err
	}

	return nil
}
