package elasticsearch

import (
	"github.com/cenkalti/backoff/v4"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/hiroyky/famiphoto/config"
	"runtime"
	"time"
)

type Client interface {
}

var searchClient Client = nil

func NewSearchClient() Client {
	if searchClient != nil {
		return searchClient
	}

	cfg := elasticsearch.Config{
		Addresses: config.Env.ElasticsearchAddresses,
		Username:  "",
		Password:  "",
	}

	c, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	searchClient = c

	return searchClient
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
