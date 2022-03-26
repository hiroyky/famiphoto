package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CSRFByForm() echo.MiddlewareFunc {
	return middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
	})
}
