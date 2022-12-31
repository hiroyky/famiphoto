package infrastructures

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/models"
	"github.com/hiroyky/famiphoto/infrastructures/repositories"
	"github.com/hiroyky/famiphoto/utils/array"
	"github.com/hiroyky/famiphoto/utils/random"
	"path"
	"path/filepath"
)

type PhotoStorageAdapter interface {
	FindDirContents(dirPath string) ([]*entities.StorageFileInfo, error)
	LoadContent(path string) (entities.StorageFileData, error)
	ParsePhotoMeta(path string) (entities.PhotoMeta, error)
	SavePreview(ctx context.Context, photoID int, data []byte, groupID, ownerID string) error
	SaveThumbnail(ctx context.Context, photoID int, data []byte, groupID, ownerID string) error
	GenerateSignToSavePhoto(ctx context.Context, userID, groupID string, expireIn int64) (string, error)
	VerifySignToken(ctx context.Context, token, userID, groupID string) error
}

func NewPhotoStorageAdapter(
	photoStorageRepo repositories.PhotoStorageRepository,
	thumbnailRepo repositories.PhotoThumbnailRepository,
	photoUploadSignRepo repositories.PhotoUploadSignRepository,
) PhotoStorageAdapter {
	return &photoStorageAdapter{
		photoStorageRepo:         photoStorageRepo,
		thumbnailRepo:            thumbnailRepo,
		photoUploadSignRepo:      photoUploadSignRepo,
		generateRandomStringFunc: random.GenerateRandomString,
	}
}

type photoStorageAdapter struct {
	photoStorageRepo         repositories.PhotoStorageRepository
	thumbnailRepo            repositories.PhotoThumbnailRepository
	photoUploadSignRepo      repositories.PhotoUploadSignRepository
	generateRandomStringFunc func(length int) string
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
			TagID:       int(t.TagId),
			TagName:     t.TagName,
			TagType:     t.TagTypeName,
			ValueString: t.ValueString,
		}
	})

	return photoMeta, err
}

func (a *photoStorageAdapter) SavePreview(ctx context.Context, photoID int, data []byte, groupID, ownerID string) error {
	return a.thumbnailRepo.SavePreview(ctx, photoID, data, groupID, ownerID)
}

func (a *photoStorageAdapter) SaveThumbnail(ctx context.Context, photoID int, data []byte, groupID, ownerID string) error {
	return a.thumbnailRepo.SaveThumbnail(ctx, photoID, data, groupID, ownerID)
}

func (a *photoStorageAdapter) GenerateSignToSavePhoto(ctx context.Context, userID, groupID string, expireIn int64) (string, error) {
	sign := a.generateRandomStringFunc(16)
	if err := a.photoUploadSignRepo.SetSignToken(ctx, sign, userID, groupID, expireIn); err != nil {
		return "", err
	}
	return sign, nil
}

func (a *photoStorageAdapter) VerifySignToken(ctx context.Context, token, userID, groupID string) error {
	sign, err := a.photoUploadSignRepo.GetSign(ctx, token)
	if err != nil {
		if errors.GetFPErrorCode(err) == errors.PhotoUploadSignNotFoundError {
			return errors.New(errors.ForbiddenError, err)
		}
		return err
	}

	if sign.UserID != userID {
		return errors.New(errors.ForbiddenError, nil)
	}
	if sign.GroupID != groupID {
		return errors.New(errors.ForbiddenError, nil)
	}

	return nil
}
