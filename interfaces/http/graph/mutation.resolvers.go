package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
