package model

import (
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/entities"
	"net/url"
	"path"
)

func NewPhotoUploadInfo(sign *entities.PhotoUploadSign) *PhotoUploadInfo {
	u, err := url.Parse(config.Env.PhotoUploadBaseURL)
	if err != nil {
		panic(err)
	}
	u.Path = path.Join(u.Path, sign.SignToken)

	return &PhotoUploadInfo{
		UploadURL: u.String(),
		ExpireAt:  sign.ExpireAt,
	}
}
