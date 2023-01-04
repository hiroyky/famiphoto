package requests

import "github.com/labstack/echo/v4"

type UploadPhotoRequest struct {
	SignToken string `param:"sign_token" validate:"required"`
}

func (r *UploadPhotoRequest) Bind(ctx echo.Context) error {
	if err := ctx.Bind(r); err != nil {
		return err
	}
	return ctx.Validate(r)
}
