package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/hiroyky/famiphoto/interfaces/http/graph/generated"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/model"
)

func (r *userResolver) Password(ctx context.Context, obj *model.User) (*model.UserPassword, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) BelongGroups(ctx context.Context, obj *model.User) ([]*model.Group, error) {
	panic(fmt.Errorf("not implemented"))
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
