package repositories

import (
	"context"
	"fmt"
	"github.com/hiroyky/famiphoto/infrastructures/dbmodels"
	"github.com/hiroyky/famiphoto/usecases"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"path"
)

func NewPhotoThumbnailRepository(fileDriver ThumbnailStorageAdapter, db SQLExecutor) usecases.PhotoThumbnailAdapter {
	return &photoThumbnailRepository{
		fileDriver: fileDriver,
		db:         db,
	}
}

type photoThumbnailRepository struct {
	fileDriver ThumbnailStorageAdapter
	db         SQLExecutor
}

func (r *photoThumbnailRepository) SavePreviewThumbnail(ctx context.Context, photoID int64, data []byte, groupID, ownerID string) error {
	name := "preview"

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
