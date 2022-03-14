//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/hiroyky/famiphoto/drivers/mysql"
	"github.com/hiroyky/famiphoto/infrastructures/repositories"
	"github.com/hiroyky/famiphoto/interfaces/http/controllers"
	"github.com/hiroyky/famiphoto/interfaces/http/graph"
	"github.com/hiroyky/famiphoto/services"
	"github.com/hiroyky/famiphoto/usecases"
)

func InitResolver() *graph.Resolver {
	wire.Build(
		graph.NewResolver,
		usecases.NewUserUseCase,
		repositories.NewUserRepository,
		services.NewPasswordService,
		mysql.NewDatabaseDriver,
		usecases.NewOauthUseCase,
		repositories.NewOauthClientRepository,
	)
	return nil
}

func InitOauthController() controllers.OauthController {
	wire.Build(
		controllers.NewOauthController,
		usecases.NewOauthUseCase,
		repositories.NewOauthClientRepository,
		services.NewPasswordService,
		mysql.NewDatabaseDriver,
	)
	return nil
}
