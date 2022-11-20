package usecases

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/infrastructures"
)

type GroupUseCase interface {
	GetGroup(ctx context.Context, groupID string) (*entities.Group, error)
	ExistGroup(ctx context.Context, groupID string) (bool, error)
	GetUserBelongingGroups(ctx context.Context, userID string) ([]*entities.Group, error)
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
