package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/hiroyky/famiphoto/interfaces/http/graph/generated"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/model"
	"github.com/hiroyky/famiphoto/utils/gql"
)

// Owner is the resolver for the owner field.
func (r *photoResolver) Owner(ctx context.Context, obj *model.Photo) (*model.User, error) {
	userID, err := gql.DecodeStrID(obj.OwnerID)
	if err != nil {
		return nil, err
	}
	user, err := r.userUseCase.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return model.NewUser(user), nil
}

// Group is the resolver for the group field.
func (r *photoResolver) Group(ctx context.Context, obj *model.Photo) (*model.Group, error) {
	groupID, err := gql.DecodeStrID(obj.GroupID)
	if err != nil {
		return nil, err
	}
	group, err := r.groupUseCase.GetGroup(ctx, groupID)
	if err != nil {
		return nil, err
	}

	return model.NewGroup(group), nil
}

// ExifData is the resolver for the exifData field.
func (r *photoResolver) ExifData(ctx context.Context, obj *model.Photo) ([]*model.PhotoExif, error) {
	panic(fmt.Errorf("not implemented"))
}

// Files is the resolver for the files field.
func (r *photoResolver) Files(ctx context.Context, obj *model.Photo) ([]*model.PhotoFile, error) {
	panic(fmt.Errorf("not implemented"))
}

// Photo returns generated.PhotoResolver implementation.
func (r *Resolver) Photo() generated.PhotoResolver { return &photoResolver{r} }

type photoResolver struct{ *Resolver }
