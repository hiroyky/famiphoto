package redis

import (
	"context"
	native "github.com/go-redis/redis/v8"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/errors"
	"time"
)

type Driver interface {
	Get(ctx context.Context, key string) (string, error)
	GetDel(ctx context.Context, key string) (string, error)
	SetEx(ctx context.Context, key string, val interface{}, expiration time.Duration) error
	Del(ctx context.Context, key string) error
	SAdd(ctx context.Context, key string, members ...string) error
	SMembers(ctx context.Context, key string) ([]string, error)
	SRem(ctx context.Context, key string, members ...string) error
}

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

func (r *redisDB) Del(ctx context.Context, key string) error {
	err := r.client.Del(ctx, key).Err()
	if err == nil {
		return nil
	}
	if err == native.Nil {
		return errors.New(errors.RedisKeyNotFound, err)
	}
	return errors.New(errors.RedisFatal, err)
}

func (r *redisDB) SAdd(ctx context.Context, key string, members ...string) error {
	return r.client.SAdd(ctx, key, members).Err()
}

func (r *redisDB) SMembers(ctx context.Context, key string) ([]string, error) {
	values, err := r.client.SMembers(ctx, key).Result()
	if err == nil {
		return values, nil
	}
	if err == native.Nil {
		return nil, errors.New(errors.RedisKeyNotFound, err)
	}
	return nil, err
}

func (r *redisDB) SRem(ctx context.Context, key string, members ...string) error {
	err := r.client.SRem(ctx, key, members).Err()
	if err == nil {
		return nil
	}
	if err == native.Nil {
		return errors.New(errors.RedisKeyNotFound, err)
	}
	return err
}

var oauthDB *redisDB = nil

func NewRedisDriver(client *native.Client) Driver {
	return &redisDB{client: client}
}

func NewOauthRedis() Driver {
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
