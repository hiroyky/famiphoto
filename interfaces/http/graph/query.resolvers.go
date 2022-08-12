package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/generated"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/model"
	"github.com/hiroyky/famiphoto/utils/gql"
	"github.com/hiroyky/famiphoto/utils/pagination"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	userID, err := gql.DecodeStrID(id)
	if err != nil {
		return nil, err
	}
	user, err := r.userUseCase.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return model.NewUser(user), nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, id *string, limit *int, offset *int) (*model.UserPagination, error) {
	userID, err := gql.DecodeStrIDPtr(id)
	if err != nil {
		return nil, err
	}

	dstLimit := pagination.GetLimitOrDefault(limit, 20, 100)
	dstOffset := pagination.GetOffsetOrDefault(offset)
	users, total, err := r.userUseCase.GetUsers(ctx, userID, dstLimit, dstOffset)
	if err != nil {
		return nil, err
	}
	return model.NewUserPagination(users, total, dstLimit, dstOffset), nil
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	sess, ok := ctx.Value(config.ClientSessionKey).(*entities.OauthSession)
	if !ok {
		return nil, errors.New("")
	}
	user, err := r.userUseCase.GetUser(ctx, sess.UserID)
	if err != nil {
		return nil, err
	}
	return model.NewUser(user), nil
}

// Photo is the resolver for the photo field.
func (r *queryResolver) Photo(ctx context.Context, id string) (*model.Photo, error) {
	photoID, err := gql.DecodeIntID(id)
	if err != nil {
		return nil, err
	}
	photo, err := r.searchUseCase.SearchPhotoByPhotoID(ctx, photoID)
	if err != nil {
		return nil, err
	}
	return model.NewPhoto(photo), nil
}

// Photos is the resolver for the photos field.
func (r *queryResolver) Photos(ctx context.Context, id *string, ownerID *string, groupID *string, limit *int, offset *int) (*model.PhotoPagination, error) {
	dstLimit := pagination.GetLimitOrDefault(limit, 20, 100)
	dstOffset := pagination.GetOffsetOrDefault(offset)
	decodedID, err := gql.DecodeIntIDPtr(id)
	if err != nil {
		return nil, err
	}
	ownerDecodedID, err := gql.DecodeStrIDPtr(ownerID)
	if err != nil {
		return nil, err
	}
	groupDecodedID, err := gql.DecodeStrIDPtr(groupID)
	if err != nil {
		return nil, err
	}
	result, err := r.searchUseCase.SearchPhotos(ctx, decodedID, ownerDecodedID, groupDecodedID, dstLimit, dstOffset)
	if err != nil {
		return nil, err
	}

	return model.NewPhotoPagination(result, dstLimit, dstOffset), nil
}

// PhotoFile is the resolver for the photoFile field.
func (r *queryResolver) PhotoFile(ctx context.Context, id string) (*model.PhotoFile, error) {
	decodedID, err := gql.DecodeIntID(id)
	if err != nil {
		return nil, err
	}
	photoFile, err := r.photoUseCase.GetPhotoFileByPhotoFileID(ctx, decodedID)
	if err != nil {
		return nil, err
	}
	return model.NewPhotoFile(photoFile), nil
}

// PhotoFiles is the resolver for the photoFiles field.
func (r *queryResolver) PhotoFiles(ctx context.Context, photoID string) ([]*model.PhotoFile, error) {
	decodedID, err := gql.DecodeIntID(photoID)
	if err != nil {
		return nil, err
	}
	photoFiles, err := r.photoUseCase.GetPhotoFilesByPhotoID(ctx, decodedID)
	if err != nil {
		return nil, err
	}
	return model.NewPhotoFiles(photoFiles), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
