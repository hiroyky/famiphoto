package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/generated"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/model"
	"github.com/hiroyky/famiphoto/utils/cast"
	"github.com/hiroyky/famiphoto/utils/gql"
)

func (r *oauthClientResolver) RedirectUrls(ctx context.Context, obj *model.OauthClient) ([]string, error) {
	clientID, err := gql.DecodeStrID(obj.ID)
	if err != nil {
		return nil, err
	}

	urls, err := r.oauthClientUseCase.GetOauthClientRedirectURLs(ctx, clientID)
	if err != nil {
		return nil, err
	}
	return cast.ArrayValues(urls, func(t *entities.OAuthClientRedirectURL) string {
		return t.RedirectURL
	}), nil
}

// OauthClient returns generated.OauthClientResolver implementation.
func (r *Resolver) OauthClient() generated.OauthClientResolver { return &oauthClientResolver{r} }

type oauthClientResolver struct{ *Resolver }
