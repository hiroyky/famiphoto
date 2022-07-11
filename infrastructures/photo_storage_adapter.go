package infrastructures

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/infrastructures/models"
	"github.com/hiroyky/famiphoto/infrastructures/repositories"
	"github.com/hiroyky/famiphoto/utils/array"
	"path"
	"path/filepath"
)

type PhotoStorageAdapter interface {
	FindDirContents(dirPath string) ([]*entities.StorageFileInfo, error)
	LoadContent(path string) (entities.StorageFileData, error)
	ParsePhotoMeta(path string) (entities.PhotoMeta, error)
	SavePreview(ctx context.Context, photoID int64, data []byte, groupID, ownerID string) error
	SaveThumbnail(ctx context.Context, photoID int64, data []byte, groupID, ownerID string) error
}

func NewPhotoStorageAdapter(
	photoStorageRepo repositories.PhotoStorageRepository,
	thumbnailRepo repositories.PhotoThumbnailRepository,
) PhotoStorageAdapter {
	return &photoStorageAdapter{
		photoStorageRepo: photoStorageRepo,
		thumbnailRepo:    thumbnailRepo,
	}
}

type photoStorageAdapter struct {
	photoStorageRepo repositories.PhotoStorageRepository
	thumbnailRepo    repositories.PhotoThumbnailRepository
}

func (a *photoStorageAdapter) FindDirContents(dirPath string) ([]*entities.StorageFileInfo, error) {
	list, err := a.photoStorageRepo.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	files := make([]*entities.StorageFileInfo, len(list))
	for i, v := range list {
		files[i] = &entities.StorageFileInfo{
			Name:  v.Name(),
			Path:  path.Join(dirPath, v.Name()),
			Ext:   filepath.Ext(v.Name()),
			IsDir: v.IsDir(),
		}
	}
	return files, nil
}

func (a *photoStorageAdapter) LoadContent(path string) (entities.StorageFileData, error) {
	return a.photoStorageRepo.LoadContent(path)
}

func (a *photoStorageAdapter) ParsePhotoMeta(path string) (entities.PhotoMeta, error) {
	ifdList, err := a.photoStorageRepo.ParsePhotoMeta(path)
	if err != nil {
		return nil, err
	}

	photoMeta := array.Map(ifdList, func(t models.IfdEntry) *entities.PhotoMetaItem {
		return &entities.PhotoMetaItem{
			TagID:       int64(t.TagId),
			TagName:     t.TagName,
			TagType:     t.TagTypeName,
			ValueString: t.ValueString,
		}
	})

	return photoMeta, err
}

func (a *photoStorageAdapter) SavePreview(ctx context.Context, photoID int64, data []byte, groupID, ownerID string) error {
	return a.thumbnailRepo.SavePreview(ctx, photoID, data, groupID, ownerID)
}

func (a *photoStorageAdapter) SaveThumbnail(ctx context.Context, photoID int64, data []byte, groupID, ownerID string) error {
	return a.thumbnailRepo.SaveThumbnail(ctx, photoID, data, groupID, ownerID)
}
