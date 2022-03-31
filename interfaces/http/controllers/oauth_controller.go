package controllers

import (
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/interfaces/http/requests"
	"github.com/hiroyky/famiphoto/interfaces/http/responses"
	"github.com/hiroyky/famiphoto/usecases"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type OauthController interface {
	PostToken(ctx echo.Context) error
	GetAuthorize(ctx echo.Context) error
	PostAuthorize(ctx echo.Context) error
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
		credential, err := c.oauthUseCase.Oauth2ClientCredential(ctx.Request().Context(), client)
		if err != nil {
			return err
		}
		return ctx.JSON(http.StatusOK, responses.NewOauthAccessTokenFromClientCredential(credential))
	case "authorization_code":
		code, err := c.oauthUseCase.Oauth2AuthorizationCode(ctx.Request().Context(), client, req.Code, req.RedirectURL, time.Now())
		if err != nil {
			return err
		}
		return ctx.JSON(http.StatusOK, responses.NewOAuthAuthorizationCodeResponse(code))
	case "refresh_token":

	}

	return nil
}

func (c *oauthController) GetAuthorize(ctx echo.Context) error {
	var req requests.OauthAuthorizeGetRequest
	if err := req.Bind(ctx); err != nil {
		return err
	}

	client, err := c.oauthUseCase.ValidateToAuthorizeUser(ctx.Request().Context(), req.ClientID, req.RedirectURI, req.Scope)
	if err != nil {
		return err
	}

	csrf := ctx.Get("csrf").(string)
	args := responses.NewAuthorizePage(csrf, req.RedirectURI, req.State, req.Scope, client)
	return ctx.Render(http.StatusOK, "authorize.html", args)
}

func (c *oauthController) PostAuthorize(ctx echo.Context) error {
	var req requests.OauthAuthorizePostRequest
	if err := req.Bind(ctx); err != nil {
		return err
	}
	code, err := c.oauthUseCase.Authorize(ctx.Request().Context(), req.UserID, req.Password, req.ClientID, req.RedirectURI, req.Scope)
	if err != nil {
		return err
	}
	u, err := responses.NewOAuthCodeRedirectURL(req.RedirectURI, code, req.State)
	if err != nil {
		return err
	}

	return ctx.Redirect(http.StatusFound, u)
}
