package di

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/hiroyky/famiphoto/drivers/es"
	"github.com/hiroyky/famiphoto/drivers/mysql"
	"github.com/hiroyky/famiphoto/drivers/redis"
	"github.com/hiroyky/famiphoto/drivers/storage"
)

func NewElasticsearchClient() *elasticsearch.Client {
	return es.NewSearchClient()
}

func NewMySQLDriver() mysql.SQLExecutor {
	return mysql.NewDatabaseDriver()
}

func NewOAuthRedisDB() redis.Driver {
	return redis.NewOauthRedis()
}

func NewMediaSambaStorageDriver() storage.Driver {
	return storage.NewMediaSambaStorage()
}

func NewPhotoThumbnailStorageDriver() storage.Driver {
	return storage.NewPhotoThumbnailDriver()
}
