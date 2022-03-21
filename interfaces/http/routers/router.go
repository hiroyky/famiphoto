package routers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hiroyky/famiphoto/di"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/generated"
	"github.com/hiroyky/famiphoto/interfaces/http/middlewares"
	"github.com/hiroyky/famiphoto/interfaces/http/validators"
	"github.com/labstack/echo/v4"
	"net/http"
)

func New() *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = middlewares.HandlerError
	e.Validator = validators.NewValidator()

	authMiddleware := di.InitAuthMiddleware()

	e.GET("/status.html", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "ok")
	})

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: di.InitResolver()}))
	e.GET("/graphql", echo.WrapHandler(playground.Handler("GraphQL playground", "/graphql")))
	e.POST("/graphql", func(ctx echo.Context) error {
		srv.ServeHTTP(ctx.Response(), ctx.Request())
		return nil
	}, echo.WrapMiddleware(authMiddleware.AuthAccessToken()))

	oauthController := di.InitOauthController()
	e.POST("/oauth/v2/token", oauthController.PostToken, echo.WrapMiddleware(authMiddleware.AuthClientSecret()))

	return e
}
