package di

import "github.com/hiroyky/famiphoto/interfaces/http/middlewares"

func NewAuthMiddleware() middlewares.AuthMiddleware {
	return middlewares.NewAuthMiddleware(NewOAuthUseCase())
}
