package graph

import (
	"github.com/hiroyky/famiphoto/usecases"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userUseCase        usecases.UserUseCase
	searchUseCase      usecases.SearchUseCase
	groupUseCase       usecases.GroupUseCase
	oauthClientUseCase usecases.OauthUseCase
}

func NewResolver(
	userUseCase usecases.UserUseCase,
	searchUseCase usecases.SearchUseCase,
	groupUseCase usecases.GroupUseCase,
	oauthClientUseCase usecases.OauthUseCase,
) *Resolver {
	return &Resolver{
		userUseCase:        userUseCase,
		searchUseCase:      searchUseCase,
		groupUseCase:       groupUseCase,
		oauthClientUseCase: oauthClientUseCase,
	}
}
