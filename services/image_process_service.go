package services

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/usecases"
	"github.com/hiroyky/famiphoto/utils/image"
)

func NewImageProcessService(thumbRepo usecases.PhotoThumbnailAdapter) usecases.ImageProcessService {
	return &imageProcessService{thumbRepo: thumbRepo}
}

type imageProcessService struct {
	thumbRepo usecases.PhotoThumbnailAdapter
}

func (s *imageProcessService) CreateThumbnails(ctx context.Context, photoFile *entities.PhotoFile, data []byte) error {
	if err := s.createPreview(ctx, photoFile, data); err != nil {
		return err
	}
	if err := s.createThumbnail(ctx, photoFile, data); err != nil {
		return err
	}
	return nil
}

func (s *imageProcessService) createPreview(ctx context.Context, photoFile *entities.PhotoFile, data []byte) error {
	dstData, err := s.resizeWidth(data, 1920)
	if err != nil {
		return err
	}

	return s.thumbRepo.SavePreview(ctx, photoFile.PhotoID, dstData, photoFile.GroupID, photoFile.OwnerID)
}

func (s *imageProcessService) createThumbnail(ctx context.Context, photoFile *entities.PhotoFile, data []byte) error {
	dstData, err := s.resizeHeight(data, 400)
	if err != nil {
		return err
	}

	return s.thumbRepo.SaveThumbnail(ctx, photoFile.PhotoID, dstData, photoFile.GroupID, photoFile.OwnerID)
}

func (s *imageProcessService) resizeWidth(data []byte, dstWidth int64) ([]byte, error) {
	srcWidth, srcHeight, err := image.GetSize(data)
	if err != nil {
		return nil, err
	}

	thumbData := data
	if dstWidth <= srcWidth {
		dstHeight := image.CalcToResizeWidth(srcWidth, srcHeight, dstWidth)
		thumbData, err = image.ResizeJPEG(data, dstWidth, dstHeight)
		if err != nil {
			return nil, err
		}
	}

	return thumbData, nil
}

func (s *imageProcessService) resizeHeight(data []byte, dstHeight int64) ([]byte, error) {
	srcWidth, srcHeight, err := image.GetSize(data)
	if err != nil {
		return nil, err
	}

	thumbData := data
	if dstHeight <= dstHeight {
		dstWidth := image.CalcToResizeHeight(srcWidth, srcHeight, dstHeight)
		thumbData, err = image.ResizeJPEG(data, dstWidth, dstHeight)
		if err != nil {
			return nil, err
		}
	}

	return thumbData, nil
}
