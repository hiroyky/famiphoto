package config

import "github.com/kelseyhightower/envconfig"

type FamiPhotoEnv struct {
	AppEnv                     string `envconfig:"APP_ENV"`
	Port                       int64  `envconfig:"PORT"`
	MySQLUserName              string `envconfig:"MYSQL_USER_NAME"`
	MySQLPassword              string `envconfig:"MYSQL_PASSWORD"`
	MySQLHostName              string `envconfig:"MYSQL_HOST_NAME"`
	MySQLPort                  string `envconfig:"MYSQL_PORT"`
	MySQLDatabase              string `envconfig:"MYSQL_DATABASE"`
	MediaSambaHostName         string `envconfig:"MEDIA_SAMBA_HOST_NAME"`
	MediaSambaUserName         string `envconfig:"MEDIA_SAMBA_USER_NAME"`
	MediaSambaPassword         string `envconfig:"MEDIA_SAMBA_PASSWORD"`
	MediaSambaShareName        string `envconfig:"MEDIA_SAMBA_SHARE_NAME"`
	OauthRedisHostName         string `envconfig:"OAUTH_REDIS_HOST_NAME"`
	OauthRedisDatabase         int64  `envconfig:"OAUTH_REDIS_DATABASE"`
	HMacKey                    string `envconfig:"HMAC_KEY"`
	CCAccessTokenExpireInSec   int64  `envconfig:"CC_ACCESS_TOKEN_EXPIRE_IN_SEC"`
	UserAccessTokenExpireInSec int64  `envconfig:"USER_ACCESS_TOKEN_EXPIRE_IN_SEC"`
	AccessTokenHashedPrefix    string `envconfig:"ACCESS_TOKEN_HASHED_PREFIX"`
}

var Env FamiPhotoEnv

func (e FamiPhotoEnv) IsDebug() bool {
	return e.AppEnv == Local
}

func init() {
	err := envconfig.Process("", &Env)
	if err != nil {
		panic(err)
	}
}

const (
	Local = "local"
	Prod  = "prod"
)
