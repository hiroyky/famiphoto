package routers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/di"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/generated"
	"github.com/hiroyky/famiphoto/interfaces/http/middlewares"
	"github.com/hiroyky/famiphoto/interfaces/http/validators"
	"github.com/hiroyky/famiphoto/utils/log"
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
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		Skipper:        nil,
		BeforeNextFunc: nil,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				log.Error(v.Method, v.URIPath, v.Error)
			}
			return nil
		},
		LogLatency:   true,
		LogProtocol:  true,
		LogRemoteIP:  false,
		LogHost:      false,
		LogMethod:    true,
		LogURI:       true,
		LogURIPath:   true,
		LogRoutePath: true,
		LogStatus:    true,
		LogError:     true,
	}))
	e.Use(middleware.Recover())
	authMiddleware := di.NewAuthMiddleware()
	authController := di.NewAuthController()
	uploadController := di.NewUploadController()
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
		e.GET("/debug/graphql", echo.WrapHandler(playground.Handler("GraphQL playground", "/debug/graphql")))
		e.POST(
			"/debug/graphql",
			func(ctx echo.Context) error {
				srv.ServeHTTP(ctx.Response(), ctx.Request())
				return nil
			},
			echo.WrapMiddleware(authMiddleware.AuthClientSecret()),
			echo.WrapMiddleware(authMiddleware.AuthAccessToken()),
			echo.WrapMiddleware(authMiddleware.VerifyClient()),
		)
	}

	e.POST("/auth/login", authController.Login, authMiddleware.MustAuthClientSecret, authMiddleware.MustVerifyAdminClient)

	e.POST("/oauth/v2/token", oauthController.PostToken, authMiddleware.MustAuthClientSecret)
	e.POST("/upload_photo/:sign_token", uploadController.UploadPhoto)
	e.Group("assets").Use(middleware.StaticWithConfig(middleware.StaticConfig{Root: config.Env.AssetRootPath}))
	download := e.Group("/download",
		echo.WrapMiddleware(authMiddleware.AuthClientSecret()),
		echo.WrapMiddleware(authMiddleware.AuthAccessToken()),
		echo.WrapMiddleware(authMiddleware.VerifyClient()),
	)
	download.GET("/files/:file_id", downloadController.GetFileDownload, authMiddleware.VerifyFileDownloadPermission)
	return e
}
