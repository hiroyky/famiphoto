package di

import (
	"github.com/hiroyky/famiphoto/drivers/elasticsearch"
	"github.com/hiroyky/famiphoto/drivers/mysql"
	"github.com/hiroyky/famiphoto/drivers/redis"
	"github.com/hiroyky/famiphoto/drivers/storage"
)

func NewElasticsearchClient() elasticsearch.Client {
	return elasticsearch.NewSearchClient()
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
