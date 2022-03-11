package routers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hiroyky/famiphoto/interfaces/http/graph"
	"github.com/hiroyky/famiphoto/interfaces/http/graph/generated"
	"github.com/labstack/echo/v4"
	"net/http"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/status.html", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "ok")
	})

	resolver := graph.NewResolver()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	e.GET("/graphql", echo.WrapHandler(playground.Handler("GraphQL playground", "/graphql/query")))
	e.POST("/graphql/query", echo.WrapHandler(srv))

	return e
}
