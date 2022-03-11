package config

import "github.com/kelseyhightower/envconfig"

type FamiPhotoEnv struct {
	Port          int64  `envconfig:"PORT"`
	MySQLUserName string `envconfig:"MYSQL_USER_NAME"`
	MySQLPassword string `envconfig:"MYSQL_PASSWORD"`
	MySQLHostName string `envconfig:"MYSQL_HOST_NAME"`
	MySQLPort     string `envconfig:"MYSQL_PORT"`
	MySQLDatabase string `envconfig:"MYSQL_DATABASE"`
	HMacKey       string `envconfig:"HMAC_KEY"`
}

var Env FamiPhotoEnv

func init() {
	err := envconfig.Process("", &Env)
	if err != nil {
		panic(err)
	}
}
