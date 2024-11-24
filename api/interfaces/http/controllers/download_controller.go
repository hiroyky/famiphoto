package controllers

import (
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/interfaces/http/requests"
	"github.com/hiroyky/famiphoto/usecases"
	"github.com/hiroyky/famiphoto/utils/gql"
	"github.com/labstack/echo/v4"
	"net/http"
)

type DownloadController interface {
	GetFileDownload(ctx echo.Context) error
}

func NewDownloadController(downloadUseCase usecases.DownloadUseCase) DownloadController {
	return &downloadController{
		downloadUseCase: downloadUseCase,
	}
}

type downloadController struct {
	downloadUseCase usecases.DownloadUseCase
}

func (c *downloadController) GetFileDownload(ctx echo.Context) error {
	var req requests.FileDownloadRequest
	if err := req.Bind(ctx); err != nil {
		return err
	}

	fileID, err := gql.DecodeIntID(req.FileID)
	if err != nil {
		return errors.New(errors.FileNotFoundError, err)
	}

	data, file, err := c.downloadUseCase.LoadPhotoFile(ctx.Request().Context(), fileID)
	if err != nil {
		return err
	}

	return ctx.Blob(http.StatusOK, file.MimeType(), data)
}
