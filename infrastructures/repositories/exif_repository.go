package repositories

import (
	"context"
	"database/sql"
	"github.com/hiroyky/famiphoto/drivers/mysql"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/dbmodels"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ExifRepository interface {
	GetPhotoMetaItemByTagID(ctx context.Context, photoID, tagID int64) (*dbmodels.Exif, error)
	InsertPhotoMetaItem(ctx context.Context, exif *dbmodels.Exif) (*dbmodels.Exif, error)
	UpdatePhotoMetaItem(ctx context.Context, exif *dbmodels.Exif) (*dbmodels.Exif, error)
}

func NewExifRepository(db mysql.SQLExecutor) ExifRepository {
	return &exifRepository{db: db}
}

type exifRepository struct {
	db mysql.SQLExecutor
}

func (r *exifRepository) GetPhotoMetaItemByTagID(ctx context.Context, photoID, tagID int64) (*dbmodels.Exif, error) {
	m, err := dbmodels.Exifs(qm.Where("photo_id = ?", photoID), qm.Where("tag_id = ?", tagID)).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(errors.DBColumnNotFoundError, err)
		}
		return nil, err
	}
	return m, err
}

func (r *exifRepository) InsertPhotoMetaItem(ctx context.Context, exif *dbmodels.Exif) (*dbmodels.Exif, error) {
	if err := exif.Insert(ctx, r.db, boil.Infer()); err != nil {
		return nil, err
	}
	return exif, nil
}

func (r *exifRepository) UpdatePhotoMetaItem(ctx context.Context, exif *dbmodels.Exif) (*dbmodels.Exif, error) {
	if _, err := exif.Update(ctx, r.db, boil.Infer()); err != nil {
		return nil, err
	}
	return exif, nil
}
