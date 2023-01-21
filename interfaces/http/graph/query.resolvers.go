package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"errors"

	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/entities"
	fperrors "github.com/hiroyky/famiphoto/errors"
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

// ExistUserID is the resolver for the existUserId field.
func (r *queryResolver) ExistUserID(ctx context.Context, id string) (bool, error) {
	return r.userUseCase.ExistUser(ctx, id)
}

// Group is the resolver for the group field.
func (r *queryResolver) Group(ctx context.Context, id string) (*model.Group, error) {
	groupID, err := gql.DecodeStrID(id)
	if err != nil {
		return nil, err
	}
	group, err := r.groupUseCase.GetGroup(ctx, groupID)
	if err != nil {
		return nil, err
	}
	return model.NewGroup(group), nil
}

// BelongingGroups is the resolver for the belongingGroups field.
func (r *queryResolver) BelongingGroups(ctx context.Context) ([]*model.Group, error) {
	sess, ok := ctx.Value(config.ClientSessionKey).(*entities.OauthSession)
	if !ok {
		return nil, fperrors.New(fperrors.UserUnauthorizedError, nil)
	}

	groups, err := r.groupUseCase.GetUserBelongingGroups(ctx, sess.UserID)
	if err != nil {
		return nil, err
	}

	return model.NewGroups(groups), nil
}

// IsBelongingGroup is the resolver for the isBelongingGroup field.
func (r *queryResolver) IsBelongingGroup(ctx context.Context, id string) (bool, error) {
	sess, ok := ctx.Value(config.ClientSessionKey).(*entities.OauthSession)
	if !ok {
		return false, fperrors.New(fperrors.UserUnauthorizedError, nil)
	}
	groupID, err := gql.DecodeStrID(id)
	if err != nil {
		return false, err
	}

	return r.groupUseCase.IsBelongingGroup(ctx, groupID, sess.UserID)
}

// ExistGroupID is the resolver for the existGroupId field.
func (r *queryResolver) ExistGroupID(ctx context.Context, id string) (bool, error) {
	return r.groupUseCase.ExistGroup(ctx, id)
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	sess, ok := ctx.Value(config.ClientSessionKey).(*entities.OauthSession)
	if !ok {
		return nil, errors.New("invalid session")
	}
	user, err := r.userUseCase.GetUser(ctx, sess.UserID)
	if err != nil {
		return nil, err
	}
	return model.NewUser(user), nil
}

// Photo is the resolver for the photo field.
func (r *queryResolver) Photo(ctx context.Context, id string) (*model.Photo, error) {
	sess, ok := ctx.Value(config.ClientSessionKey).(*entities.OauthSession)
	if !ok {
		return nil, fperrors.New(fperrors.UserUnauthorizedError, nil)
	}

	photoID, err := gql.DecodeIntID(id)
	if err != nil {
		return nil, err
	}
	photo, err := r.searchUseCase.SearchPhotoByPhotoID(ctx, photoID, sess.UserID)
	if err != nil {
		return nil, err
	}
	return model.NewPhoto(photo), nil
}

// Photos is the resolver for the photos field.
func (r *queryResolver) Photos(ctx context.Context, groupID string, id *string, ownerID *string, limit *int, offset *int) (*model.PhotoPagination, error) {
	sess, ok := ctx.Value(config.ClientSessionKey).(*entities.OauthSession)
	if !ok {
		return nil, fperrors.New(fperrors.UserUnauthorizedError, nil)
	}

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
	groupDecodedID, err := gql.DecodeStrID(groupID)
	if err != nil {
		return nil, err
	}
	result, err := r.searchUseCase.SearchPhotos(ctx, groupDecodedID, sess.UserID, decodedID, ownerDecodedID, dstLimit, dstOffset)
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
