package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/generated"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/model"
	"github.com/hiroyky/famiphoto/utils/gql"
)

// Password is the resolver for the password field.
func (r *userResolver) Password(ctx context.Context, obj *model.User) (*model.UserPassword, error) {
	userID, err := gql.DecodeStrID(obj.ID)
	if err != nil {
		return nil, err
	}
	password, err := r.userUseCase.GetUserPassword(ctx, userID)
	if err != nil {
		return nil, err
	}
	return model.NewUserPassword(password), nil
}

// BelongGroups is the resolver for the belongGroups field.
func (r *userResolver) BelongGroups(ctx context.Context, obj *model.User) ([]*model.Group, error) {
	groupID, err := gql.DecodeStrID(obj.ID)
	if err != nil {
		return nil, err
	}
	groups, err := r.groupUseCase.GetUserBelongingGroups(ctx, groupID)
	if err != nil {
		return nil, err
	}
	return model.NewGroups(groups), nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
