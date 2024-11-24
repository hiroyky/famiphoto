package controllers

import (
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/interfaces/http/requests"
	"github.com/hiroyky/famiphoto/interfaces/http/responses"
	"github.com/hiroyky/famiphoto/usecases"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func NewAuthController(userUseCase usecases.UserUseCase) AuthController {
	return &authController{
		userUseCase: userUseCase,
	}
}

type AuthController interface {
	Login(ctx echo.Context) error
}

type authController struct {
	userUseCase usecases.UserUseCase
}

func (c *authController) Login(ctx echo.Context) error {
	client, ok := ctx.Get(config.OauthClientKey).(*entities.OauthClient)
	if !ok {
		return errors.New(errors.ContextValueNotFoundFatal, nil)
	}

	var req requests.LoginRequest
	if err := req.Bind(ctx); err != nil {
		return err
	}

	code, err := c.userUseCase.Login(ctx.Request().Context(), client, req.UserID, req.Password, time.Now())
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, responses.NewOAuthAuthorizationCodeResponse(code))
}
