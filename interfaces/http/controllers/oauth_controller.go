package controllers

import (
	"context"
	"github.com/hiroyky/famiphoto/interfaces/http/requests"
	"github.com/hiroyky/famiphoto/interfaces/http/responses"
	"github.com/hiroyky/famiphoto/usecases"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
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
	var req requests.OauthGrantTokenRequest
	if err := req.Bind(ctx); err != nil {
		return err
	}

	switch req.GrantType {
	case "client_credentials":
		credential, err := c.oauthUseCase.Oauth2ClientCredential(context.Background(), req.ClientID, req.ClientSecret, time.Now())
		if err != nil {
			return err
		}
		return ctx.JSON(http.StatusOK, responses.NewOauthAccessTokenFromClientCredential(credential))
	}

	return nil
}
