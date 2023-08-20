package es

import (
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/errors"
	"io"
	"net/http"
	"runtime"
	"time"
)

type Search interface {
	CreateIndex(index string, mapping Mapping) error
	Index(index string, body IndexBody) error
	Search(index string, body *SearchRequestBody) (*SearchResponse, error)
}

func NewSearch(addresses []string, userName, password, fingerPrint string) Search {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses:               addresses,
		Username:                userName,
		Password:                password,
		CloudID:                 "",
		APIKey:                  "",
		ServiceToken:            "",
		CertificateFingerprint:  fingerPrint,
		Header:                  nil,
		CACert:                  nil,
		RetryOnStatus:           nil,
		DisableRetry:            false,
		MaxRetries:              0,
		RetryOnError:            nil,
		CompressRequestBody:     false,
		DiscoverNodesOnStart:    false,
		DiscoverNodesInterval:   0,
		EnableMetrics:           false,
		EnableDebugLogger:       false,
		EnableCompatibilityMode: false,
		DisableMetaHeader:       false,
		RetryBackoff:            nil,
		Transport:               nil,
		Logger:                  nil,
		Selector:                nil,
		ConnectionPoolFunc:      nil,
	})
	if err != nil {
		panic(err)
	}

	return &driver{
		client: client,
	}
}

type driver struct {
	client *elasticsearch.Client
}

func (d *driver) handleError(res *esapi.Response, err error) error {
	if err != nil {
		return err
	}

	if res.StatusCode >= http.StatusBadRequest {
		body, _ := io.ReadAll(res.Body)
		return errors.New(errors.ElasticSearchFatal, fmt.Errorf(string(body)))
	}

	return nil
}

func NewBulkClient() esutil.BulkIndexer {
	retryBackoff := backoff.NewExponentialBackOff()
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses:     config.Env.ElasticsearchAddresses,
		Username:      "",
		Password:      "",
		RetryOnStatus: []int{502, 503, 504, 429},
		RetryBackoff: func(i int) time.Duration {
			if i == 1 {
				retryBackoff.Reset()
			}
			return retryBackoff.NextBackOff()
		},
		MaxRetries: 5,
	})
	if err != nil {
		panic(err)
	}

	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		NumWorkers:    runtime.NumCPU(),
		FlushBytes:    5e+6,
		FlushInterval: 30 * time.Second,
		Client:        es,
		Index:         config.ElasticsearchPhotoIndexName,
	})
	if err != nil {
		panic(err)
	}

	return bi
}
