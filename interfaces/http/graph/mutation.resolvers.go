package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/generated"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/model"
)

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
