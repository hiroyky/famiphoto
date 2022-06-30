package infrastructures

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/infrastructures/models"
	"github.com/hiroyky/famiphoto/infrastructures/repositories"
	"github.com/hiroyky/famiphoto/utils/array"
	"time"
)

type SearchAdapter interface {
	BulkInsertPhotos(ctx context.Context, photos entities.PhotoList, dateTimeOriginal *entities.PhotoMetaItem) (*esutil.BulkIndexerStats, error)
}

func NewSearchAdapter(esRepo repositories.ElasticSearchRepository) SearchAdapter {
	return &searchAdapter{
		esRepo: esRepo,
	}
}

type searchAdapter struct {
	esRepo repositories.ElasticSearchRepository
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
			DateTimeOriginal: time.Now().Unix(),
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
