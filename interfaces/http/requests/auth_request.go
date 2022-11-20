package requests

import (
	"github.com/hiroyky/famiphoto/errors"
	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	UserID   string `json:"user_id" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r *LoginRequest) Bind(ctx echo.Context) error {
	if err := ctx.Bind(r); err != nil {
		return errors.New(errors.InvalidRequestError, err)
	}
	return ctx.Validate(r)
}
