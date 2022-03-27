package middlewares

import (
	"context"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/interfaces/http/responses"
	"github.com/hiroyky/famiphoto/usecases"
	"github.com/hiroyky/famiphoto/utils"
	"net/http"
)

type AuthMiddleware interface {
	AuthClientSecret() func(handler http.Handler) http.Handler
	AuthAccessToken() func(handler http.Handler) http.Handler
}

func NewAuthMiddleware(oauthUseCase usecases.OauthUseCase) AuthMiddleware {
	return &authMiddleware{
		oauthUseCase: oauthUseCase,
	}
}

type authMiddleware struct {
	oauthUseCase usecases.OauthUseCase
	userUseCase  usecases.UserUseCase
}

func (m *authMiddleware) AuthClientSecret() func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
			ctx := req.Context()
			clientID, clientSecret, ok := req.BasicAuth()
			if !ok {
				http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			oauthClient, err := m.oauthUseCase.AuthClientSecret(ctx, clientID, clientSecret)
			if err != nil {
				code := responses.GetStatusCode(err)
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

			sess, err := m.oauthUseCase.AuthAccessToken(ctx, token)
			if err != nil {
				code := responses.GetStatusCode(err)
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
