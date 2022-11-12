package routers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/di"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/generated"
	"github.com/hiroyky/famiphoto/interfaces/http/middlewares"
	"github.com/hiroyky/famiphoto/interfaces/http/responses"
	"github.com/hiroyky/famiphoto/interfaces/http/validators"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echotrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/labstack/echo.v4"
	"net/http"
)

func New() *echo.Echo {
	e := echo.New()
	e.HTTPErrorHandler = middlewares.HandlerError
	e.Validator = validators.NewValidator()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(echotrace.Middleware())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	authMiddleware := di.NewAuthMiddleware()

	e.GET("/status.html", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "ok")
	})

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: di.NewResolver()}))
	e.POST("/graphql", func(ctx echo.Context) error {
		srv.ServeHTTP(ctx.Response(), ctx.Request())
		return nil
	}, echo.WrapMiddleware(authMiddleware.AuthAccessToken()))

	if config.Env.IsDebug() {
		e.GET("/debug/graphql", echo.WrapHandler(playground.Handler("GraphQL playground", "/graphql")))
	}

	e.Renderer = responses.NewHtmlTemplateRenderer()

	oauthController := di.NewOAuthController()
	e.POST("/oauth/v2/token", oauthController.PostToken, authMiddleware.AuthClientSecret)
	e.GET("/oauth/authorize", oauthController.GetAuthorize, middlewares.CSRFByForm())
	e.POST("/oauth/authorize", oauthController.PostAuthorize, middlewares.CSRFByForm())

	e.Group("assets").Use(middleware.StaticWithConfig(middleware.StaticConfig{Root: "assets"}))

	downloadController := di.NewDownloadController()
	e.GET("/download/files/:file_id", downloadController.GetFileDownload)
	return e
}
