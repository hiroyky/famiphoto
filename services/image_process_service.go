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
	return s.createPreview(ctx, photoFile, data)
}

func (s *imageProcessService) createPreview(ctx context.Context, photoFile *entities.PhotoFile, data []byte) error {
	srcWidth, srcHeight, err := image.GetSize(data)
	if err != nil {
		return err
	}

	dstWidth := int64(1920)
	thumbData := data
	if dstWidth <= srcWidth {
		dstHeight := image.CalcToResizeWidth(srcWidth, srcHeight, dstWidth)
		thumbData, err = image.ResizeJPEG(data, dstWidth, dstHeight)
		if err != nil {
			return err
		}
	}

	return s.thumbRepo.SavePreviewThumbnail(ctx, photoFile.PhotoID, thumbData, photoFile.GroupID, photoFile.OwnerID)
}
