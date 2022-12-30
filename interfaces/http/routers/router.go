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
	e.Renderer = responses.NewHtmlTemplateRenderer()
	authMiddleware := di.NewAuthMiddleware()
	authController := di.NewAuthController()
	downloadController := di.NewDownloadController()
	oauthController := di.NewOAuthController()

	e.GET("/status.html", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "ok")
	})

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: di.NewResolver()}))
	srv.SetErrorPresenter(middlewares.HandleGraphQLError)
	e.POST(
		"/graphql",
		func(ctx echo.Context) error {
			srv.ServeHTTP(ctx.Response(), ctx.Request())
			return nil
		},
		echo.WrapMiddleware(authMiddleware.AuthClientSecret()),
		echo.WrapMiddleware(authMiddleware.AuthAccessToken()),
		echo.WrapMiddleware(authMiddleware.VerifyClient()),
	)

	if config.Env.IsDebug() {
		e.GET("/debug/graphql", echo.WrapHandler(playground.Handler("GraphQL playground", "/graphql")))
	}

	e.POST("/auth/login", authController.Login, authMiddleware.MustAuthClientSecret, authMiddleware.MustVerifyAdminClient)

	e.POST("/oauth/v2/token", oauthController.PostToken, authMiddleware.MustAuthClientSecret)

	e.Group("assets").Use(middleware.StaticWithConfig(middleware.StaticConfig{Root: "assets"}))
	download := e.Group("/download",
		echo.WrapMiddleware(authMiddleware.AuthClientSecret()),
		echo.WrapMiddleware(authMiddleware.AuthAccessToken()),
		echo.WrapMiddleware(authMiddleware.VerifyClient()),
	)
	download.GET("/files/:file_id", downloadController.GetFileDownload, authMiddleware.VerifyFileDownloadPermission)
	return e
}
