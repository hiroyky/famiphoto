package services

import (
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/usecases"
)

func NewImageProcessService(thumbRepo usecases.PhotoThumbnailAdapter) usecases.ImageProcessService {
	return &imageProcessService{thumbRepo: thumbRepo}
}

type imageProcessService struct {
	thumbRepo usecases.PhotoThumbnailAdapter
}

func (s *imageProcessService) CreateThumbnails(photoFile *entities.PhotoFile, data []byte) error {
	return nil
}
