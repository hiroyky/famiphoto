package middlewares

import (
	"context"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/interfaces/http/responses"
	"github.com/hiroyky/famiphoto/usecases"
	"github.com/hiroyky/famiphoto/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthMiddleware interface {
	AuthClientSecret(next echo.HandlerFunc) echo.HandlerFunc
	AuthAccessToken() func(handler http.Handler) http.Handler
}

func NewAuthMiddleware(oauthUseCase usecases.OauthUseCase) AuthMiddleware {
	return &authMiddleware{
		oauthUseCase: oauthUseCase,
	}
}

type authMiddleware struct {
	oauthUseCase usecases.OauthUseCase
}

func (m *authMiddleware) AuthClientSecret(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		clientID, clientSecret, ok := c.Request().BasicAuth()
		if !ok {
			return errors.New(errors.InvalidRequestError, nil)
		}
		oauthClient, err := m.oauthUseCase.AuthClientSecret(ctx, clientID, clientSecret)
		if err != nil {
			return responses.ConvertIfNotFatal(err, errors.OAuthClientUnauthorizedError)
		}
		c.Set(config.OauthClientKey, oauthClient)
		return next(c)
	}
}

func (m *authMiddleware) AuthAccessToken() func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
			ctx := req.Context()
			token, ok := utils.ParseAuthHeader(req.Header.Get("authorization"), "Bearer")
			if !ok {
				next.ServeHTTP(writer, req)
				return
			}

			sess, err := m.oauthUseCase.AuthAccessToken(ctx, token)
			if err != nil {
				code := responses.GetStatusCode(responses.ConvertIfNotFatal(err, errors.UnauthorizedError))
				http.Error(writer, http.StatusText(code), code)
				return
			}
			if sess == nil {
				http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			ctx = context.WithValue(ctx, config.ClientSessionKey, sess)
			req = req.WithContext(ctx)
			next.ServeHTTP(writer, req)
		})
	}
}
