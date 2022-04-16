package repositories

import (
	"context"
	"fmt"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/infrastructures/models"
	"github.com/hiroyky/famiphoto/usecases"
	"time"
)

func NewOauthCodeAdapter(db RedisAdapter) usecases.OauthCodeAdapter {
	return &oauthCodeAdapter{
		db: db,
	}
}

type oauthCodeAdapter struct {
	db RedisAdapter
}

func (r *oauthCodeAdapter) SetCode(ctx context.Context, code *entities.OAuthCode) error {
	data, err := (&models.OauthCode{
		ClientID:    code.ClientID,
		UserID:      code.UserID,
		RedirectURL: code.RedirectURL,
	}).JSONString()
	if err != nil {
		return err
	}
	return r.db.SetEx(
		ctx,
		r.genKey(code.Code),
		data,
		60*time.Second,
	)
}

func (r *oauthCodeAdapter) GetCode(ctx context.Context, code string) (*entities.OAuthCode, error) {
	val, err := r.db.GetDel(ctx, r.genKey(code))
	if err != nil {
		return nil, err
	}

	var m models.OauthCode
	if err := m.BindFromJSON(val); err != nil {
		return nil, err
	}

	return &entities.OAuthCode{
		Code:        code,
		ClientID:    m.ClientID,
		UserID:      m.UserID,
		RedirectURL: m.RedirectURL,
	}, nil
}

func (r *oauthCodeAdapter) genKey(code string) string {
	return fmt.Sprintf("oauthcode::%s", code)
}
