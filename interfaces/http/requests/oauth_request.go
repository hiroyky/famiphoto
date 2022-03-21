package requests

import (
	"github.com/hiroyky/famiphoto/errors"
	"github.com/labstack/echo/v4"
)

type OauthGrantTokenRequest struct {
	GrantType    string `form:"grant_type" validators:"required,oneof=client_credentials authorization_code refresh_token"`
	Scope        string `form:"scope" validators:"required"`
	Code         string `form:"code"`
	RedirectURL  string `form:"redirect_url"`
	RefreshToken string `form:"refresh_token"`
}

func (r *OauthGrantTokenRequest) Bind(ctx echo.Context) error {
	if err := ctx.Bind(r); err != nil {
		return errors.New(errors.InvalidRequestError, err)
	}
	return ctx.Validate(r)
}
