package responses

import (
	"github.com/hiroyky/famiphoto/errors"
	"net/http"
)

func GetStatusCode(err error) int {
	appCode := errors.GetFPErrorCode(err)
	switch appCode {
	case errors.InvalidRequestError,
		errors.PasswordWeakError:
		return http.StatusBadRequest
	case errors.UnauthorizedError,
		errors.UserUnauthorizedError,
		errors.OAuthClientUnauthorizedError:
		return http.StatusUnauthorized
	case errors.ForbiddenError:
		return http.StatusForbidden
	case errors.UserNotFoundError,
		errors.GroupNotFoundError,
		errors.UserPasswordNotFoundError,
		errors.FileNotFoundError:
		return http.StatusNotFound
	case errors.UserAlreadyExists:
		return http.StatusConflict
	}
	return http.StatusInternalServerError
}

func IsFatalError(err error) bool {
	code := GetStatusCode(err)
	return code >= http.StatusInternalServerError
}

func ConvertIfNotFatal(err error, dst errors.FamiPhotoErrorCode) error {
	if IsFatalError(err) {
		return err
	}
	return errors.New(dst, err)
}
