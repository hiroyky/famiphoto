package controllers

import (
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/interfaces/http/requests"
	"github.com/hiroyky/famiphoto/usecases"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

type UploadController interface {
	UploadPhoto(ctx echo.Context) error
}

func NewUploadController(photoImportUseCase usecases.PhotoImportUseCase) UploadController {
	return &uploadController{photoImportUseCase: photoImportUseCase}
}

type uploadController struct {
	photoImportUseCase usecases.PhotoImportUseCase
}

func (c *uploadController) UploadPhoto(ctx echo.Context) error {
	var req requests.UploadPhotoRequest
	if err := req.Bind(ctx); err != nil {
		return err
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		return errors.New(errors.InvalidRequestError, err)
	}
	fp, err := file.Open()
	if err != nil {
		return errors.New(errors.InvalidRequestError, err)
	}
	defer fp.Close()

	data, err := io.ReadAll(fp)
	if err != nil {
		return errors.New(errors.InvalidRequestError, err)
	}

	if err := c.photoImportUseCase.UploadPhoto(ctx.Request().Context(), req.SignToken, file.Filename, data); err != nil {
		return err
	}

	return ctx.String(http.StatusCreated, "")
}
