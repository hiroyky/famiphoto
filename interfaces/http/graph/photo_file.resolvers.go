package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/hiroyky/famiphoto/interfaces/http/graph/generated"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/model"
)

// Photo is the resolver for the photo field.
func (r *photoFileResolver) Photo(ctx context.Context, obj *model.PhotoFile) (*model.Photo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Group is the resolver for the group field.
func (r *photoFileResolver) Group(ctx context.Context, obj *model.PhotoFile) (*model.Group, error) {
	panic(fmt.Errorf("not implemented"))
}

// Owner is the resolver for the owner field.
func (r *photoFileResolver) Owner(ctx context.Context, obj *model.PhotoFile) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// PhotoFile returns generated.PhotoFileResolver implementation.
func (r *Resolver) PhotoFile() generated.PhotoFileResolver { return &photoFileResolver{r} }

type photoFileResolver struct{ *Resolver }
