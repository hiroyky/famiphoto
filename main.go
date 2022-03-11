package main

import (
	"fmt"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/interfaces/http/routers"
)

func main() {
	app := routers.New()
	if err := app.Start(fmt.Sprintf(":%d", config.Env.Port)); err != nil {
		app.Logger.Fatal(err)
	}
}
