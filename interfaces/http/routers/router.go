package routers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/status.html", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "ok")
	})
	return e
}
