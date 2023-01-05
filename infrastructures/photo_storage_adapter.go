package infrastructures

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/models"
	"github.com/hiroyky/famiphoto/infrastructures/repositories"
	"github.com/hiroyky/famiphoto/utils"
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
	SavePhotoFile(ctx context.Context, userID, groupID, fileName string, data []byte) (*entities.StorageFileInfo, error)
	GenerateSignToSavePhoto(ctx context.Context, userID, groupID string, expireIn int64) (string, error)
	VerifySignToken(ctx context.Context, token string) (*entities.PhotoUploadInfo, error)
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

func (a *photoStorageAdapter) SavePhotoFile(ctx context.Context, userID, groupID, fileName string, data []byte) (*entities.StorageFileInfo, error) {
	fileInfo, p, err := a.photoStorageRepo.SaveContent(groupID, userID, fileName, data)
	if err != nil {
		if errors.IsErrCode(err, errors.FileAlreadyExistError) {
			return a.SavePhotoFile(ctx, userID, groupID, utils.IncrementFileNameSuffix(fileName), data)
		}
		return nil, err
	}
	return &entities.StorageFileInfo{
		Name:  filepath.Base(fileInfo.Name()),
		Path:  p,
		Ext:   filepath.Ext(fileInfo.Name()),
		IsDir: fileInfo.IsDir(),
	}, nil
}

func (a *photoStorageAdapter) GenerateSignToSavePhoto(ctx context.Context, userID, groupID string, expireIn int64) (string, error) {
	sign := a.generateRandomStringFunc(16)
	if err := a.photoUploadSignRepo.SetSignToken(ctx, sign, userID, groupID, expireIn); err != nil {
		return "", err
	}
	return sign, nil
}

func (a *photoStorageAdapter) VerifySignToken(ctx context.Context, token string) (*entities.PhotoUploadInfo, error) {
	sign, err := a.photoUploadSignRepo.GetSign(ctx, token)
	if err != nil {
		if errors.GetFPErrorCode(err) == errors.PhotoUploadSignNotFoundError {
			return nil, errors.New(errors.ForbiddenError, err)
		}
		return nil, err
	}

	return &entities.PhotoUploadInfo{
		UserID:  sign.UserID,
		GroupID: sign.GroupID,
	}, nil
}
