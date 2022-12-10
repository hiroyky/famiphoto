package main

import (
	"github.com/hiroyky/famiphoto/di"
	"github.com/labstack/gommon/log"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "client-id",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "name",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "client-secret",
			Required: true,
		},
	}

	app.Action = func(ctx *cli.Context) error {
		uc := di.NewOAuthUseCase()
		err := uc.CreateSpecialOauthClient(
			ctx.Context,
			ctx.String("client-id"),
			ctx.String("name"),
			ctx.String("client-secret"),
		)
		if err != nil {
			log.Warn(err)
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Error(err)
		panic(err)
	}
}
