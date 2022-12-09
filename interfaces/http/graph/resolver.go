package graph

import (
	"github.com/hiroyky/famiphoto/usecases"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userUseCase        usecases.UserUseCase
	photoUseCase       usecases.PhotoUseCase
	searchUseCase      usecases.SearchUseCase
	groupUseCase       usecases.GroupUseCase
	oauthClientUseCase usecases.OauthUseCase
	photoImportUseCase usecases.PhotoImportUseCase
}

func NewResolver(
	userUseCase usecases.UserUseCase,
	photoUseCase usecases.PhotoUseCase,
	searchUseCase usecases.SearchUseCase,
	groupUseCase usecases.GroupUseCase,
	oauthClientUseCase usecases.OauthUseCase,
	photoImportUseCase usecases.PhotoImportUseCase,
) *Resolver {
	return &Resolver{
		userUseCase:        userUseCase,
		photoUseCase:       photoUseCase,
		searchUseCase:      searchUseCase,
		groupUseCase:       groupUseCase,
		oauthClientUseCase: oauthClientUseCase,
		photoImportUseCase: photoImportUseCase,
	}
}
