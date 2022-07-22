package infrastructures

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/dbmodels"
	"github.com/hiroyky/famiphoto/infrastructures/repositories"
	"github.com/hiroyky/famiphoto/utils/array"
)

type PhotoAdapter interface {
	GetPhotos(ctx context.Context, limit, offset int) (entities.PhotoList, error)
	CountPhotos(ctx context.Context) (int, error)
	UpsertPhotoByFilePath(ctx context.Context, photo *entities.Photo) (*entities.Photo, error)
	UpsertPhotoMetaItemByPhotoTagID(ctx context.Context, photoID int, metaItem *entities.PhotoMetaItem) (*entities.PhotoMetaItem, error)
}

func NewPhotoAdapter(
	photoRepo repositories.PhotoRepository,
	photoFileRepo repositories.PhotoFileRepository,
	exifRepo repositories.ExifRepository,
) PhotoAdapter {
	return &photoAdapter{
		photoRepo:     photoRepo,
		photoFileRepo: photoFileRepo,
		exifRepo:      exifRepo,
	}
}

type photoAdapter struct {
	photoRepo     repositories.PhotoRepository
	photoFileRepo repositories.PhotoFileRepository
	exifRepo      repositories.ExifRepository
}

func (a *photoAdapter) GetPhotos(ctx context.Context, limit, offset int) (entities.PhotoList, error) {
	photos, err := a.photoRepo.GetPhotos(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	photoIDs := array.Map(photos, func(p *dbmodels.Photo) int {
		return p.PhotoID
	})

	files, err := a.photoFileRepo.GetPhotoFilesByPhotoID(ctx, photoIDs)
	if err != nil {
		return nil, err
	}

	photoEntities := array.Map(photos, func(photo *dbmodels.Photo) *entities.Photo {
		return a.toPhotoEntity(photo, files)
	})

	return photoEntities, nil
}

func (a *photoAdapter) CountPhotos(ctx context.Context) (int, error) {
	return a.photoRepo.CountPhotos(ctx)
}

func (a *photoAdapter) UpsertPhotoByFilePath(ctx context.Context, photo *entities.Photo) (*entities.Photo, error) {
	dstDBPhoto, err := a.upsertPhotoByFilePath(ctx, &dbmodels.Photo{
		Name:         photo.Name,
		ImportedAt:   photo.ImportedAt,
		GroupID:      photo.GroupID,
		OwnerID:      photo.OwnerID,
		FileNameHash: photo.FileNameHash,
	})
	if err != nil {
		return nil, err
	}

	dstDBFiles := make([]*dbmodels.PhotoFile, 0)
	for _, file := range photo.Files {
		dst, err := a.upsertPhotoFileByFilePath(ctx, &dbmodels.PhotoFile{
			PhotoID:    dstDBPhoto.PhotoID,
			FileType:   file.FileType().ToString(),
			FilePath:   file.FilePath,
			ImportedAt: file.ImportedAt,
			GroupID:    file.GroupID,
			OwnerID:    file.OwnerID,
			FileHash:   file.FileHash,
		})
		if err != nil {
			return nil, err
		}
		dstDBFiles = append(dstDBFiles, dst)
	}

	return a.toPhotoEntity(dstDBPhoto, dstDBFiles), nil
}

func (a *photoAdapter) upsertPhotoByFilePath(ctx context.Context, photo *dbmodels.Photo) (*dbmodels.Photo, error) {
	existPhoto, err := a.photoRepo.GetPhotoByFilePath(ctx, photo.FileNameHash)
	if err != nil && !errors.IsErrCode(err, errors.DBColumnNotFoundError) {
		return nil, err
	}
	if err == nil && existPhoto != nil {
		photo.PhotoID = existPhoto.PhotoID
		return a.photoRepo.UpdatePhoto(ctx, photo)
	}
	return a.photoRepo.InsertPhoto(ctx, photo)
}

func (a *photoAdapter) upsertPhotoFileByFilePath(ctx context.Context, photoFile *dbmodels.PhotoFile) (*dbmodels.PhotoFile, error) {
	exist, err := a.photoFileRepo.GetPhotoFileByFilePath(ctx, photoFile.FilePath)
	if err != nil && !errors.IsErrCode(err, errors.DBColumnNotFoundError) {
		return nil, err
	}
	if err == nil && exist != nil {
		photoFile.PhotoFileID = exist.PhotoFileID
		return a.photoFileRepo.UpdatePhotoFile(ctx, photoFile)
	}
	return a.photoFileRepo.InsertPhotoFile(ctx, photoFile)
}

func (a *photoAdapter) toPhotoEntity(photo *dbmodels.Photo, files []*dbmodels.PhotoFile) *entities.Photo {
	photoFiles := array.Filter(files, func(t *dbmodels.PhotoFile) bool {
		return t.PhotoID == photo.PhotoID
	})
	photoFileEntities := array.Map(photoFiles, func(t *dbmodels.PhotoFile) *entities.PhotoFile {
		return &entities.PhotoFile{
			PhotoFileID: t.PhotoFileID,
			PhotoID:     t.PhotoID,
			FilePath:    t.FilePath,
			ImportedAt:  t.ImportedAt,
			GroupID:     t.GroupID,
			OwnerID:     t.OwnerID,
			FileHash:    t.FileHash,
		}
	})

	return &entities.Photo{
		PhotoID:      photo.PhotoID,
		Name:         photo.Name,
		ImportedAt:   photo.ImportedAt,
		GroupID:      photo.GroupID,
		OwnerID:      photo.OwnerID,
		FileNameHash: photo.FileNameHash,
		Files:        photoFileEntities,
	}
}

func (a *photoAdapter) UpsertPhotoMetaItemByPhotoTagID(ctx context.Context, photoID int, metaItem *entities.PhotoMetaItem) (*entities.PhotoMetaItem, error) {
	dbMetaItem := &dbmodels.Exif{
		PhotoID:     photoID,
		TagID:       metaItem.TagID,
		TagName:     metaItem.TagName,
		TagType:     metaItem.TagType,
		ValueString: metaItem.ValueString,
		SortOrder:   0,
	}

	exitTag, err := a.exifRepo.GetPhotoMetaItemByTagID(ctx, photoID, metaItem.TagID)
	if err == nil && exitTag != nil {
		dbMetaItem.ExifID = exitTag.ExifID
		dst, err := a.exifRepo.UpdatePhotoMetaItem(ctx, dbMetaItem)
		if err != nil {
			return nil, err
		}
		return a.toPhotoMetaItemEntity(dst), nil
	}

	dst, err := a.exifRepo.InsertPhotoMetaItem(ctx, dbMetaItem)
	if err != nil {
		return nil, err
	}
	return a.toPhotoMetaItemEntity(dst), nil
}

func (a *photoAdapter) toPhotoMetaItemEntity(m *dbmodels.Exif) *entities.PhotoMetaItem {
	return &entities.PhotoMetaItem{
		PhotoMetaItemID: m.ExifID,
		TagID:           m.TagID,
		TagName:         m.TagName,
		TagType:         m.TagType,
		ValueString:     m.ValueString,
	}
}
