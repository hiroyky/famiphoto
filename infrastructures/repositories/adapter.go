package repositories

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"time"
)

type SQLExecutor interface {
	boil.ContextExecutor
	boil.ContextBeginner
}

type RedisAdapter interface {
	Get(ctx context.Context, key string) (string, error)
	GetDel(ctx context.Context, key string) (string, error)
	SetEx(ctx context.Context, key string, val interface{}, expiration time.Duration) error
}
