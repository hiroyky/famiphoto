package infrastructures

import (
	"context"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/infrastructures/dbmodels"
	"github.com/hiroyky/famiphoto/infrastructures/repositories"
)

type GroupAdapter interface {
	GetGroup(ctx context.Context, groupID string) (*entities.Group, error)
}

func NewGroupAdapter(groupRepo repositories.GroupRepository) GroupAdapter {
	return &groupAdapter{groupRepo: groupRepo}
}

type groupAdapter struct {
	groupRepo repositories.GroupRepository
}

func (a *groupAdapter) GetGroup(ctx context.Context, groupID string) (*entities.Group, error) {
	dbGroup, err := a.groupRepo.GetGroup(ctx, groupID)
	if err != nil {
		return nil, err
	}
	return a.toGroupEntity(dbGroup), nil
}

func (a *groupAdapter) toGroupEntity(group *dbmodels.Group) *entities.Group {
	return &entities.Group{
		GroupID: group.GroupID,
		Name:    group.Name,
	}
}
