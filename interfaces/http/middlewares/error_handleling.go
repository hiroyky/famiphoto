package middlewares

import (
	"fmt"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/interfaces/http/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandlerError(err error, ctx echo.Context) {
	if responses.IsFatalError(err) {
		ctx.Logger().Error(err.Error())
		fmt.Println(err.Error())
		ctx.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error":         http.StatusText(http.StatusInternalServerError),
				"error_code":    http.StatusText(http.StatusInternalServerError),
				"error_message": http.StatusText(http.StatusInternalServerError),
			},
		)
	}
	ctx.JSON(
		responses.GetStatusCode(err),
		map[string]string{
			"error":         http.StatusText(responses.GetStatusCode(err)),
			"error_code":    errors.GetFPErrorCode(err).ToString(),
			"error_message": err.Error(),
		},
	)
}
