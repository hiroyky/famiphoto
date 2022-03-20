package config

import "github.com/kelseyhightower/envconfig"

type FamiPhotoEnv struct {
	Port                     int64  `envconfig:"PORT"`
	MySQLUserName            string `envconfig:"MYSQL_USER_NAME"`
	MySQLPassword            string `envconfig:"MYSQL_PASSWORD"`
	MySQLHostName            string `envconfig:"MYSQL_HOST_NAME"`
	MySQLPort                string `envconfig:"MYSQL_PORT"`
	MySQLDatabase            string `envconfig:"MYSQL_DATABASE"`
	OauthRedisHostName       string `envconfig:"OAUTH_REDIS_HOST_NAME"`
	OauthRedisDatabase       int64  `envconfig:"OAUTH_REDIS_DATABASE"`
	HMacKey                  string `envconfig:"HMAC_KEY"`
	CCAccessTokenExpireInSec int64  `envconfig:"CC_ACCESS_TOKEN_EXPIRE_IN_SEC"`
	AccessTokenHashedPrefix  string `envconfig:"ACCESS_TOKEN_HASHED_PREFIX"`
}

var Env FamiPhotoEnv

func init() {
	err := envconfig.Process("", &Env)
	if err != nil {
		panic(err)
	}
}
