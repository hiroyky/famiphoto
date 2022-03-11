package graph

import (
	"github.com/hiroyky/famiphoto/di"
	"github.com/hiroyky/famiphoto/usecases"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userUseCase usecases.UserUseCase
}

func NewResolver() *Resolver {
	return &Resolver{userUseCase: di.InitUserUseCase()}
}
