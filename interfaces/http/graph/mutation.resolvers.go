package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/entities"
	fperrors "github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/generated"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/model"
	"github.com/hiroyky/famiphoto/utils/gql"
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

// CreateGroup is the resolver for the createGroup field.
func (r *mutationResolver) CreateGroup(ctx context.Context, input model.CreateGroupInput) (*model.Group, error) {
	sess, ok := ctx.Value(config.ClientSessionKey).(*entities.OauthSession)
	if !ok {
		return nil, fperrors.New(fperrors.UserUnauthorizedError, nil)
	}
	group, err := r.groupUseCase.CreateGroup(ctx, input.GroupID, input.Name, sess.UserID)
	if err != nil {
		return nil, err
	}

	return model.NewGroup(group), nil
}

// AlterGroupMembers is the resolver for the alterGroupMembers field.
func (r *mutationResolver) AlterGroupMembers(ctx context.Context, input model.AlterGroupMembersInput) (*model.Group, error) {
	sess, ok := ctx.Value(config.ClientSessionKey).(*entities.OauthSession)
	if !ok {
		return nil, fperrors.New(fperrors.UserUnauthorizedError, nil)
	}

	groupID, err := gql.DecodeStrID(input.GroupID)
	if err != nil {
		return nil, err
	}
	appendUserIDs, err := gql.DecodeStrIDs(input.AppendUserIds)
	if err != nil {
		return nil, err
	}
	removeUserIDs, err := gql.DecodeStrIDs(input.RemoveUserIds)
	if err != nil {
		return nil, err
	}

	group, err := r.groupUseCase.AlterGroupMembers(ctx, sess.UserID, groupID, appendUserIDs, removeUserIDs)
	if err != nil {
		return nil, err
	}
	return model.NewGroup(group), nil
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
