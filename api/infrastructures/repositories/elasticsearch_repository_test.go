package repositories

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/hiroyky/famiphoto/drivers/es"
	"github.com/hiroyky/famiphoto/infrastructures/filters"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type mockTransport struct {
	RoundTripFunc func(req *http.Request) (*http.Response, error)
}

func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.RoundTripFunc(req)
}

func TestElasticSearchRepository_searchPhotos_リクエスト(t *testing.T) {
	tp := &mockTransport{
		RoundTripFunc: func(req *http.Request) (*http.Response, error) {
			fmt.Printf("%#v\n", req.URL.Query())
			return &http.Response{
				Status:     "200 OK",
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(strings.NewReader("HELLO")),
			}, nil
		},
	}
	esClient := es.NewSearchWithConfig(elasticsearch.Config{
		Addresses: []string{"http://test_es"},
		Transport: tp,
	})

	esRepo := &elasticSearchRepository{
		searchClient: esClient,
	}

	_, _ = esRepo.searchPhotos(context.Background(), &filters.PhotoSearchQuery{})
}
