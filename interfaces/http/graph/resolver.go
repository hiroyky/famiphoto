package graph

import (
	"github.com/hiroyky/famiphoto/usecases"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userUseCase        usecases.UserUseCase
	oauthClientUseCase usecases.OauthUseClientCase
}

func NewResolver(
	userUseCase usecases.UserUseCase,
	oauthClientUseCase usecases.OauthUseClientCase,
) *Resolver {
	return &Resolver{
		userUseCase:        userUseCase,
		oauthClientUseCase: oauthClientUseCase,
	}
}
