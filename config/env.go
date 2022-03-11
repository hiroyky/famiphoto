package config

import "github.com/kelseyhightower/envconfig"

type FamiPhotoEnv struct {
	Port int64 `envconfig:"PORT"`
}

var Env FamiPhotoEnv

func init() {
	err := envconfig.Process("", &Env)
	if err != nil {
		panic(err)
	}
}
