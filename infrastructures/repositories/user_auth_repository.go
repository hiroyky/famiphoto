package repositories

import (
	"context"
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"fmt"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/dbmodels"
	"github.com/hiroyky/famiphoto/usecases"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func NewUserAuthAdapter(db SQLExecutor) usecases.UserAuthAdapter {
	return &userAuthRepository{db: db}
}

type userAuthRepository struct {
	db SQLExecutor
}

func (r *userAuthRepository) UpsertUserAuth(ctx context.Context, m *entities.UserAuth) (*entities.UserAuth, error) {
	userAuth := &dbmodels.UserAuth{
		UserID:                  m.UserID,
		OauthClientID:           m.OAuthClientID,
		RefreshToken:            r.refreshTokenHashed(m.RefreshToken),
		RefreshTokenPublishedAt: m.RefreshTokenPublishedAt,
	}
	if err := userAuth.Upsert(ctx, r.db, boil.Infer(), boil.Infer()); err != nil {
		return nil, err
	}
	return r.toEntity(userAuth), nil
}

func (r *userAuthRepository) GetUserAuth(ctx context.Context, userID, clientID string) (*entities.UserAuth, error) {
	row, err := dbmodels.FindUserAuth(ctx, r.db, userID, clientID)
	if err != nil {
		return nil, errors.New(errors.UserAuthNotFoundError, err)
	}
	return r.toEntity(row), nil
}

func (r *userAuthRepository) GetUserAuthByRefreshToken(ctx context.Context, refreshToken string) (*entities.UserAuth, error) {
	hashed := r.refreshTokenHashed(refreshToken)
	ua, err := dbmodels.UserAuths(qm.Where(fmt.Sprintf("%s=?", dbmodels.UserAuthColumns.RefreshToken), hashed)).One(ctx, r.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(errors.UserAuthNotFoundError, err)
		}
		return nil, err
	}
	return r.toEntity(ua), nil
}

func (r *userAuthRepository) DeleteUserAuth(ctx context.Context, userID, clientID string) error {
	row, err := dbmodels.FindUserAuth(ctx, r.db, userID, clientID)
	if err != nil {
		return err
	}
	if _, err := row.Delete(ctx, r.db); err != nil {
		return err
	}
	return nil
}

func (r *userAuthRepository) DeleteClientAllAuth(ctx context.Context, clientID string) error {
	rows, err := dbmodels.UserAuths(qm.Where(fmt.Sprintf("%s = ?", dbmodels.UserAuthColumns.OauthClientID), clientID)).All(ctx, r.db)
	if err != nil {
		return err
	}
	if _, err := rows.DeleteAll(ctx, r.db); err != nil {
		return err
	}
	return nil
}

func (r *userAuthRepository) toEntity(m *dbmodels.UserAuth) *entities.UserAuth {
	return &entities.UserAuth{
		UserID:                  m.UserID,
		OAuthClientID:           m.OauthClientID,
		RefreshToken:            m.RefreshToken,
		RefreshTokenPublishedAt: m.RefreshTokenPublishedAt,
	}
}

func (r *userAuthRepository) refreshTokenHashed(refreshToken string) string {
	hashed := sha512.Sum512([]byte(refreshToken))
	return base64.StdEncoding.EncodeToString(hashed[:])
}
