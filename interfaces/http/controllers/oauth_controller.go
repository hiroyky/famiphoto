package controllers

import (
	"github.com/hiroyky/famiphoto/usecases"
	"github.com/labstack/echo/v4"
)

type OauthController interface {
	PostToken(ctx echo.Context) error
}

func NewOauthController(oauthUseCase usecases.OauthUseCase) OauthController {
	return &oauthController{
		oauthUseCase: oauthUseCase,
	}
}

type oauthController struct {
	oauthUseCase usecases.OauthUseCase
}

func (c *oauthController) PostToken(ctx echo.Context) error {
	return nil
}
