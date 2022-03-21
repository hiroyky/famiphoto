package middlewares

import (
	"context"
	"github.com/hiroyky/famiphoto/config"
	"github.com/hiroyky/famiphoto/usecases"
	"github.com/hiroyky/famiphoto/utils"
	"net/http"
)

type AuthMiddleware interface {
	AuthClientCredential() func(handler http.Handler) http.Handler
}

func NewAuthMiddleware(oauthUseCase usecases.OauthUseCase) AuthMiddleware {
	return &authMiddleware{
		oauthUseCase: oauthUseCase,
	}
}

type authMiddleware struct {
	oauthUseCase usecases.OauthUseCase
}

func (m *authMiddleware) AuthClientCredential() func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
			ctx := req.Context()
			token, ok := utils.ParseAuthHeader(req.Header.Get("authorization"), "Bearer")
			if !ok {
				http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			sess, err := m.oauthUseCase.AuthAccessToken(ctx, token)
			if err != nil {
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			if sess == nil {
				next.ServeHTTP(writer, req)
				return
			}
			ctx = context.WithValue(ctx, config.ClientSessionKey, sess)
			req = req.WithContext(ctx)
			next.ServeHTTP(writer, req)
		})
	}
}
