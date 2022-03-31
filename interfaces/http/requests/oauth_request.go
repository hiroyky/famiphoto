package requests

import (
	"github.com/hiroyky/famiphoto/errors"
	"github.com/labstack/echo/v4"
)

type OauthGrantTokenRequest struct {
	GrantType    string `form:"grant_type" validate:"required,oneof=client_credentials authorization_code refresh_token token"`
	Scope        string `form:"scope" validate:"required"`
	Code         string `form:"code"`
	RedirectURL  string `form:"redirect_url"`
	RefreshToken string `form:"refresh_token"`
	State        string `form:"state"`
}

func (r *OauthGrantTokenRequest) Bind(ctx echo.Context) error {
	if err := ctx.Bind(r); err != nil {
		return errors.New(errors.InvalidRequestError, err)
	}
	return ctx.Validate(r)
}

type OauthAuthorizeGetRequest struct {
	ResponseType string `query:"response_type" validate:"required,oneof=code"`
	ClientID     string `query:"client_id" validate:"required"`
	RedirectURI  string `query:"redirect_uri" validate:"required,uri"`
	State        string `query:"state" validate:"required"`
	Scope        string `query:"scope"`
}

func (r *OauthAuthorizeGetRequest) Bind(ctx echo.Context) error {
	if err := ctx.Bind(r); err != nil {
		return errors.New(errors.InvalidRequestError, err)
	}
	return ctx.Validate(r)
}

type OauthAuthorizePostRequest struct {
	ClientID    string `form:"client_id" validate:"required"`
	RedirectURI string `form:"redirect_uri" validate:"required,uri"`
	State       string `form:"state" validate:"required"`
	Scope       string `form:"scope"`
	UserID      string `form:"user_id" validate:"required"`
	Password    string `form:"password" validate:"required"`
}

func (r *OauthAuthorizePostRequest) Bind(ctx echo.Context) error {
	if err := ctx.Bind(r); err != nil {
		return errors.New(errors.InvalidRequestError, err)
	}
	return ctx.Validate(r)
}
