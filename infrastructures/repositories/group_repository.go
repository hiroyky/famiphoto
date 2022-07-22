package repositories

import (
	"context"
	"database/sql"
	"github.com/hiroyky/famiphoto/drivers/mysql"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/dbmodels"
)

type GroupRepository interface {
	GetGroup(ctx context.Context, groupID string) (*dbmodels.Group, error)
}

func NewGroupRepository(db mysql.SQLExecutor) GroupRepository {
	return &groupRepository{
		db: db,
	}
}

type groupRepository struct {
	db mysql.SQLExecutor
}

func (r *groupRepository) GetGroup(ctx context.Context, groupID string) (*dbmodels.Group, error) {
	group, err := dbmodels.FindGroup(ctx, r.db, groupID)
	if err == sql.ErrNoRows {
		return nil, errors.New(errors.GroupNotFoundError, err)
	}
	if err != nil {
		return nil, err
	}
	return group, nil
}
