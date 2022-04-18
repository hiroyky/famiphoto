package services

import (
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/usecases"
)

func NewPhotoService(photoStorage usecases.PhotoStorageAdapter) usecases.PhotoService {
	return &photoService{
		photoStorage: photoStorage,
	}
}

type photoService struct {
	photoStorage usecases.PhotoStorageAdapter
}

func (s *photoService) ParsePhotoExif(filePath string) (*entities.PhotoMeta, error) {
	_, err := s.photoStorage.ParsePhotoMeta(filePath)
	if err != nil {
		return nil, err
	}

	return nil, err
}
