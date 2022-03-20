package redis

import (
	"context"
	native "github.com/go-redis/redis/v8"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/infrastructures/repositories"
	"time"
)

type redisDB struct {
	client *native.Client
}

func (r *redisDB) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *redisDB) SetEx(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	return r.client.SetEX(ctx, key, val, expiration).Err()
}

var oauthDB *redisDB = nil

func NewOAuthRedis() repositories.RedisAdapter {
	if oauthDB != nil {
		return oauthDB
	}

	db := &redisDB{
		client: native.NewClient(&native.Options{
			Addr:     config.Env.OauthRedisHostName,
			Password: "",
			DB:       int(config.Env.OauthRedisDatabase),
		}),
	}

	oauthDB = db

	return oauthDB
}
