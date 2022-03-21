package middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandlerError(err error, ctx echo.Context) {
	ctx.Logger().Error(err.Error())
	ctx.String(http.StatusInternalServerError, err.Error())
}
