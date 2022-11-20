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
	ExistGroup(ctx context.Context, groupID string) (bool, error)
	GetGroupsByUserID(ctx context.Context, userID string) (dbmodels.GroupSlice, error)
	GetUsersByGroupID(ctx context.Context, groupID string, limit, offset int) (dbmodels.UserSlice, error)
	CountUsersByGroupID(ctx context.Context, groupID string) (int, error)
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

func (r *groupRepository) ExistGroup(ctx context.Context, groupID string) (bool, error) {
	return dbmodels.GroupExists(ctx, r.db, groupID)
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

func (r *groupRepository) GetUsersByGroupID(ctx context.Context, groupID string, limit, offset int) (dbmodels.UserSlice, error) {
	userGroups, err := dbmodels.GroupUsers(
		qm.Load(qm.Rels(dbmodels.GroupUserRels.User)),
		qm.Where("group_id = ?", groupID),
		qm.Limit(limit),
		qm.Offset(offset),
	).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	users := make(dbmodels.UserSlice, len(userGroups))
	for i, ug := range userGroups {
		users[i] = ug.R.User
	}
	return users, nil
}

func (r *groupRepository) CountUsersByGroupID(ctx context.Context, groupID string) (int, error) {
	val, err := dbmodels.GroupUsers(qm.Where("group_id = ?", groupID)).Count(ctx, r.db)
	if err != nil {
		return 0, err
	}
	return int(val), nil
}
