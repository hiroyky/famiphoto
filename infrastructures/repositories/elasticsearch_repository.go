package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/infrastructures/models"
	"github.com/hiroyky/famiphoto/usecases"
	"github.com/hiroyky/famiphoto/utils/array"
	"time"
)

func NewElasticSearchRepository(bulkIndexer esutil.BulkIndexer) usecases.SearchAdapter {
	return &elasticSearchRepository{
		bulkIndexer: bulkIndexer,
	}
}

type elasticSearchRepository struct {
	bulkIndexer esutil.BulkIndexer
}

func (r *elasticSearchRepository) BulkInsertPhoto(ctx context.Context, photos []*entities.Photo, photoFiles entities.PhotoFileList, dateTimeOriginal *entities.PhotoMetaItem) (*esutil.BulkIndexerStats, error) {
	for _, photo := range photos {
		//fmt.Println(photo)
		p := &models.PhotoIndex{
			PhotoID: photo.PhotoID,
			OwnerID: photo.OwnerID,
			GroupID: photo.GroupID,
			FileTypes: array.Map(photoFiles.FindFileTypesByPhotoID(photo.PhotoID), func(t entities.PhotoFileType) string {
				return t.ToString()
			}),
			Name:             photo.Name,
			ImportedAt:       photo.ImportedAt.Unix(),
			DateTimeOriginal: time.Now().Unix(),
			PreviewURL:       photo.PreviewURL(),
			ThumbnailURL:     photo.ThumbnailURL(),
		}

		data, err := json.Marshal(p)
		if err != nil {
			return nil, err
		}

		if err := r.bulkIndexer.Add(ctx, esutil.BulkIndexerItem{
			Action:     "index",
			DocumentID: p.PhotoIndexID(),
			Body:       bytes.NewReader(data),
			OnSuccess:  r.onBulkInsertSuccess,
			OnFailure:  r.onBulkInsertFail,
		}); err != nil {
			return nil, err
		}
	}
	if err := r.bulkIndexer.Close(ctx); err != nil {
		return nil, err
	}

	stats := r.bulkIndexer.Stats()

	return &stats, nil
}

func (r *elasticSearchRepository) onBulkInsertSuccess(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
}

func (r *elasticSearchRepository) onBulkInsertFail(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
	if err != nil {
		fmt.Println("error item ", err)
		return
	}
	fmt.Println("error item", item, res.Error.Type, res.Error.Reason)
}
