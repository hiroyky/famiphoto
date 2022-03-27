// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/hiroyky/famiphoto/drivers/mysql"
	"github.com/hiroyky/famiphoto/drivers/redis"
	"github.com/hiroyky/famiphoto/infrastructures/repositories"
	"github.com/hiroyky/famiphoto/interfaces/http/controllers"
	"github.com/hiroyky/famiphoto/interfaces/http/graph"
	"github.com/hiroyky/famiphoto/interfaces/http/middlewares"
	"github.com/hiroyky/famiphoto/services"
	"github.com/hiroyky/famiphoto/usecases"
)

// Injectors from wire.go:

func InitResolver() *graph.Resolver {
	sqlExecutor := mysql.NewDatabaseDriver()
	userAdapter := repositories.NewUserRepository(sqlExecutor)
	userPasswordAdapter := repositories.NewUserPasswordRepository(sqlExecutor)
	passwordService := services.NewPasswordService()
	userUseCase := usecases.NewUserUseCase(userAdapter, userPasswordAdapter, passwordService)
	oauthClientAdapter := repositories.NewOauthClientRepository(sqlExecutor)
	oauthClientRedirectURLAdapter := repositories.NewOauthClientRedirectURLRepository(sqlExecutor)
	redisAdapter := redis.NewOauthRedis()
	oauthAccessTokenAdapter := repositories.NewOauthAccessTokenRepository(redisAdapter)
	oauthCodeAdapter := repositories.NewOauthCodeAdapter(redisAdapter)
	userService := services.NewUserService(userAdapter, userPasswordAdapter, passwordService)
	randomService := services.NewRandomService()
	oauthUseCase := usecases.NewOauthUseCase(oauthClientAdapter, oauthClientRedirectURLAdapter, oauthAccessTokenAdapter, oauthCodeAdapter, userService, passwordService, randomService)
	resolver := graph.NewResolver(userUseCase, oauthUseCase)
	return resolver
}

func InitOauthController() controllers.OauthController {
	sqlExecutor := mysql.NewDatabaseDriver()
	oauthClientAdapter := repositories.NewOauthClientRepository(sqlExecutor)
	oauthClientRedirectURLAdapter := repositories.NewOauthClientRedirectURLRepository(sqlExecutor)
	redisAdapter := redis.NewOauthRedis()
	oauthAccessTokenAdapter := repositories.NewOauthAccessTokenRepository(redisAdapter)
	oauthCodeAdapter := repositories.NewOauthCodeAdapter(redisAdapter)
	userAdapter := repositories.NewUserRepository(sqlExecutor)
	userPasswordAdapter := repositories.NewUserPasswordRepository(sqlExecutor)
	passwordService := services.NewPasswordService()
	userService := services.NewUserService(userAdapter, userPasswordAdapter, passwordService)
	randomService := services.NewRandomService()
	oauthUseCase := usecases.NewOauthUseCase(oauthClientAdapter, oauthClientRedirectURLAdapter, oauthAccessTokenAdapter, oauthCodeAdapter, userService, passwordService, randomService)
	oauthController := controllers.NewOauthController(oauthUseCase)
	return oauthController
}

func InitAuthMiddleware() middlewares.AuthMiddleware {
	sqlExecutor := mysql.NewDatabaseDriver()
	oauthClientAdapter := repositories.NewOauthClientRepository(sqlExecutor)
	oauthClientRedirectURLAdapter := repositories.NewOauthClientRedirectURLRepository(sqlExecutor)
	redisAdapter := redis.NewOauthRedis()
	oauthAccessTokenAdapter := repositories.NewOauthAccessTokenRepository(redisAdapter)
	oauthCodeAdapter := repositories.NewOauthCodeAdapter(redisAdapter)
	userAdapter := repositories.NewUserRepository(sqlExecutor)
	userPasswordAdapter := repositories.NewUserPasswordRepository(sqlExecutor)
	passwordService := services.NewPasswordService()
	userService := services.NewUserService(userAdapter, userPasswordAdapter, passwordService)
	randomService := services.NewRandomService()
	oauthUseCase := usecases.NewOauthUseCase(oauthClientAdapter, oauthClientRedirectURLAdapter, oauthAccessTokenAdapter, oauthCodeAdapter, userService, passwordService, randomService)
	authMiddleware := middlewares.NewAuthMiddleware(oauthUseCase)
	return authMiddleware
}
