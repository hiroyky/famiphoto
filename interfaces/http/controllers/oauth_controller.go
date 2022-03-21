package controllers

import (
	"context"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/interfaces/http/requests"
	"github.com/hiroyky/famiphoto/interfaces/http/responses"
	"github.com/hiroyky/famiphoto/usecases"
	"github.com/labstack/echo/v4"
	"net/http"
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
	client, ok := ctx.Request().Context().Value(config.OauthClientKey).(*entities.OauthClient)
	if !ok {
		return ctx.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	var req requests.OauthGrantTokenRequest
	if err := req.Bind(ctx); err != nil {
		return err
	}

	switch req.GrantType {
	case "client_credentials":
		credential, err := c.oauthUseCase.Oauth2ClientCredential(context.Background(), client)
		if err != nil {
			return err
		}
		return ctx.JSON(http.StatusOK, responses.NewOauthAccessTokenFromClientCredential(credential))
	}

	return nil
}
