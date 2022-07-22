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

type PhotoFileRepository interface {
	GetPhotoFilesByPhotoID(ctx context.Context, photoIDs []int) ([]*dbmodels.PhotoFile, error)
	InsertPhotoFile(ctx context.Context, photoFile *dbmodels.PhotoFile) (*dbmodels.PhotoFile, error)
	UpdatePhotoFile(ctx context.Context, photoFile *dbmodels.PhotoFile) (*dbmodels.PhotoFile, error)
	GetPhotoFileByFilePath(ctx context.Context, filePath string) (*dbmodels.PhotoFile, error)
}

func NewPhotoFileRepository(db mysql.SQLExecutor) PhotoFileRepository {
	return &photoFileRepository{db: db}
}

type photoFileRepository struct {
	db mysql.SQLExecutor
}

func (r *photoFileRepository) GetPhotoFilesByPhotoID(ctx context.Context, photoIDs []int) ([]*dbmodels.PhotoFile, error) {
	return dbmodels.PhotoFiles(qm.WhereIn("photo_id in ?", toInterfaceSlice(photoIDs)...)).All(ctx, r.db)
}

func (r *photoFileRepository) InsertPhotoFile(ctx context.Context, photoFile *dbmodels.PhotoFile) (*dbmodels.PhotoFile, error) {
	if err := photoFile.Insert(ctx, r.db, boil.Infer()); err != nil {
		return nil, err
	}
	return photoFile, nil
}

func (r *photoFileRepository) UpdatePhotoFile(ctx context.Context, photoFile *dbmodels.PhotoFile) (*dbmodels.PhotoFile, error) {
	if _, err := photoFile.Update(ctx, r.db, boil.Blacklist(
		dbmodels.PhotoColumns.ImportedAt,
		dbmodels.PhotoColumns.GroupID,
		dbmodels.PhotoColumns.OwnerID,
	)); err != nil {
		return nil, err
	}
	return photoFile, nil
}

func (r *photoFileRepository) GetPhotoFileByFilePath(ctx context.Context, filePath string) (*dbmodels.PhotoFile, error) {
	p, err := dbmodels.PhotoFiles(qm.Where("file_path = ?", filePath)).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(errors.DBColumnNotFoundError, err)
		}
		return nil, err
	}
	return p, nil
}
