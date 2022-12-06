package main

import (
	"fmt"
	"github.com/hiroyky/famiphoto/di"
	"github.com/hiroyky/famiphoto/utils/array"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "group-id",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "user-id",
			Required: true,
		},
		&cli.StringFlag{
			Name:        "extensions",
			Usage:       "You can specify file extensions",
			Value:       ".jpeg, .jpg, .arw, .raw",
			DefaultText: ".jpeg, .jpg, .arw, .raw",
		},
		&cli.BoolFlag{
			Name:     "fast",
			Usage:    "Specify true if you want fast mode, otherwise false",
			Required: true,
		},
	}

	app.Action = func(ctx *cli.Context) error {
		groupID := ctx.String("group-id")
		userID := ctx.String("user-id")
		extensions := array.Map(strings.Split(ctx.String("extensions"), ","), strings.TrimSpace)
		fmt.Printf("%s-%s, extensions: %v, fast: %v \n", groupID, userID, extensions, ctx.Bool("fast"))
		uc := di.NewPhotoImportUseCase()
		return uc.IndexingPhotos(ctx.Context, "", groupID, userID, extensions, ctx.Bool("fast"))
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
