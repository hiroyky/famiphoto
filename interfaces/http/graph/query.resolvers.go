package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"

	"github.com/hiroyky/famiphoto/interfaces/http/graph/generated"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/model"
	"github.com/hiroyky/famiphoto/utils/gql"
	"github.com/hiroyky/famiphoto/utils/pagination"
)

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	userID, err := gql.DecodeStrID(id)
	if err != nil {
		return nil, err
	}
	user, err := r.userUseCase.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return model.NewUser(user), nil
}

func (r *queryResolver) Users(ctx context.Context, id *string, limit *int, offset *int) (*model.UserPagination, error) {
	userID, err := gql.DecodeStrIDPtr(id)
	if err != nil {
		return nil, err
	}

	dstLimit := pagination.GetLimitOrDefault(limit, 20, 100)
	dstOffset := pagination.GetOffsetOrDefault(offset)
	users, total, err := r.userUseCase.GetUsers(ctx, userID, dstLimit, dstOffset)
	return model.NewUserPagination(users, total, dstLimit, dstOffset), nil
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	sess, ok := ctx.Value(config.ClientSessionKey).(*entities.OauthSession)
	if !ok {
		return nil, errors.New(errors.UserUnauthorizedError, nil)
	}
	user, err := r.userUseCase.GetUser(ctx, sess.UserID)
	if err != nil {
		return nil, err
	}
	return model.NewUser(user), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
