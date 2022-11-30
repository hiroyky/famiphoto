package infrastructures

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/infrastructures/dbmodels"
	"github.com/hiroyky/famiphoto/infrastructures/repositories"
	"github.com/hiroyky/famiphoto/utils/array"
)

type GroupAdapter interface {
	GetGroup(ctx context.Context, groupID string) (*entities.Group, error)
	ExistGroup(ctx context.Context, groupID string) (bool, error)
	GetGroupsByUserID(ctx context.Context, userID string) ([]*entities.Group, error)
	IsBelongGroupUser(ctx context.Context, groupID, userID string) (bool, error)
	CreateGroup(ctx context.Context, group *entities.Group, userID string) error
	EditGroupMembers(ctx context.Context, groupID string, appendUserIDs, deleteUserIDs []string) error
}

func NewGroupAdapter(
	groupRepo repositories.GroupRepository,
	groupUserRepo repositories.GroupUserRepository,
	photoStorageRepo repositories.PhotoStorageRepository,
) GroupAdapter {
	return &groupAdapter{
		groupRepo:        groupRepo,
		groupUserRepo:    groupUserRepo,
		photoStorageRepo: photoStorageRepo,
	}
}

type groupAdapter struct {
	groupRepo        repositories.GroupRepository
	groupUserRepo    repositories.GroupUserRepository
	photoStorageRepo repositories.PhotoStorageRepository
}

func (a *groupAdapter) GetGroup(ctx context.Context, groupID string) (*entities.Group, error) {
	dbGroup, err := a.groupRepo.GetGroup(ctx, groupID)
	if err != nil {
		return nil, err
	}
	return a.toGroupEntity(dbGroup), nil
}

func (a *groupAdapter) ExistGroup(ctx context.Context, groupID string) (bool, error) {
	return a.groupRepo.ExistGroup(ctx, groupID)
}

func (a *groupAdapter) GetGroupsByUserID(ctx context.Context, userID string) ([]*entities.Group, error) {
	dbGroups, err := a.groupRepo.GetGroupsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return array.Map(dbGroups, a.toGroupEntity), nil
}

func (a *groupAdapter) IsBelongGroupUser(ctx context.Context, groupID, userID string) (bool, error) {
	return a.groupRepo.ExistGroupUser(ctx, groupID, userID)
}

func (a *groupAdapter) CreateGroup(ctx context.Context, group *entities.Group, userID string) error {
	if err := a.photoStorageRepo.CreateGroupUserDir(group.GroupID, userID); err != nil {
		return err
	}

	groupModel := &dbmodels.Group{
		GroupID: group.GroupID,
		Name:    group.Name,
	}
	groupUser := &dbmodels.GroupUser{
		GroupID: group.GroupID,
		UserID:  userID,
	}
	if err := a.groupRepo.CreateGroup(ctx, groupModel, groupUser); err != nil {
		return err
	}
	return nil
}

func (a *groupAdapter) toGroupEntity(group *dbmodels.Group) *entities.Group {
	return &entities.Group{
		GroupID: group.GroupID,
		Name:    group.Name,
	}
}

func (a *groupAdapter) EditGroupMembers(ctx context.Context, groupID string, appendUserIDs, deleteUserIDs []string) error {
	if appendUserIDs != nil {
		appendGroupUsers := make(dbmodels.GroupUserSlice, len(appendUserIDs))
		for i, v := range appendUserIDs {
			appendGroupUsers[i] = &dbmodels.GroupUser{
				GroupID: groupID,
				UserID:  v,
			}
		}
		if err := a.groupUserRepo.CreateGroupUsers(ctx, appendGroupUsers); err != nil {
			return nil
		}
	}

	if deleteUserIDs != nil {
		deleteGroupUsers := make(dbmodels.GroupUserSlice, len(deleteUserIDs))
		for i, v := range deleteUserIDs {
			deleteGroupUsers[i] = &dbmodels.GroupUser{
				GroupID: groupID,
				UserID:  v,
			}
		}
		if err := a.groupUserRepo.DeleteGroupUsers(ctx, deleteGroupUsers); err != nil {
			return err
		}
	}

	return nil
}
