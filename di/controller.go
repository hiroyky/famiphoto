package di

import "github.com/hiroyky/famiphoto/interfaces/http/controllers"

func NewOAuthController() controllers.OauthController {
	return controllers.NewOauthController(NewOAuthUseCase())
}

func NewDownloadController() controllers.DownloadController {
	return controllers.NewDownloadController(NewDownloadUseCase())
}
