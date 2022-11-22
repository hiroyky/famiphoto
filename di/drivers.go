package di

import (
	"github.com/elastic/go-elasticsearch/v8"
	native "github.com/go-redis/redis/v8"
	"github.com/hiroyky/famiphoto/config"
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

var redisDriver redis.Driver = nil

func NewOAuthRedisDB() redis.Driver {
	if redisDriver != nil {
		return redisDriver
	}
	redisDriver = newRedis()
	return redisDriver
}

func newRedis() redis.Driver {
	n := native.NewClient(&native.Options{
		Addr:     config.Env.OauthRedisHostName,
		Password: "",
		DB:       int(config.Env.OauthRedisDatabase),
	})
	db := redis.NewRedisDriver(n)
	return db
}

func NewMediaSambaStorageDriver() storage.Driver {
	return storage.NewMediaSambaStorage()
}

func NewPhotoThumbnailStorageDriver() storage.Driver {
	return storage.NewPhotoThumbnailDriver()
}

func NewMediaLocalStorageDriver() storage.Driver {
	return storage.NewLocalStorageDriver("/mnt/famiphoto")
}
