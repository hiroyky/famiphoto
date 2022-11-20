package model

import (
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/utils/array"
	"github.com/hiroyky/famiphoto/utils/gql"
)

const GroupName = "Group"

func NewGroup(g *entities.Group) *Group {
	return &Group{
		ID:      gql.EncodeStrID(GroupName, g.GroupID),
		GroupID: g.GroupID,
		Name:    g.Name,
	}
}

func NewGroups(l []*entities.Group) []*Group {
	return array.Map(l, NewGroup)
}
