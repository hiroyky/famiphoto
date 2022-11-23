package usecases

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/infrastructures"
)

type GroupUseCase interface {
	GetGroup(ctx context.Context, groupID string) (*entities.Group, error)
	ExistGroup(ctx context.Context, groupID string) (bool, error)
	GetUserBelongingGroups(ctx context.Context, userID string) ([]*entities.Group, error)
	CreateGroup(ctx context.Context, groupID, groupName, userID string) (*entities.Group, error)
}

func NewGroupUseCase(groupAdapter infrastructures.GroupAdapter) GroupUseCase {
	return &groupUseCase{groupAdapter: groupAdapter}
}

type groupUseCase struct {
	groupAdapter infrastructures.GroupAdapter
}

func (u *groupUseCase) GetGroup(ctx context.Context, groupID string) (*entities.Group, error) {
	return u.groupAdapter.GetGroup(ctx, groupID)
}

func (u *groupUseCase) ExistGroup(ctx context.Context, groupID string) (bool, error) {
	return u.groupAdapter.ExistGroup(ctx, groupID)
}

func (u *groupUseCase) GetUserBelongingGroups(ctx context.Context, userID string) ([]*entities.Group, error) {
	return u.groupAdapter.GetGroupsByUserID(ctx, userID)
}

func (u *groupUseCase) CreateGroup(ctx context.Context, groupID, groupName, userID string) (*entities.Group, error) {
	if exist, err := u.groupAdapter.ExistGroup(ctx, groupID); err != nil {
		return nil, err
	} else if exist {
		return nil, errors.New(errors.GroupAlreadyExistError, nil)
	}

	if err := u.groupAdapter.CreateGroup(
		ctx,
		&entities.Group{
			GroupID: groupID,
			Name:    groupName,
		},
		userID,
	); err != nil {
		return nil, err
	}

	return u.groupAdapter.GetGroup(ctx, groupID)
}
