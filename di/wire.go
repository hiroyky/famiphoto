//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/hiroyky/famiphoto/drivers/mysql"
	"github.com/hiroyky/famiphoto/infrastructures/repositories"
	"github.com/hiroyky/famiphoto/services"
	"github.com/hiroyky/famiphoto/usecases"
)

func InitUserUseCase() usecases.UserUseCase {
	wire.Build(
		usecases.NewUserUseCase,
		repositories.NewUserRepository,
		services.NewPasswordService,
		mysql.NewDatabaseDriver,
	)
	return nil
}