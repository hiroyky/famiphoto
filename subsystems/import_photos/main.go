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
	app.Name = "FAMIPHOTO import batch"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "base-dir",
			Usage:    "You have to specify base dir for import",
			Required: true,
		},
		&cli.StringFlag{
			Name:        "extensions",
			Usage:       "You can specify file extensions",
			Value:       ".jpeg, .jpg, .arw, .raw",
			DefaultText: ".jpeg, .jpg, .arw, .raw",
		},
	}

	app.Action = func(ctx *cli.Context) error {
		baseDir := ctx.String("base-dir")
		extensions := array.Map(strings.Split(ctx.String("extensions"), ","), strings.TrimSpace)
		fmt.Printf("base-dir: %s, extensions: %d %v\n", baseDir, len(extensions), extensions)
		uc := di.InitPhotoImportUseCase()
		if err := uc.ImportPhotos(ctx.Context, baseDir, extensions); err != nil {
			panic(err)
		}
		return nil
	}

	app.Run(os.Args)
}
