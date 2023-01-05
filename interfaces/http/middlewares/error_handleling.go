package middlewares

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/interfaces/http/responses"
	"github.com/labstack/echo/v4"
	"github.com/vektah/gqlparser/v2/gqlerror"
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

func HandleGraphQLError(ctx context.Context, e error) *gqlerror.Error {
	err := graphql.DefaultErrorPresenter(ctx, e)
	if responses.IsFatalError(e) {
		fmt.Println(err.Error())
		err.Message = http.StatusText(http.StatusInternalServerError)
		err.Extensions = map[string]interface{}{
			"error":         http.StatusText(http.StatusInternalServerError),
			"error_code":    http.StatusText(http.StatusInternalServerError),
			"error_message": http.StatusText(http.StatusInternalServerError),
		}
	} else {
		err.Message = err.Error()
		err.Extensions = map[string]interface{}{
			"error":         http.StatusText(responses.GetStatusCode(err)),
			"error_code":    errors.GetFPErrorCode(err).ToString(),
			"error_message": err.Error(),
		}
	}
	return err
}
