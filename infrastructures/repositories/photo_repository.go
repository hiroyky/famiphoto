package repositories

import (
	"context"
	"database/sql"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/dbmodels"
	"github.com/hiroyky/famiphoto/usecases"
	"github.com/hiroyky/famiphoto/utils/array"
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
		Name:         photo.Name,
		ImportedAt:   photo.ImportedAt,
		GroupID:      photo.GroupID,
		OwnerID:      photo.OwnerID,
		FileNameHash: photo.FileNameHash(),
	}

	if err := p.Insert(ctx, r.db, boil.Infer()); err != nil {
		return nil, err
	}

	return r.toPhotoEntity(p), nil
}

func (r *photoRepository) UpdatePhoto(ctx context.Context, photo *entities.Photo) (*entities.Photo, error) {
	p := &dbmodels.Photo{
		PhotoID:      int(photo.PhotoID),
		Name:         photo.Name,
		FileNameHash: photo.FileNameHash(),
	}
	if _, err := p.Update(ctx, r.db, boil.Blacklist(
		dbmodels.PhotoColumns.ImportedAt,
		dbmodels.PhotoColumns.GroupID,
		dbmodels.PhotoColumns.OwnerID,
	)); err != nil {
		return nil, err
	}
	return r.toPhotoEntity(p), nil
}

func (r *photoRepository) CountPhotos(ctx context.Context) (int64, error) {
	return dbmodels.Photos().Count(ctx, r.db)
}

func (r *photoRepository) GetPhotos(ctx context.Context, limit, offset int64) (entities.PhotoList, error) {
	photos, err := dbmodels.Photos(qm.Limit(int(limit)), qm.Offset(int(offset))).All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return array.Map(photos, r.toPhotoEntity), nil
}

func (r *photoRepository) GetPhotoByFilePath(ctx context.Context, filePath string) (*entities.Photo, error) {
	hash := entities.Photo{FilePath: filePath}.FileNameHash()
	p, err := dbmodels.Photos(qm.Where("file_name_hash = ?", hash)).One(ctx, r.db)
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
		ImportedAt: p.ImportedAt,
		GroupID:    p.GroupID,
		OwnerID:    p.OwnerID,
	}
}

func (r *photoRepository) GetPhotoFileByFilePath(ctx context.Context, filePath string) (*entities.PhotoFile, error) {
	p, err := dbmodels.PhotoFiles(qm.Where("file_path = ?", filePath)).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(errors.DBColumnNotFoundError, err)
		}
		return nil, err
	}
	return r.toPhotoFileEntity(p), err
}

func (r *photoRepository) toPhotoFileEntity(p *dbmodels.PhotoFile) *entities.PhotoFile {
	return &entities.PhotoFile{
		PhotoFileID: int64(p.PhotoFileID),
		PhotoID:     int64(p.PhotoID),
		FilePath:    p.FilePath,
		ImportedAt:  p.ImportedAt,
		GroupID:     p.GroupID,
		OwnerID:     p.OwnerID,
		FileHash:    p.FileHash,
	}
}

func (r *photoRepository) InsertPhotoFile(ctx context.Context, file *entities.PhotoFile) (*entities.PhotoFile, error) {
	m := &dbmodels.PhotoFile{
		PhotoID:    int(file.PhotoID),
		FileType:   file.FileType().ToString(),
		FilePath:   file.FilePath,
		ImportedAt: file.ImportedAt,
		GroupID:    file.GroupID,
		OwnerID:    file.OwnerID,
		FileHash:   file.FileHash,
	}
	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		return nil, err
	}
	return r.toPhotoFileEntity(m), nil
}

func (r *photoRepository) UpdatePhotoFile(ctx context.Context, file *entities.PhotoFile) (*entities.PhotoFile, error) {
	m := &dbmodels.PhotoFile{
		PhotoFileID: int(file.PhotoFileID),
		PhotoID:     int(file.PhotoID),
		FileType:    file.FileType().ToString(),
		FilePath:    file.FilePath,
		FileHash:    file.FileHash,
	}
	if _, err := m.Update(ctx, r.db, boil.Blacklist(
		dbmodels.PhotoColumns.ImportedAt,
		dbmodels.PhotoColumns.GroupID,
		dbmodels.PhotoColumns.OwnerID,
	)); err != nil {
		return nil, err
	}
	return r.toPhotoFileEntity(m), nil
}

func (r *photoRepository) GetPhotoFilesByPhotoIDs(ctx context.Context, photoIDs []int64) ([]*entities.PhotoFile, error) {
	files, err := dbmodels.PhotoFiles(qm.WhereIn("photo_id in ?", toInterfaceSlice(photoIDs)...)).All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return array.Map(files, r.toPhotoFileEntity), nil
}

func (r *photoRepository) InsertPhotoMetaItem(ctx context.Context, photoID int64, item *entities.PhotoMetaItem) (*entities.PhotoMetaItem, error) {
	m := &dbmodels.Exif{
		PhotoID:     int(photoID),
		TagID:       int(item.TagID),
		TagName:     item.TagName,
		TagType:     item.TagType,
		ValueString: item.ValueString,
		SortOrder:   0,
	}
	if err := m.Insert(ctx, r.db, boil.Infer()); err != nil {
		return nil, err
	}
	return r.toPhotoMetaItem(m), nil
}

func (r *photoRepository) UpdatePhotoMetaItem(ctx context.Context, photoID int64, item *entities.PhotoMetaItem) (*entities.PhotoMetaItem, error) {
	m := &dbmodels.Exif{
		ExifID:      int(item.PhotoMetaItemID),
		PhotoID:     int(photoID),
		TagID:       int(item.TagID),
		TagName:     item.TagName,
		TagType:     item.TagType,
		ValueString: item.ValueString,
	}
	if _, err := m.Update(ctx, r.db, boil.Infer()); err != nil {
		return nil, err
	}
	return r.toPhotoMetaItem(m), nil
}

func (r *photoRepository) GetPhotoMetaItemByTagID(ctx context.Context, photoID, tagID int64) (*entities.PhotoMetaItem, error) {
	m, err := dbmodels.Exifs(qm.Where("photo_id = ?", photoID), qm.Where("tag_id = ?", tagID)).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(errors.DBColumnNotFoundError, err)
		}
		return nil, err
	}
	return r.toPhotoMetaItem(m), nil
}

func (r *photoRepository) toPhotoMetaItem(m *dbmodels.Exif) *entities.PhotoMetaItem {
	return &entities.PhotoMetaItem{
		PhotoMetaItemID: int64(m.ExifID),
		TagID:           int64(m.TagID),
		TagName:         m.TagName,
		TagType:         m.TagType,
		ValueString:     m.ValueString,
	}
}
