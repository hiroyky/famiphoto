package repositories

import (
	"context"
	"fmt"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/infrastructures/models"
	"github.com/hiroyky/famiphoto/utils/cast"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type oauthClientRedirectURLRepository struct {
	db SQLExecutor
}

func (r *oauthClientRedirectURLRepository) GetOAuthClientRedirectURLsByOAuthClientID(ctx context.Context, oauthClientID string) ([]*entities.OAuthClientRedirectURL, error) {
	urls, err := models.OauthClientRedirectUrls(
		qm.Where(fmt.Sprintf("%s = ?", models.OauthClientRedirectURLColumns.OauthClientID), oauthClientID),
	).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return cast.Array(urls, r.toEntity), nil
}

func (r *oauthClientRedirectURLRepository) CreatOAuthClientRedirectURL(ctx context.Context, url *entities.OAuthClientRedirectURL) (*entities.OAuthClientRedirectURL, error) {
	data := &models.OauthClientRedirectURL{
		OauthClientID:            url.OauthClientID,
		RedirectURL:              url.RedirectURL,
		OauthClientRedirectURLID: url.OAuthClientRedirectUrlID,
	}

	if err := data.Insert(ctx, r.db, boil.Infer()); err != nil {
		return nil, err
	}

	return r.toEntity(data), nil
}

func (r *oauthClientRedirectURLRepository) toEntity(t *models.OauthClientRedirectURL) *entities.OAuthClientRedirectURL {
	return &entities.OAuthClientRedirectURL{
		OAuthClientRedirectUrlID: t.OauthClientRedirectURLID,
		OauthClientID:            t.OauthClientID,
		RedirectURL:              t.RedirectURL,
	}
}
