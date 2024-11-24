package middlewares

import (
	"context"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/entities"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/hiroyky/famiphoto/interfaces/http/responses"
	"github.com/hiroyky/famiphoto/usecases"
	"github.com/hiroyky/famiphoto/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthMiddleware interface {
	MustAuthClientSecret(next echo.HandlerFunc) echo.HandlerFunc
	MustVerifyAdminClient(next echo.HandlerFunc) echo.HandlerFunc
	AuthClientSecret() func(handler http.Handler) http.Handler
	AuthAccessToken() func(handler http.Handler) http.Handler
	VerifyClient() func(handler http.Handler) http.Handler
	VerifyFileDownloadPermission(next echo.HandlerFunc) echo.HandlerFunc
}

func NewAuthMiddleware(oauthUseCase usecases.OauthUseCase, downloadUseCase usecases.DownloadUseCase) AuthMiddleware {
	return &authMiddleware{
		oauthUseCase:    oauthUseCase,
		downloadUseCase: downloadUseCase,
	}
}

type authMiddleware struct {
	oauthUseCase    usecases.OauthUseCase
	downloadUseCase usecases.DownloadUseCase
}

func (m *authMiddleware) MustAuthClientSecret(next echo.HandlerFunc) echo.HandlerFunc {
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

func (m *authMiddleware) MustVerifyAdminClient(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		client, ok := c.Get(config.OauthClientKey).(*entities.OauthClient)
		if !ok {
			return errors.New(errors.ContextValueNotFoundFatal, nil)
		}
		if client.Scope != entities.OauthScopeAdmin {
			return errors.New(errors.ForbiddenError, nil)
		}
		return next(c)
	}
}

func (m *authMiddleware) AuthClientSecret() func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
			ctx := req.Context()
			clientID, clientSecret, ok := req.BasicAuth()
			if !ok {
				next.ServeHTTP(writer, req)
				return
			}
			oauthClient, err := m.oauthUseCase.AuthClientSecret(ctx, clientID, clientSecret)
			if err != nil {
				code := responses.GetStatusCode(responses.ConvertIfNotFatal(err, errors.UnauthorizedError))
				http.Error(writer, http.StatusText(code), code)
				return
			}
			ctx = context.WithValue(ctx, config.OauthClientKey, oauthClient)
			req = req.WithContext(ctx)
			next.ServeHTTP(writer, req)
		})
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

			sess, client, err := m.oauthUseCase.AuthAccessToken(ctx, token)
			if err != nil {
				next.ServeHTTP(writer, req)
				return
			}
			if sess == nil {
				next.ServeHTTP(writer, req)
				return
			}
			ctx = context.WithValue(ctx, config.ClientSessionKey, sess)
			ctx = context.WithValue(ctx, config.OauthClientKey, client)
			req = req.WithContext(ctx)
			next.ServeHTTP(writer, req)
		})
	}
}

// VerifyClient クライアント情報が確認できなければエラーとする、APIにアクセスさせない。
func (m *authMiddleware) VerifyClient() func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
			ctx := req.Context()
			_, ok := ctx.Value(config.OauthClientKey).(*entities.OauthClient)
			if !ok {
				http.Error(writer, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}
			next.ServeHTTP(writer, req)
		})
	}
}

func (m *authMiddleware) VerifyFileDownloadPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		_, ok := ctx.Value(config.ClientSessionKey).(*entities.OauthSession)
		if !ok {
			return errors.New(errors.UserUnauthorizedError, nil)
		}

		return next(c)
	}
}
