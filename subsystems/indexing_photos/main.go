package main

import (
	"fmt"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/di"
	"github.com/hiroyky/famiphoto/utils/array"
	"github.com/hiroyky/famiphoto/utils/log"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
	"time"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "env",
			Usage: "Path to env file",
			Value: "/etc/famiphoto/env",
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
		if err := godotenv.Load(ctx.String("env")); err != nil {
			log.Error("Failed to load env file, ", ctx.String("env"))
			return err
		}
		config.InitEnv()

		// インデックス中フラグを立ててプロセスを二重実行しないようにする
		pFilePath := "indexing"
		storage := di.NewTempStorageDriver()
		defer func() {
			storage.Delete(pFilePath)
		}()
		if storage.Exist(pFilePath) {
			log.Info("already running")
			return nil
		}
		if err := storage.CreateFile(pFilePath, []byte(fmt.Sprintf("%d", time.Now().Unix()))); err != nil {
			return err
		}

		extensions := array.Map(strings.Split(ctx.String("extensions"), ","), strings.TrimSpace)
		log.Info(fmt.Sprintf("extensions: %v, fast: %v \n", extensions, ctx.Bool("fast")))

		uc := di.NewPhotoImportUseCase()
		return uc.IndexingPhotos(ctx.Context, "", extensions, ctx.Bool("fast"))
	}
	if err := app.Run(os.Args); err != nil {
		log.Error("[FATAL]", err)
		panic(err)
	}
}
