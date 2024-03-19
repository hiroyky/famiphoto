package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/drivers/es"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/filters"
	"github.com/hiroyky/famiphoto/infrastructures/models"
)

func NewElasticSearchRepository(searchClient es.Search, newBulkIndexerFunc func() esutil.BulkIndexer) ElasticSearchRepository {
	return &elasticSearchRepository{
		searchClient:       searchClient,
		newBulkIndexerFunc: newBulkIndexerFunc,
	}
}

type ElasticSearchRepository interface {
	InsertPhoto(ctx context.Context, photo *models.PhotoIndex) error
	BulkInsertPhotos(ctx context.Context, photos []*models.PhotoIndex) (*esutil.BulkIndexerStats, error)
	SearchPhotos(ctx context.Context, query *filters.PhotoSearchQuery) (*models.PhotoResult, error)
	AggregateByDateTimeOriginalYear(ctx context.Context) ([]*models.PhotoDateHistogram, error)
	AggregateByDateTimeOriginalYearMonth(ctx context.Context, year int) ([]*models.PhotoDateHistogram, error)
	AggregateByDateTimeOriginalYearMonthDate(ctx context.Context, year, month int) ([]*models.PhotoDateHistogram, error)
}

type elasticSearchRepository struct {
	searchClient       es.Search
	newBulkIndexerFunc func() esutil.BulkIndexer
}

func (r *elasticSearchRepository) SearchPhotos(ctx context.Context, query *filters.PhotoSearchQuery) (*models.PhotoResult, error) {
	res, err := r.searchPhotos(ctx, query)
	if err != nil {
		return nil, err
	}

	return r.parsePhotoResult(res)
}

func (r *elasticSearchRepository) searchPhotos(_ context.Context, query *filters.PhotoSearchQuery) (*es.SearchResponse, error) {
	res, err := r.searchClient.Search("photo", query.Body())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *elasticSearchRepository) parsePhotoResult(res *es.SearchResponse) (*models.PhotoResult, error) {
	photos := make([]*models.PhotoIndex, 0)
	for _, v := range res.Hits.Hits {
		item, err := models.NewPhotoIndexFromMap(v.Source)
		if err != nil {
			return nil, errors.New(errors.ElasticSearchFatal, err)
		}
		photos = append(photos, item)
	}

	total := res.Hits.Total.Value

	return &models.PhotoResult{
		Total:  int(total),
		Photos: photos,
	}, nil
}

func (r *elasticSearchRepository) InsertPhoto(ctx context.Context, photo *models.PhotoIndex) error {
	_, err := r.BulkInsertPhotos(ctx, []*models.PhotoIndex{photo})
	return err
}

func (r *elasticSearchRepository) BulkInsertPhotos(ctx context.Context, photos []*models.PhotoIndex) (*esutil.BulkIndexerStats, error) {
	bulkIndexer := r.newBulkIndexerFunc()
	for _, p := range photos {
		data, err := json.Marshal(p)
		if err != nil {
			return nil, err
		}

		if err := bulkIndexer.Add(ctx, esutil.BulkIndexerItem{
			Action:     "index",
			DocumentID: p.PhotoIndexID(),
			Body:       bytes.NewReader(data),
			OnSuccess:  r.onBulkInsertSuccess,
			OnFailure:  r.onBulkInsertFail,
		}); err != nil {
			return nil, err
		}
	}
	if err := bulkIndexer.Close(ctx); err != nil {
		return nil, err
	}

	stats := bulkIndexer.Stats()

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

func (r *elasticSearchRepository) AggregateByDateTimeOriginalYear(ctx context.Context) ([]*models.PhotoDateHistogram, error) {
	key := "date_time_original"
	query := filters.NewAggregateByDateTimeOriginalYear(key)

	res, err := r.searchClient.Search("photo", query)
	if err != nil {
		return nil, err
	}

	result := r.parseAggregationByDateResult(res)
	return result, nil
}

func (r *elasticSearchRepository) AggregateByDateTimeOriginalYearMonth(ctx context.Context, year int) ([]*models.PhotoDateHistogram, error) {
	key := "date_time_original"
	query := filters.NewAggregateByDateTimeOriginalYearMonth(key, year, config.Env.ExifTimezone)

	res, err := r.searchClient.Search("photo", query)
	if err != nil {
		return nil, err
	}

	result := r.parseAggregationByDateResult(res)
	return result, nil
}

func (r *elasticSearchRepository) AggregateByDateTimeOriginalYearMonthDate(ctx context.Context, year, month int) ([]*models.PhotoDateHistogram, error) {
	key := "date_time_original"
	query := filters.NewAggregateByDateTimeOriginalYearMonthDate(key, year, month, config.Env.ExifTimezone)

	res, err := r.searchClient.Search("photo", query)
	if err != nil {
		return nil, err
	}

	result := r.parseAggregationByDateResult(res)
	return result, nil
}

func (r *elasticSearchRepository) parseAggregationByDateResult(res *es.SearchResponse) []*models.PhotoDateHistogram {
	key := "date_time_original"
	result := make([]*models.PhotoDateHistogram, len(res.Aggregations[key].Buckets))
	for i, b := range res.Aggregations[key].Buckets {
		result[i] = &models.PhotoDateHistogram{
			EpochSec: b.Key / 1000,
			DocCount: int(b.DocCount),
		}
	}
	return result
}
