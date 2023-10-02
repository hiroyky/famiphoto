package repositories

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/drivers/redis"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/models"
	"time"
)

type PhotoUploadSignRepository interface {
	GetSign(ctx context.Context, token string) (*models.PhotoUploadSign, error)
	SetSignToken(ctx context.Context, token, userID string, expireIn int64) error
}

func NewPhotoUploadSignRepository(db redis.Driver) PhotoUploadSignRepository {
	return &photoUploadSignRepository{
		db:     db,
		prefix: config.Env.UploadTokenHashedPrefix,
	}
}

type photoUploadSignRepository struct {
	db     redis.Driver
	prefix string
}

func (r *photoUploadSignRepository) GetSign(ctx context.Context, token string) (*models.PhotoUploadSign, error) {
	str, err := r.db.Get(ctx, r.toHash(token))
	if err != nil {
		if errors.GetFPErrorCode(err) == errors.RedisKeyNotFoundError {
			return nil, errors.New(errors.PhotoUploadSignNotFoundError, err)
		}
		return nil, err
	}

	var val models.PhotoUploadSign
	if err := json.Unmarshal([]byte(str), &val); err != nil {
		return nil, err
	}
	return &val, nil
}

func (r *photoUploadSignRepository) SetSignToken(ctx context.Context, token, userID string, expireIn int64) error {
	m := &models.PhotoUploadSign{
		UserID: userID,
	}
	val, err := json.Marshal(m)
	if err != nil {
		return err
	}

	return r.db.SetEx(ctx, r.toHash(token), val, time.Duration(expireIn)*time.Second)
}

func (r *photoUploadSignRepository) toHash(token string) string {
	base := fmt.Sprintf("%s-%s", r.prefix, token)

	hashed := sha256.Sum256([]byte(base))
	return base64.StdEncoding.EncodeToString(hashed[:])
}
