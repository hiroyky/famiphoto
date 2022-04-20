package repositories

import (
	"context"
	"database/sql"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/dbmodels"
	"github.com/hiroyky/famiphoto/usecases"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func NewPhotoRepository(db SQLExecutor) usecases.PhotoAdapter {
	return &photoRepository{db: db}
}

type photoRepository struct {
	db SQLExecutor
}

func (r *photoRepository) InsertPhoto(ctx context.Context, photo *entities.Photo) (*entities.Photo, error) {
	p := &dbmodels.Photo{
		Name:       photo.Name,
		FilePath:   photo.FilePath,
		ImportedAt: photo.ImportedAt,
		GroupID:    photo.GroupID,
		OwnerID:    photo.OwnerID,
	}

	if err := p.Insert(ctx, r.db, boil.Infer()); err != nil {
		return nil, err
	}

	return r.toPhotoEntity(p), nil
}

func (r *photoRepository) UpdatePhoto(ctx context.Context, photo *entities.Photo) (*entities.Photo, error) {
	p := &dbmodels.Photo{
		PhotoID:    int(photo.PhotoID),
		Name:       photo.Name,
		FilePath:   photo.FilePath,
		ImportedAt: photo.ImportedAt,
		GroupID:    photo.GroupID,
		OwnerID:    photo.OwnerID,
	}
	if _, err := p.Update(ctx, r.db, boil.Infer()); err != nil {
		return nil, err
	}
	return r.toPhotoEntity(p), nil
}

func (r *photoRepository) GetPhotoByFilePath(ctx context.Context, filePath string) (*entities.Photo, error) {
	p, err := dbmodels.Photos(qm.Where("file_path = ?", filePath)).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(errors.DBColumnNotFoundError, err)
		}
		return nil, err
	}
	return r.toPhotoEntity(p), nil
}

func (r *photoRepository) toPhotoEntity(p *dbmodels.Photo) *entities.Photo {
	return &entities.Photo{
		PhotoID:    int64(p.PhotoID),
		Name:       p.Name,
		FilePath:   p.FilePath,
		ImportedAt: p.ImportedAt,
		GroupID:    p.GroupID,
		OwnerID:    p.OwnerID,
	}
}

func (r *photoRepository) InsertPhotoMetaItem(ctx context.Context, photoID int64, meta *entities.PhotoMetaItem) (entities.PhotoMeta, error) {
	panic("")
}

func (r *photoRepository) UpdatePhotoMetaItem(ctx context.Context, photoID int64, meta *entities.PhotoMetaItem) (entities.PhotoMeta, error) {
	panic("")
}

func (r *photoRepository) GetPhotoMetaItemByTagID(ctx context.Context, photoID, tagID int64) (*entities.PhotoMetaItem, error) {
	m, err := dbmodels.Exifs(qm.Where("photo_id = ?", photoID), qm.Where("tag_id = ?", tagID)).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(errors.DBColumnNotFoundError, err)
		}
		return nil, err
	}
	return &entities.PhotoMetaItem{
		TagID:       int64(m.TagID),
		TagName:     m.TagName,
		TagType:     m.TagType,
		Value:       m.Value,
		ValueString: m.ValueString,
	}, nil
}
