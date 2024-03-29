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
	GetPhotoFileByPhotoFileID(ctx context.Context, photoFileID int) (*dbmodels.PhotoFile, error)
	GetPhotoFilesByPhotoID(ctx context.Context, photoID int) ([]*dbmodels.PhotoFile, error)
	GetPhotoFilesByPhotoIDs(ctx context.Context, photoIDs []int) ([]*dbmodels.PhotoFile, error)
	InsertPhotoFile(ctx context.Context, photoFile *dbmodels.PhotoFile) (*dbmodels.PhotoFile, error)
	UpdatePhotoFile(ctx context.Context, photoFile *dbmodels.PhotoFile) (*dbmodels.PhotoFile, error)
	GetPhotoFileByFilePath(ctx context.Context, filePath string) (*dbmodels.PhotoFile, error)
	ExistPhotoFileByFilePath(ctx context.Context, filePath string) (bool, error)
}

func NewPhotoFileRepository(db mysql.SQLExecutor) PhotoFileRepository {
	return &photoFileRepository{db: db}
}

type photoFileRepository struct {
	db mysql.SQLExecutor
}

func (r *photoFileRepository) GetPhotoFileByPhotoFileID(ctx context.Context, photoFileID int) (*dbmodels.PhotoFile, error) {
	file, err := dbmodels.FindPhotoFile(ctx, r.db, photoFileID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New(errors.DBRowNotFoundError, err)
	}
	return file, nil
}

func (r *photoFileRepository) GetPhotoFilesByPhotoID(ctx context.Context, photoID int) ([]*dbmodels.PhotoFile, error) {
	return dbmodels.PhotoFiles(qm.Where("photo_id = ?", photoID)).All(ctx, r.db)
}

func (r *photoFileRepository) GetPhotoFilesByPhotoIDs(ctx context.Context, photoIDs []int) ([]*dbmodels.PhotoFile, error) {
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
	)); err != nil {
		return nil, err
	}
	return photoFile, nil
}

func (r *photoFileRepository) GetPhotoFileByFilePath(ctx context.Context, filePath string) (*dbmodels.PhotoFile, error) {
	p, err := dbmodels.PhotoFiles(qm.Where("file_path = ?", filePath)).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(errors.DBRowNotFoundError, err)
		}
		return nil, err
	}
	return p, nil
}

func (r *photoFileRepository) ExistPhotoFileByFilePath(ctx context.Context, filePath string) (bool, error) {
	return dbmodels.PhotoFiles(qm.Where("file_path = ?", filePath)).Exists(ctx, r.db)
}
