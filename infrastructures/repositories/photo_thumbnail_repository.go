package repositories

import (
	"context"
	"fmt"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/drivers/mysql"
	"github.com/hiroyky/famiphoto/drivers/storage"
	"github.com/hiroyky/famiphoto/infrastructures/dbmodels"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"path"
)

type PhotoThumbnailRepository interface {
	SavePreview(ctx context.Context, photoID int64, data []byte, groupID, ownerID string) error
	SaveThumbnail(ctx context.Context, photoID int64, data []byte, groupID, ownerID string) error
}

func NewPhotoThumbnailRepository(fileDriver storage.Driver, db mysql.SQLExecutor) PhotoThumbnailRepository {
	return &photoThumbnailRepository{
		fileDriver: fileDriver,
		db:         db,
	}
}

type photoThumbnailRepository struct {
	fileDriver storage.Driver
	db         mysql.SQLExecutor
}

func (r *photoThumbnailRepository) SavePreview(ctx context.Context, photoID int64, data []byte, groupID, ownerID string) error {
	return r.saveImage(ctx, photoID, data, groupID, ownerID, config.AssetPreviewImageName)
}

func (r *photoThumbnailRepository) SaveThumbnail(ctx context.Context, photoID int64, data []byte, groupID, ownerID string) error {
	return r.saveImage(ctx, photoID, data, groupID, ownerID, config.AssetThumbnailImageName)
}

func (r *photoThumbnailRepository) saveImage(ctx context.Context, photoID int64, data []byte, groupID, ownerID, name string) error {
	m := &dbmodels.PhotoThumbnail{
		PhotoID:       int(photoID),
		ThumbnailName: name,
		FilePath:      r.genFilePath(groupID, ownerID, name, photoID),
		GroupID:       groupID,
		OwnerID:       ownerID,
	}

	if err := r.fileDriver.CreateFile(m.FilePath, data); err != nil {
		return err
	}

	if exist, err := dbmodels.PhotoThumbnailExists(ctx, r.db, int(photoID), name); err != nil {
		return err
	} else if exist {
		if _, err := m.Update(ctx, r.db, boil.Infer()); err != nil {
			return err
		}
		return nil
	}

	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func (r *photoThumbnailRepository) genFilePath(groupID, ownerID, thumbnailName string, photoID int64) string {
	return path.Join(groupID, ownerID, fmt.Sprintf("%d-%s.jpg", photoID, thumbnailName))
}
