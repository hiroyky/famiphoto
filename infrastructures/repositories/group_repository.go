package repositories

import (
	"context"
	"database/sql"
	"github.com/hiroyky/famiphoto/drivers/mysql"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures/dbmodels"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type GroupRepository interface {
	GetGroup(ctx context.Context, groupID string) (*dbmodels.Group, error)
	GetGroupsByUserID(ctx context.Context, userID string) (dbmodels.GroupSlice, error)
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

func (r *groupRepository) GetGroupsByUserID(ctx context.Context, userID string) (dbmodels.GroupSlice, error) {
	userGroups, err := dbmodels.GroupUsers(
		qm.Load(qm.Rels(dbmodels.GroupUserRels.Group)),
		qm.Where("user_id = ?", userID),
	).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	groups := make(dbmodels.GroupSlice, len(userGroups))
	for i, ug := range userGroups {
		groups[i] = ug.R.Group
	}

	return groups, nil
}
