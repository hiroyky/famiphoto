package requests

import "github.com/labstack/echo/v4"

type FileDownloadRequest struct {
	FileID string `param:"file_id" validate:"required"`
}

func (r *FileDownloadRequest) Bind(ctx echo.Context) error {
	if err := ctx.Bind(r); err != nil {
		return err
	}
	return ctx.Validate(r)
}
