package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/hiroyky/famiphoto/interfaces/http/graph/generated"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/model"
)

func (r *groupResolver) UserPagination(ctx context.Context, obj *model.Group) (*model.UserPagination, error) {
	panic(fmt.Errorf("not implemented"))
}

// Group returns generated.GroupResolver implementation.
func (r *Resolver) Group() generated.GroupResolver { return &groupResolver{r} }

type groupResolver struct{ *Resolver }
