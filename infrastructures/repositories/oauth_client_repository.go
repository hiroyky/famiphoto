package repositories

import (
	"context"
	"database/sql"
	"github.com/hiroyky/famiphoto/drivers/mysql"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/dbmodels"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type OAuthClientRepository interface {
	GetByOauthClientID(ctx context.Context, id string) (*dbmodels.OauthClient, error)
	ExistOauthClient(ctx context.Context, id string) (bool, error)
	CreateOAuthClient(ctx context.Context, client *dbmodels.OauthClient, redirectURLs []*dbmodels.OauthClientRedirectURL) (*dbmodels.OauthClient, []*dbmodels.OauthClientRedirectURL, error)
}

func NewOAuthClientRepository(db mysql.SQLExecutor) OAuthClientRepository {
	return &oauthClientRepository{db: db}
}

type oauthClientRepository struct {
	db mysql.SQLExecutor
}

func (r *oauthClientRepository) GetByOauthClientID(ctx context.Context, id string) (*dbmodels.OauthClient, error) {
	oa, err := dbmodels.FindOauthClient(ctx, r.db, id)
	if err == sql.ErrNoRows {
		return nil, errors.New(errors.OAuthClientNotFoundError, err)
	}
	if err != nil {
		return nil, err
	}
	return oa, nil
}

func (r *oauthClientRepository) ExistOauthClient(ctx context.Context, id string) (bool, error) {
	return dbmodels.OauthClientExists(ctx, r.db, id)
}

func (r *oauthClientRepository) CreateOAuthClient(ctx context.Context, client *dbmodels.OauthClient, redirectURLs []*dbmodels.OauthClientRedirectURL) (*dbmodels.OauthClient, []*dbmodels.OauthClientRedirectURL, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, errors.New(errors.TxnBeginFatal, err)
	}

	if err := client.Insert(ctx, tx, boil.Infer()); err != nil {
		return nil, nil, err
	}

	urls := make([]*dbmodels.OauthClientRedirectURL, len(redirectURLs))
	for i, u := range redirectURLs {
		if err := u.Insert(ctx, tx, boil.Infer()); err != nil {
			return nil, nil, err
		}
		urls[i] = u
	}

	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, nil, errors.New(errors.TxnRollbackFatal, err)
		}
		return nil, nil, errors.New(errors.OAuthClientCreateFatal, err)
	}

	return client, urls, nil
}
