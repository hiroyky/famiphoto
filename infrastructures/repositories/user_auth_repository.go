package repositories

import (
	"context"
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"fmt"
	"github.com/hiroyky/famiphoto/drivers/mysql"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/dbmodels"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func NewUserAuthRepository(db mysql.SQLExecutor) UserAuthRepository {
	return &userAuthRepository{db: db}
}

type UserAuthRepository interface {
	UpsertUserAuth(ctx context.Context, m *dbmodels.UserAuth, refreshTokenRaw string) (*dbmodels.UserAuth, error)
	GetUserAuth(ctx context.Context, userID, clientID string) (*dbmodels.UserAuth, error)
	GetUserAuthByRefreshToken(ctx context.Context, refreshToken string) (*dbmodels.UserAuth, error)
	DeleteUserAuth(ctx context.Context, userID, clientID string) error
	DeleteClientAllAuth(ctx context.Context, clientID string) error
}

type userAuthRepository struct {
	db mysql.SQLExecutor
}

func (r *userAuthRepository) UpsertUserAuth(ctx context.Context, m *dbmodels.UserAuth, refreshTokenRaw string) (*dbmodels.UserAuth, error) {
	m.RefreshToken = r.refreshTokenHashed(refreshTokenRaw)
	if err := m.Upsert(ctx, r.db, boil.Infer(), boil.Infer()); err != nil {
		return nil, err
	}
	return m, nil
}

func (r *userAuthRepository) GetUserAuth(ctx context.Context, userID, clientID string) (*dbmodels.UserAuth, error) {
	row, err := dbmodels.FindUserAuth(ctx, r.db, userID, clientID)
	if err != nil {
		return nil, errors.New(errors.UserAuthNotFoundError, err)
	}
	return row, nil
}

func (r *userAuthRepository) GetUserAuthByRefreshToken(ctx context.Context, refreshToken string) (*dbmodels.UserAuth, error) {
	hashed := r.refreshTokenHashed(refreshToken)
	ua, err := dbmodels.UserAuths(qm.Where(fmt.Sprintf("%s=?", dbmodels.UserAuthColumns.RefreshToken), hashed)).One(ctx, r.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(errors.UserAuthNotFoundError, err)
		}
		return nil, err
	}
	return ua, nil
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

func (r *userAuthRepository) refreshTokenHashed(refreshToken string) string {
	hashed := sha512.Sum512([]byte(refreshToken))
	return base64.StdEncoding.EncodeToString(hashed[:])
}
