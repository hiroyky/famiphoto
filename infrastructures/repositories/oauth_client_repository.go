package repositories

import (
	"context"
	"database/sql"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/models"
	"github.com/hiroyky/famiphoto/usecases"
	"github.com/volatiletech/sqlboiler/boil"
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
	oa, err := models.FindOauthClient(ctx, r.db, id)
	if err == sql.ErrNoRows {
		return nil, errors.New(errors.OAuthClientNotFoundError, err)
	}
	if err != nil {
		return nil, err
	}
	return r.toEntityOauthClient(oa), nil
}

func (r *oauthClientRepository) toEntityOauthClient(m *models.OauthClient) *entities.OauthClient {
	return &entities.OauthClient{
		OauthClientID: m.OauthClientID,
		Name:          m.Name,
		Scope:         entities.OauthScope(m.Scope),
		ClientType:    entities.OauthClientType(m.ClientType),
	}
}

func (r *oauthClientRepository) ExistOauthClient(ctx context.Context, id string) (bool, error) {
	return models.OauthClientExists(ctx, r.db, id)
}

func (r *oauthClientRepository) CreateOAuthClient(ctx context.Context, client *entities.OauthClient, secret string) (*entities.OauthClient, error) {
	c := &models.OauthClient{
		OauthClientID: client.OauthClientID,
		Name:          client.Name,
		ClientSecret:  secret,
		Scope:         string(client.Scope),
		ClientType:    int(client.ClientType),
	}
	if err := c.Insert(ctx, r.db, boil.Infer()); err != nil {
		return nil, err
	}
	return r.toEntityOauthClient(c), nil
}
