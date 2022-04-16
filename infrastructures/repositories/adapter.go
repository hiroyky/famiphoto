package repositories

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"os"
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

type SambaAdapter interface {
	CreateFile(filePath string, data []byte) error
	CreateDir(dirPath string, perm os.FileMode) error
	Rename(old, file string) error
	ReadDir(dirPath string) ([]os.FileInfo, error)
	ReadFile(filePath string) ([]byte, error)
	Glob(pattern string) ([]string, error)
	Exist(filePath string) bool
	Delete(filePath string) error
	DeleteAll(path string) error
}
