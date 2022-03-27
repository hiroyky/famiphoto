//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/hiroyky/famiphoto/drivers/mysql"
	"github.com/hiroyky/famiphoto/drivers/redis"
	"github.com/hiroyky/famiphoto/infrastructures/repositories"
	"github.com/hiroyky/famiphoto/interfaces/http/controllers"
	"github.com/hiroyky/famiphoto/interfaces/http/graph"
	"github.com/hiroyky/famiphoto/interfaces/http/middlewares"
	"github.com/hiroyky/famiphoto/services"
	"github.com/hiroyky/famiphoto/usecases"
)

func InitResolver() *graph.Resolver {
	wire.Build(
		graph.NewResolver,
		usecases.NewUserUseCase,
		usecases.NewOauthUseCase,
		services.NewPasswordService,
		services.NewUserService,
		services.NewRandomService,
		repositories.NewUserRepository,
		repositories.NewOauthCodeAdapter,
		repositories.NewUserPasswordRepository,
		repositories.NewOauthClientRepository,
		repositories.NewOauthAccessTokenRepository,
		repositories.NewOauthClientRedirectURLRepository,
		mysql.NewDatabaseDriver,
		redis.NewOauthRedis,
	)
	return nil
}

func InitOauthController() controllers.OauthController {
	wire.Build(
		controllers.NewOauthController,
		usecases.NewOauthUseCase,
		services.NewPasswordService,
		services.NewRandomService,
		services.NewUserService,
		repositories.NewUserRepository,
		repositories.NewOauthClientRepository,
		repositories.NewOauthAccessTokenRepository,
		repositories.NewUserPasswordRepository,
		repositories.NewOauthClientRedirectURLRepository,
		repositories.NewOauthCodeAdapter,
		mysql.NewDatabaseDriver,
		redis.NewOauthRedis,
	)
	return nil
}

func InitAuthMiddleware() middlewares.AuthMiddleware {
	wire.Build(
		middlewares.NewAuthMiddleware,
		usecases.NewOauthUseCase,
		services.NewPasswordService,
		services.NewRandomService,
		services.NewUserService,
		repositories.NewUserRepository,
		repositories.NewOauthClientRepository,
		repositories.NewOauthAccessTokenRepository,
		repositories.NewUserPasswordRepository,
		repositories.NewOauthClientRedirectURLRepository,
		repositories.NewOauthCodeAdapter,
		mysql.NewDatabaseDriver,
		redis.NewOauthRedis,
	)
	return nil
}
