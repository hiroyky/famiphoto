//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/hiroyky/famiphoto/drivers/mysql"
	"github.com/hiroyky/famiphoto/drivers/redis"
	"github.com/hiroyky/famiphoto/drivers/samba"
	"github.com/hiroyky/famiphoto/infrastructures/repositories"
	"github.com/hiroyky/famiphoto/interfaces/http/controllers"
	"github.com/hiroyky/famiphoto/interfaces/http/graph"
	"github.com/hiroyky/famiphoto/interfaces/http/middlewares"
	"github.com/hiroyky/famiphoto/services"
	"github.com/hiroyky/famiphoto/usecases"
)

func InitPhotoImportUseCase() usecases.PhotoImportUseCase {
	wire.Build(
		usecases.NewPhotoImportUseCase,
		services.NewPhotoService,
		repositories.NewPhotoStorageRepository,
		samba.NewMediaSambaStorage,
	)
	return nil
}

func InitResolver() *graph.Resolver {
	wire.Build(
		graph.NewResolver,
		usecases.NewUserUseCase,
		initOauthUseCase,
		services.NewPasswordService,
		repositories.NewUserRepository,
		repositories.NewUserPasswordRepository,
		mysql.NewDatabaseDriver,
	)
	return nil
}

func InitOauthController() controllers.OauthController {
	wire.Build(
		controllers.NewOauthController,
		initOauthUseCase,
	)
	return nil
}

func initOauthUseCase() usecases.OauthUseCase {
	wire.Build(
		usecases.NewOauthUseCase,
		services.NewAuthService,
		services.NewPasswordService,
		services.NewRandomService,
		services.NewUserService,
		repositories.NewUserRepository,
		repositories.NewUserAuthAdapter,
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
		initOauthUseCase,
	)
	return nil
}
