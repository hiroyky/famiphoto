package services

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/usecases"
	"github.com/hiroyky/famiphoto/utils"
	"time"
)

func NewPhotoService(photoRepo usecases.PhotoAdapter, photoStorage usecases.PhotoStorageAdapter) usecases.PhotoService {
	return &photoService{
		photoRepo:    photoRepo,
		photoStorage: photoStorage,
		nowFunc:      time.Now,
	}
}

type photoService struct {
	photoRepo    usecases.PhotoAdapter
	photoStorage usecases.PhotoStorageAdapter
	nowFunc      func() time.Time
}

func (s *photoService) RegisterPhoto(ctx context.Context, filePath, fileHash, ownerID, groupID string) error {
	now := s.nowFunc()

	photo, err := s.insertPhotoIfNotExist(ctx, filePath, ownerID, groupID, now)
	if err != nil {
		return err
	}

	if err := s.upsertPhotoFile(ctx, photo, filePath, fileHash); err != nil {
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

func (s *photoService) insertPhotoIfNotExist(ctx context.Context, filePath, ownerID, groupID string, now time.Time) (*entities.Photo, error) {
	existPhoto, err := s.photoRepo.GetPhotoByFilePath(ctx, filePath)
	if err == nil && existPhoto != nil {
		return existPhoto, nil
	}
	return s.photoRepo.InsertPhoto(
		ctx,
		&entities.Photo{
			Name:       utils.FileNameExceptExt(filePath),
			GroupID:    groupID,
			OwnerID:    ownerID,
			ImportedAt: now,
			FilePath:   filePath,
		},
	)
}

func (s *photoService) upsertPhotoFile(ctx context.Context, photo *entities.Photo, filePath, fileHash string) error {
	photoFile := &entities.PhotoFile{
		PhotoFileID: 0,
		PhotoID:     photo.PhotoID,
		FilePath:    filePath,
		ImportedAt:  photo.ImportedAt,
		GroupID:     photo.GroupID,
		OwnerID:     photo.OwnerID,
		FileHash:    fileHash,
	}

	existPhotoFile, err := s.photoRepo.GetPhotoFileByFilePath(ctx, filePath)
	if err != nil && !errors.IsErrCode(err, errors.DBColumnNotFoundError) {
		return err
	}

	if err == nil && existPhotoFile != nil {
		photoFile.PhotoFileID = existPhotoFile.PhotoFileID
		if _, err := s.photoRepo.UpdatePhotoFile(ctx, photoFile); err != nil {
			return err
		}
	}

	if _, err := s.photoRepo.InsertPhotoFile(ctx, photoFile); err != nil {
		return err
	}

	return nil
}

func (s *photoService) upsertPhotoMetaItem(ctx context.Context, photoID int64, metaItem *entities.PhotoMetaItem) error {
	existTag, err := s.photoRepo.GetPhotoMetaItemByTagID(ctx, photoID, metaItem.TagID)
	if err == nil && existTag != nil {
		if _, err := s.photoRepo.UpdatePhotoMetaItem(ctx, photoID, metaItem); err != nil {
			return err
		}
		return nil
	}

	if _, err := s.photoRepo.InsertPhotoMetaItem(ctx, photoID, metaItem); err != nil {
		return err
	}

	return nil
}
