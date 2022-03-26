package responses

import (
	"github.com/hiroyky/famiphoto/errors"
	"net/http"
)

func GetStatusCode(err error) int {
	appCode := errors.GetFPErrorCode(err)
	switch appCode {
	case errors.UnauthorizedError,
		errors.OAuthClientUnauthorizedError:
		return http.StatusUnauthorized
	}
	return http.StatusInternalServerError
}

func IsFatalError(err error) bool {
	code := GetStatusCode(err)
	return code >= http.StatusInternalServerError
}
