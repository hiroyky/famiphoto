package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/hiroyky/famiphoto/interfaces/http/graph/generated"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/model"
)

func (r *photoResolver) Group(ctx context.Context, obj *model.Photo) (*model.Group, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *photoResolver) Owner(ctx context.Context, obj *model.Photo) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Photo returns generated.PhotoResolver implementation.
func (r *Resolver) Photo() generated.PhotoResolver { return &photoResolver{r} }

type photoResolver struct{ *Resolver }
