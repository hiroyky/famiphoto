package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/hiroyky/famiphoto/interfaces/http/graph/generated"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/model"
	"github.com/hiroyky/famiphoto/utils/gql"
)

// Photo is the resolver for the photo field.
func (r *photoFileResolver) Photo(ctx context.Context, obj *model.PhotoFile) (*model.Photo, error) {
	photoID, err := gql.DecodeIntID(obj.PhotoID)
	if err != nil {
		return nil, err
	}
	photo, err := r.searchUseCase.SearchPhotoByPhotoID(ctx, photoID)
	if err != nil {
		return nil, err
	}
	return model.NewPhoto(photo), nil
}

// Group is the resolver for the group field.
func (r *photoFileResolver) Group(ctx context.Context, obj *model.PhotoFile) (*model.Group, error) {
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

// Owner is the resolver for the owner field.
func (r *photoFileResolver) Owner(ctx context.Context, obj *model.PhotoFile) (*model.User, error) {
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

// PhotoFile returns generated.PhotoFileResolver implementation.
func (r *Resolver) PhotoFile() generated.PhotoFileResolver { return &photoFileResolver{r} }

type photoFileResolver struct{ *Resolver }
