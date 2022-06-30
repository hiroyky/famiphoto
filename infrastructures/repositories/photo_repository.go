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

type PhotoRepository interface {
	GetPhotos(ctx context.Context, limit, offset int64) ([]*dbmodels.Photo, error)
	CountPhotos(ctx context.Context) (int64, error)
	GetPhotoByFilePath(ctx context.Context, fileHash string) (*dbmodels.Photo, error)
	InsertPhoto(ctx context.Context, photo *dbmodels.Photo) (*dbmodels.Photo, error)
	UpdatePhoto(ctx context.Context, photo *dbmodels.Photo) (*dbmodels.Photo, error)
}

func NewPhotoRepository(db mysql.SQLExecutor) PhotoRepository {
	return &photoRepository{db: db}
}

type photoRepository struct {
	db mysql.SQLExecutor
}

func (r *photoRepository) GetPhotos(ctx context.Context, limit, offset int64) ([]*dbmodels.Photo, error) {
	return dbmodels.Photos(qm.Limit(int(limit)), qm.Offset(int(offset))).All(ctx, r.db)
}

func (r *photoRepository) CountPhotos(ctx context.Context) (int64, error) {
	return dbmodels.Photos().Count(ctx, r.db)
}

func (r *photoRepository) GetPhotoByFilePath(ctx context.Context, fileHash string) (*dbmodels.Photo, error) {
	p, err := dbmodels.Photos(qm.Where("file_name_hash = ?", fileHash)).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(errors.DBColumnNotFoundError, err)
		}
		return nil, err
	}
	return p, nil
}

func (r *photoRepository) InsertPhoto(ctx context.Context, photo *dbmodels.Photo) (*dbmodels.Photo, error) {
	if err := photo.Insert(ctx, r.db, boil.Infer()); err != nil {
		return nil, err
	}
	return photo, nil
}

func (r *photoRepository) UpdatePhoto(ctx context.Context, photo *dbmodels.Photo) (*dbmodels.Photo, error) {
	if _, err := photo.Update(ctx, r.db, boil.Blacklist(
		dbmodels.PhotoColumns.ImportedAt,
		dbmodels.PhotoColumns.GroupID,
		dbmodels.PhotoColumns.OwnerID,
	)); err != nil {
		return nil, err
	}
	return photo, nil
}
