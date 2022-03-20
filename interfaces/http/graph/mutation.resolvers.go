package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/utils/gql"
	"time"

	"github.com/hiroyky/famiphoto/interfaces/http/graph/generated"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	if err := r.userUseCase.ValidateToCreateUser(ctx, input.UserID, input.Name, input.Password); err != nil {
		return nil, err
	}

	user, err := r.userUseCase.CreateUser(ctx, input.UserID, input.Name, input.Password, time.Now())
	if err != nil {
		return nil, err
	}
	return model.NewUser(user), nil
}

func (r *mutationResolver) CreateGroup(ctx context.Context, input model.CreateGroupInput) (*model.Group, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateOauthClient(ctx context.Context, input model.CreateOauthClientInput) (*model.OauthClient, error) {
	clientID, err := gql.DecodeStrID(input.ClientID)
	if err != nil {
		return nil, err
	}

	oauthClient, secret, err := r.oauthClientUseCase.CreateOauthClient(ctx, &entities.OauthClient{
		OauthClientID: clientID,
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
