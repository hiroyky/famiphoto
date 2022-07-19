package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/filters"
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
	searchClient *elasticsearch.Client
	bulkIndexer  esutil.BulkIndexer
}

func (r *elasticSearchRepository) SearchPhotos(ctx context.Context, query filters.PhotoSearchRequest) (*models.PhotoResult, error) {
	res, err := r.searchPhotos(ctx, query)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New(errors.ElasticSearchFatal, err)
	}
	defer res.Body.Close()

	var body map[string]any
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return nil, errors.New(errors.ElasticSearchFatal, err)
	}

	return r.parsePhotoResult(body)
}

func (r *elasticSearchRepository) searchPhotos(ctx context.Context, query filters.PhotoSearchRequest) (*esapi.Response, error) {
	es := r.searchClient
	return es.Search(
		es.Search.WithContext(ctx),
		es.Search.WithIndex("photo"),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithBody(query.Body().MustBuffer()),
	)
}

func (r *elasticSearchRepository) parsePhotoResult(body map[string]any) (*models.PhotoResult, error) {
	photos := make([]*models.PhotoIndex, 0)
	for _, v := range body["hits"].(map[string]any)["hits"].([]any) {
		item, err := models.NewPhotoIndexFromMap(v.(map[string]any)["_source"].(map[string]any))
		if err != nil {
			return nil, errors.New(errors.ElasticSearchFatal, err)
		}
		photos = append(photos, item)
	}

	total := body["hits"].(map[string]any)["total"].(map[string]any)["value"].(float64)

	return &models.PhotoResult{
		Total:  int64(total),
		Photos: photos,
	}, nil
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
