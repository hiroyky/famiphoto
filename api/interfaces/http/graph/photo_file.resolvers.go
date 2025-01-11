package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.63

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

// PhotoFile returns generated.PhotoFileResolver implementation.
func (r *Resolver) PhotoFile() generated.PhotoFileResolver { return &photoFileResolver{r} }

type photoFileResolver struct{ *Resolver }
