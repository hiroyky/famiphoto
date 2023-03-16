package main

import (
	"fmt"
	"github.com/hiroyky/famiphoto/utils"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "in",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "out",
			Required: false,
		},
	}

	app.Action = func(ctx *cli.Context) error {
		data, err := os.ReadFile(ctx.String("in"))
		if err != nil {
			return err
		}

		thumbnail, err := utils.ExtractThumbnail(data)
		if err != nil {
			return err
		}

		fmt.Println(len(thumbnail))

		return os.WriteFile(ctx.String("out"), thumbnail, 0644)
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
