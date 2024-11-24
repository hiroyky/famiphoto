package repositories

import (
	"context"
	"fmt"
	"github.com/hiroyky/famiphoto/drivers/redis"
	"github.com/hiroyky/famiphoto/infrastructures/models"
	"time"
)

type OauthCodeRepository interface {
	SetCode(ctx context.Context, code string, m *models.OauthCode) error
	GetCode(ctx context.Context, code string) (*models.OauthCode, error)
}

func NewOauthCodeRepository(db redis.Driver) OauthCodeRepository {
	return &oauthCodeRepository{
		db:     db,
		expire: 60 * time.Second,
	}
}

type oauthCodeRepository struct {
	db     redis.Driver
	expire time.Duration
}

func (r *oauthCodeRepository) SetCode(ctx context.Context, code string, m *models.OauthCode) error {
	data, err := m.JSONString()
	if err != nil {
		return err
	}

	return r.db.SetEx(ctx, r.genKey(code), data, r.expire)
}

func (r *oauthCodeRepository) GetCode(ctx context.Context, code string) (*models.OauthCode, error) {
	val, err := r.db.GetDel(ctx, r.genKey(code))
	if err != nil {
		return nil, err
	}

	var m models.OauthCode
	if err := m.BindFromJSON(val); err != nil {
		return nil, err
	}

	return &m, nil
}

func (r *oauthCodeRepository) genKey(code string) string {
	return fmt.Sprintf("oauthcode::%s", code)
}
