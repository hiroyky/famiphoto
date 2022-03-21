package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/dbmodels"
	"github.com/hiroyky/famiphoto/usecases"
	"github.com/hiroyky/famiphoto/utils/cast"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func NewOauthClientRepository(db SQLExecutor) usecases.OauthClientAdapter {
	return &oauthClientRepository{
		db: db,
	}
}

type oauthClientRepository struct {
	db SQLExecutor
}

func (r *oauthClientRepository) GetByOauthClientID(ctx context.Context, id string) (*entities.OauthClient, error) {
	oa, err := dbmodels.FindOauthClient(ctx, r.db, id)
	if err == sql.ErrNoRows {
		return nil, errors.New(errors.OAuthClientNotFoundError, err)
	}
	if err != nil {
		return nil, err
	}
	urls, err := dbmodels.OauthClientRedirectUrls(
		qm.Where(fmt.Sprintf("%s = ?", dbmodels.OauthClientRedirectURLColumns.OauthClientID), oa.OauthClientID),
	).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	return r.toEntityOauthClient(oa, urls), nil
}

func (r *oauthClientRepository) toEntityOauthClient(m *dbmodels.OauthClient, urls []*dbmodels.OauthClientRedirectURL) *entities.OauthClient {
	return &entities.OauthClient{
		OauthClientID:      m.OauthClientID,
		ClientSecretHashed: m.ClientSecret,
		Name:               m.Name,
		Scope:              entities.OauthScope(m.Scope),
		ClientType:         entities.OauthClientType(m.ClientType),
		RedirectURLs: cast.ArrayValues(urls, func(t *dbmodels.OauthClientRedirectURL) string {
			return t.RedirectURL
		}),
	}
}

func (r *oauthClientRepository) ExistOauthClient(ctx context.Context, id string) (bool, error) {
	return dbmodels.OauthClientExists(ctx, r.db, id)
}

func (r *oauthClientRepository) CreateOAuthClient(ctx context.Context, client *entities.OauthClient, secret string) (*entities.OauthClient, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.New(errors.TxnBeginFatal, err)
	}

	c := &dbmodels.OauthClient{
		OauthClientID: client.OauthClientID,
		Name:          client.Name,
		ClientSecret:  secret,
		Scope:         string(client.Scope),
		ClientType:    int(client.ClientType),
	}
	if err := c.Insert(ctx, tx, boil.Infer()); err != nil {
		return nil, err
	}

	urls := make([]*dbmodels.OauthClientRedirectURL, len(client.RedirectURLs))
	for i, url := range client.RedirectURLs {
		u := &dbmodels.OauthClientRedirectURL{
			OauthClientID: client.OauthClientID,
			RedirectURL:   url,
		}
		if err := u.Insert(ctx, tx, boil.Infer()); err != nil {
			return nil, err
		}
		urls[i] = u
	}

	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, errors.New(errors.TxnRollbackFatal, err)
		}
		return nil, errors.New(errors.OAuthClientCreateFatal, err)
	}

	return r.toEntityOauthClient(c, urls), nil
}
