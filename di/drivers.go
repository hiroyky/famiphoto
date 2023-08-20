package di

import (
	native "github.com/go-redis/redis/v8"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/drivers/es"
	"github.com/hiroyky/famiphoto/drivers/mysql"
	"github.com/hiroyky/famiphoto/drivers/redis"
	"github.com/hiroyky/famiphoto/drivers/storage"
)

var search es.Search = nil

func NewSearch() es.Search {
	if search != nil {
		return search
	}

	s := es.NewSearch(
		config.Env.ElasticsearchAddresses,
		config.Env.ElasticsearchUserName,
		config.Env.ElasticsearchPassword,
		config.Env.ElasticsearchFingerPrint,
	)
	search = s
	return search
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

func NewPhotoThumbnailStorageDriver() storage.Driver {
	return storage.NewPhotoThumbnailDriver()
}

func NewMediaLocalStorageDriver() storage.Driver {
	return storage.NewLocalStorageDriver("/mnt/famiphoto")
}

func NewTempStorageDriver() storage.Driver {
	return storage.NewLocalStorageDriver("/var/famiphoto")
}
