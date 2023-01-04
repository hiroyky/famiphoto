package infrastructures

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/infrastructures/dbmodels"
	"github.com/hiroyky/famiphoto/infrastructures/filters"
	"github.com/hiroyky/famiphoto/infrastructures/models"
	"github.com/hiroyky/famiphoto/infrastructures/repositories"
	"github.com/hiroyky/famiphoto/utils"
	"github.com/hiroyky/famiphoto/utils/array"
	"time"
)

type SearchAdapter interface {
	InsertPhoto(ctx context.Context, photo *entities.Photo) error
	BulkInsertPhotos(ctx context.Context, photos entities.PhotoList) (*esutil.BulkIndexerStats, error)
	SearchPhotos(ctx context.Context, q *filters.PhotoSearchQuery) (*entities.PhotoSearchResult, error)
}

func NewSearchAdapter(esRepo repositories.ElasticSearchRepository, exifRepo repositories.ExifRepository) SearchAdapter {
	return &searchAdapter{
		esRepo:   esRepo,
		exifRepo: exifRepo,
		nowFunc:  time.Now,
	}
}

type searchAdapter struct {
	esRepo   repositories.ElasticSearchRepository
	exifRepo repositories.ExifRepository
	nowFunc  func() time.Time
}

func (a *searchAdapter) InsertPhoto(ctx context.Context, photo *entities.Photo) error {
	dateTimeOriginalTag, err := a.exifRepo.GetPhotoMetaItemByTagID(ctx, photo.PhotoID, config.ExifTagIDDateTimeOriginal)
	if err != nil {
		return err
	}

	var dateTimeOriginalEpoc int64
	if dateTimeOriginal, err := a.parseExifDatetime(dateTimeOriginalTag.ValueString); err == nil {
		dateTimeOriginalEpoc = dateTimeOriginal.Unix()
	}

	photoIndex := models.NewPhotoIndex(photo, dateTimeOriginalEpoc)
	return a.esRepo.InsertPhoto(ctx, photoIndex)
}

func (a *searchAdapter) BulkInsertPhotos(ctx context.Context, photos entities.PhotoList) (*esutil.BulkIndexerStats, error) {
	dateTimeOriginals, err := a.exifRepo.GetPhotoMetaItemsByPhotoIDsTagID(ctx, photos.PhotoIDs(), config.ExifTagIDDateTimeOriginal)
	if err != nil {
		return nil, err
	}

	photoIndexes := array.Map(photos, func(photo *entities.Photo) *models.PhotoIndex {
		var dateTimeOriginalEpoc int64
		dateTimeOriginalTag := array.Filter(dateTimeOriginals, func(t *dbmodels.Exif) bool { return t.PhotoID == photo.PhotoID })
		if len(dateTimeOriginalTag) > 0 {
			if dateTimeOriginal, err := a.parseExifDatetime(dateTimeOriginalTag[0].ValueString); err == nil {
				dateTimeOriginalEpoc = dateTimeOriginal.Unix()
			}
		}

		return models.NewPhotoIndex(photo, dateTimeOriginalEpoc)
	})

	stats, err := a.esRepo.BulkInsertPhotos(ctx, photoIndexes)
	if err != nil {
		return nil, err
	}
	return stats, nil
}

func (a *searchAdapter) parseExifDatetime(valueString string) (time.Time, error) {
	return utils.ParseDatetime(valueString, utils.MustLoadLocation(config.Env.ExifTimezone))
}

func (a *searchAdapter) SearchPhotos(ctx context.Context, q *filters.PhotoSearchQuery) (*entities.PhotoSearchResult, error) {
	res, err := a.esRepo.SearchPhotos(ctx, q)
	if err != nil {
		return nil, err
	}

	dstItems := array.Map(res.Photos, func(photo *models.PhotoIndex) *entities.PhotoSearchResultItem {
		return photo.ToEntityItem()
	})

	return &entities.PhotoSearchResult{
		Items: dstItems,
		Total: res.Total,
	}, nil
}
