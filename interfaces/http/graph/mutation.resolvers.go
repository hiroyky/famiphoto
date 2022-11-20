package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/errors"
	"time"

	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/generated"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	client, ok := ctx.Value(config.OauthClientKey).(*entities.OauthClient)
	if !ok || client.Scope != entities.OauthScopeAdmin {
		return nil, errors.New(errors.ForbiddenError, nil)
	}

	user, err := r.userUseCase.CreateUser(ctx, input.UserID, input.Name, input.Password, time.Now())
	if err != nil {
		return nil, err
	}
	return model.NewUser(user), nil
}

// CreateGroup is the resolver for the createGroup field.
func (r *mutationResolver) CreateGroup(ctx context.Context, input model.CreateGroupInput) (*model.Group, error) {
	panic(fmt.Errorf("not implemented"))
}

// CreateOauthClient is the resolver for the createOauthClient field.
func (r *mutationResolver) CreateOauthClient(ctx context.Context, input model.CreateOauthClientInput) (*model.OauthClient, error) {
	oauthClient, secret, err := r.oauthClientUseCase.CreateOauthClient(ctx, &entities.OauthClient{
		OauthClientID: input.ClientID,
		Name:          input.Name,
		Scope:         input.Scope.ToEntity(),
		ClientType:    input.ClientType.ToEntity(),
		RedirectURLs:  input.RedirectUrls,
	})
	if err != nil {
		return nil, err
	}

	return model.NewOauthClientWithSecret(oauthClient, secret), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
