package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/infrastructures/models"
)

func NewElasticSearchRepository(bulkIndexer esutil.BulkIndexer) ElasticSearchRepository {
	return &elasticSearchRepository{
		bulkIndexer: bulkIndexer,
	}
}

type ElasticSearchRepository interface {
	BulkInsertPhotos(ctx context.Context, photos []*models.PhotoIndex, dateTimeOriginal *entities.PhotoMetaItem) (*esutil.BulkIndexerStats, error)
}

type elasticSearchRepository struct {
	bulkIndexer esutil.BulkIndexer
}

func (r *elasticSearchRepository) BulkInsertPhotos(ctx context.Context, photos []*models.PhotoIndex, dateTimeOriginal *entities.PhotoMetaItem) (*esutil.BulkIndexerStats, error) {
	for _, p := range photos {
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
