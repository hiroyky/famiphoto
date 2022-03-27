package redis

import (
	"context"
	native "github.com/go-redis/redis/v8"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/repositories"
	"time"
)

type redisDB struct {
	client *native.Client
}

func (r *redisDB) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == nil {
		return val, nil
	}
	if err == native.Nil {
		return "", errors.New(errors.RedisKeyNotFound, err)
	}
	return val, errors.New(errors.RedisFatal, err)
}

func (r *redisDB) GetDel(ctx context.Context, key string) (string, error) {
	val, err := r.client.GetDel(ctx, key).Result()
	if err == nil {
		return val, nil
	}
	if err == native.Nil {
		return "", errors.New(errors.RedisKeyNotFound, err)
	}
	return val, nil
}

func (r *redisDB) SetEx(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	err := r.client.SetEX(ctx, key, val, expiration).Err()
	if err != nil {
		return errors.New(errors.RedisFatal, err)
	}
	return nil
}

var oauthDB *redisDB = nil

func NewOauthRedis() repositories.RedisAdapter {
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
