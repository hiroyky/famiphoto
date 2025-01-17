package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.63

import (
	"context"
	"time"

	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/entities"
	fperrors "github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/generated"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	client, ok := ctx.Value(config.OauthClientKey).(*entities.OauthClient)
	if !ok || client.Scope != entities.OauthScopeAdmin {
		return nil, fperrors.New(fperrors.ForbiddenError, nil)
	}

	user, err := r.userUseCase.CreateUser(ctx, input.UserID, input.Name, input.Password, time.Now())
	if err != nil {
		return nil, err
	}
	return model.NewUser(user), nil
}

// CreateOauthClient is the resolver for the createOauthClient field.
func (r *mutationResolver) CreateOauthClient(ctx context.Context, input model.CreateOauthClientInput) (*model.OauthClient, error) {
	oauthClient, secret, err := r.oauthClientUseCase.CreateOauthClient(ctx, &entities.OauthClient{
		OauthClientID: input.ClientID,
		Name:          input.Name,
		Scope:         input.Scope.ToEntity(),
		ClientType:    input.ClientType.ToEntity(),
		RedirectURLs:  input.RedirectUrls,
	})
	if err != nil {
		return nil, err
	}

	return model.NewOauthClientWithSecret(oauthClient, secret), nil
}

// IndexingPhotos is the resolver for the indexingPhotos field.
func (r *mutationResolver) IndexingPhotos(ctx context.Context, input *model.IndexingPhotosInput) (bool, error) {
	if _, ok := ctx.Value(config.ClientSessionKey).(*entities.OauthSession); !ok {
		return false, fperrors.New(fperrors.UserUnauthorizedError, nil)
	}

	if err := r.photoImportUseCase.ExecuteBatch(ctx, input.Fast); err != nil {
		return false, err
	}
	return true, nil
}

// UploadPhoto is the resolver for the uploadPhoto field.
func (r *mutationResolver) UploadPhoto(ctx context.Context) (*model.PhotoUploadInfo, error) {
	sess, ok := ctx.Value(config.ClientSessionKey).(*entities.OauthSession)
	if !ok {
		return nil, fperrors.New(fperrors.UserUnauthorizedError, nil)
	}

	client, ok := ctx.Value(config.OauthClientKey).(*entities.OauthClient)
	if !ok && client != nil { // TODO クライアントアプリの権限チェック
		return nil, fperrors.New(fperrors.ForbiddenError, nil)
	}

	sign, err := r.photoImportUseCase.GenerateUploadURL(ctx, sess.UserID, time.Now())
	if err != nil {
		return nil, err
	}
	return model.NewPhotoUploadInfo(sign), nil
}

// UpdateMe is the resolver for the updateMe field.
func (r *mutationResolver) UpdateMe(ctx context.Context, input model.UpdateMeInput) (*model.User, error) {
	sess, ok := ctx.Value(config.ClientSessionKey).(*entities.OauthSession)
	if !ok {
		return nil, fperrors.New(fperrors.UserUnauthorizedError, nil)
	}
	user, err := r.userUseCase.UpdateUserProfile(ctx, sess.UserID, input.Name)
	if err != nil {
		return nil, err
	}
	return model.NewUser(user), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
