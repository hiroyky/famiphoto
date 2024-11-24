package responses

import (
	"github.com/hiroyky/famiphoto/errors"
	"net/http"
)

func GetStatusCode(err error) int {
	appCode := errors.GetFPErrorCode(err)
	switch appCode {
	case errors.InvalidRequestError,
		errors.OAuthClientInvalidRedirectURLError,
		errors.PasswordWeakError:
		return http.StatusBadRequest
	case errors.UnauthorizedError,
		errors.UserUnauthorizedError,
		errors.OAuthAccessTokenNotFoundError,
		errors.UserAuthNotFoundError,
		errors.OAuthClientUnauthorizedError:
		return http.StatusUnauthorized
	case errors.ForbiddenError:
		return http.StatusForbidden
	case errors.UserNotFoundError,
		errors.DBRowNotFoundError,
		errors.GroupNotFoundError,
		errors.UserPasswordNotFoundError,
		errors.OAuthClientNotFoundError,
		errors.PhotoUploadSignNotFoundError,
		errors.RedisKeyNotFoundError,
		errors.PhotoNotFoundError,
		errors.FileNotFoundError:
		return http.StatusNotFound
	case errors.UserAlreadyExists,
		errors.OAuthClientAlreadyExist,
		errors.GroupAlreadyExistError:
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
