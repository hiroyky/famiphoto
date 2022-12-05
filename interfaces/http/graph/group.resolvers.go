package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/hiroyky/famiphoto/interfaces/http/graph/generated"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/model"
	"github.com/hiroyky/famiphoto/utils/gql"
	"github.com/hiroyky/famiphoto/utils/pagination"
)

// UserPagination is the resolver for the userPagination field.
func (r *groupResolver) UserPagination(ctx context.Context, obj *model.Group, limit *int, offset *int) (*model.UserPagination, error) {
	groupID, err := gql.DecodeStrID(obj.ID)
	if err != nil {
		return nil, err
	}

	dstLimit := pagination.GetLimitOrDefault(limit, 20, 100)
	dstOffset := pagination.GetOffsetOrDefault(offset)

	users, total, err := r.userUseCase.GetUsersBelongingGroup(ctx, groupID, dstLimit, dstOffset)
	if err != nil {
		return nil, err
	}
	return model.NewUserPagination(users, total, dstLimit, dstOffset), nil
}

// Group returns generated.GroupResolver implementation.
func (r *Resolver) Group() generated.GroupResolver { return &groupResolver{r} }

type groupResolver struct{ *Resolver }
