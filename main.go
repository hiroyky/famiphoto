package main

import (
	"fmt"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/interfaces/http/routers"
	"github.com/hiroyky/famiphoto/utils/log"
)

func main() {
	config.InitEnv()
	app := routers.New()
	log.Info("Hello! Famiphoto API Server", config.Env.Port, config.Env.AppEnv)
	if err := app.Start(fmt.Sprintf(":%d", config.Env.Port)); err != nil {
		app.Logger.Fatal(err)
		log.Error(err)
	}
}
