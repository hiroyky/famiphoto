package services

import (
	"context"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/infrastructures"
	"github.com/hiroyky/famiphoto/utils/image"
)

type ImageProcessService interface {
	CreateThumbnails(ctx context.Context, photoFile *entities.PhotoFile, data []byte, orientation int) error
}

func NewImageProcessService(thumbRepo infrastructures.PhotoStorageAdapter) ImageProcessService {
	return &imageProcessService{thumbRepo: thumbRepo}
}

type imageProcessService struct {
	thumbRepo infrastructures.PhotoStorageAdapter
}

func (s *imageProcessService) CreateThumbnails(ctx context.Context, photoFile *entities.PhotoFile, data []byte, orientation int) error {
	if err := s.createPreview(ctx, photoFile, data, orientation); err != nil {
		return err
	}
	if err := s.createThumbnail(ctx, photoFile, data, orientation); err != nil {
		return err
	}
	return nil
}

func (s *imageProcessService) createPreview(ctx context.Context, photoFile *entities.PhotoFile, data []byte, orientation int) error {
	dstData, err := s.resizeWidth(data, 1920)
	if err != nil {
		return err
	}

	dstData, err = s.rotateByOrientation(dstData, orientation)
	if err != nil {
		return err
	}

	return s.thumbRepo.SavePreview(ctx, photoFile.PhotoID, dstData)
}

func (s *imageProcessService) createThumbnail(ctx context.Context, photoFile *entities.PhotoFile, data []byte, orientation int) error {
	dstData, err := s.resizeHeight(data, 400)
	if err != nil {
		return err
	}

	dstData, err = s.rotateByOrientation(dstData, orientation)
	if err != nil {
		return err
	}

	return s.thumbRepo.SaveThumbnail(ctx, photoFile.PhotoID, dstData)
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

func (s *imageProcessService) rotateByOrientation(data []byte, orientation int) ([]byte, error) {
	switch orientation {
	case config.ExifOrientationHorizontal:
		return image.FlipHJPEG(data)
	case config.ExifOrientationRotate180:
		return image.Rotate180JPEG(data)
	case config.ExifOrientationVertical:
		return image.FlipVJPEG(data)
	case config.ExifOrientationHorizontalRotate270:
		dst, err := image.FlipHJPEG(data)
		if err != nil {
			return nil, err
		}
		return image.Rotate90JPEG(dst)
	case config.ExifOrientationRotate90:
		return image.Rotate270JPEG(data)
	case config.ExifOrientationHorizontalRotate90:
		dst, err := image.FlipHJPEG(data)
		if err != nil {
			return nil, err
		}
		return image.Rotate270JPEG(dst)
	case config.ExifOrientationRotate270:
		return image.Rotate90JPEG(data)
	}
	return data, nil
}
