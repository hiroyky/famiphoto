package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/hiroyky/famiphoto/config"
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
