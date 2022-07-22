package infrastructures

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/infrastructures/filters"
	"github.com/hiroyky/famiphoto/infrastructures/models"
	"github.com/hiroyky/famiphoto/infrastructures/repositories"
	"github.com/hiroyky/famiphoto/utils/array"
	"time"
)

type SearchAdapter interface {
	BulkInsertPhotos(ctx context.Context, photos entities.PhotoList, dateTimeOriginal *entities.PhotoMetaItem) (*esutil.BulkIndexerStats, error)
	SearchPhotos(ctx context.Context, q *filters.PhotoSearchQuery) (*entities.PhotoSearchResult, error)
}

func NewSearchAdapter(esRepo repositories.ElasticSearchRepository) SearchAdapter {
	return &searchAdapter{
		esRepo:  esRepo,
		nowFunc: time.Now,
	}
}

type searchAdapter struct {
	esRepo  repositories.ElasticSearchRepository
	nowFunc func() time.Time
}

func (a *searchAdapter) BulkInsertPhotos(ctx context.Context, photos entities.PhotoList, dateTimeOriginal *entities.PhotoMetaItem) (*esutil.BulkIndexerStats, error) {
	photoIndexes := array.Map(photos, func(photo *entities.Photo) *models.PhotoIndex {
		return &models.PhotoIndex{
			PhotoID: photo.PhotoID,
			OwnerID: photo.OwnerID,
			GroupID: photo.GroupID,
			FileTypes: array.Map(photo.Files.FindFileTypesByPhotoID(photo.PhotoID), func(t entities.PhotoFileType) string {
				return t.ToString()
			}),
			Name:             photo.Name,
			ImportedAt:       photo.ImportedAt.Unix(),
			DateTimeOriginal: a.nowFunc().Unix(),
			PreviewURL:       photo.PreviewURL(),
			ThumbnailURL:     photo.ThumbnailURL(),
		}
	})

	stats, err := a.esRepo.BulkInsertPhotos(ctx, photoIndexes, dateTimeOriginal)
	if err != nil {
		return nil, err
	}
	return stats, nil
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
