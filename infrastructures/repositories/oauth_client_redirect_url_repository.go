package repositories

import (
	"context"
	"fmt"
	"github.com/hiroyky/famiphoto/drivers/mysql"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/infrastructures/dbmodels"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type OAuthClientRedirectURLRepository interface {
	GetOAuthClientRedirectURLsByOAuthClientID(ctx context.Context, oauthClientID string) ([]*dbmodels.OauthClientRedirectURL, error)
	CreateOAuthClientRedirectURL(ctx context.Context, url *entities.OAuthClientRedirectURL) (*dbmodels.OauthClientRedirectURL, error)
}

func NewOauthClientRedirectURLRepository(db mysql.SQLExecutor) OAuthClientRedirectURLRepository {
	return &oauthClientRedirectURLRepository{db: db}
}

type oauthClientRedirectURLRepository struct {
	db mysql.SQLExecutor
}

func (r *oauthClientRedirectURLRepository) GetOAuthClientRedirectURLsByOAuthClientID(ctx context.Context, oauthClientID string) ([]*dbmodels.OauthClientRedirectURL, error) {
	urls, err := dbmodels.OauthClientRedirectUrls(
		qm.Where(fmt.Sprintf("%s = ?", dbmodels.OauthClientRedirectURLColumns.OauthClientID), oauthClientID),
	).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return urls, nil
}

func (r *oauthClientRedirectURLRepository) CreateOAuthClientRedirectURL(ctx context.Context, url *entities.OAuthClientRedirectURL) (*dbmodels.OauthClientRedirectURL, error) {
	data := &dbmodels.OauthClientRedirectURL{
		OauthClientID:            url.OauthClientID,
		RedirectURL:              url.RedirectURL,
		OauthClientRedirectURLID: url.OAuthClientRedirectUrlID,
	}

	if err := data.Insert(ctx, r.db, boil.Infer()); err != nil {
		return nil, err
	}

	return data, nil
}
