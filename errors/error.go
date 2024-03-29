package errors

import (
	native "errors"
	"fmt"
)

type FamiPhotoError struct {
	errorCode FamiPhotoErrorCode
	baseError error
}

func (e *FamiPhotoError) Error() string {
	if e.baseError == nil {
		return e.ErrorCode().ToString()
	}
	return fmt.Sprintf("code: %s, %s", e.errorCode, e.baseError.Error())
}

func (e *FamiPhotoError) ErrorCode() FamiPhotoErrorCode {
	return e.errorCode
}

type FamiPhotoErrorCode string

func (c FamiPhotoErrorCode) ToString() string {
	return string(c)
}

const (
	Unknown                            FamiPhotoErrorCode = "Unknown"
	InvalidRequestError                FamiPhotoErrorCode = "InvalidRequestError"
	UserNotFoundError                  FamiPhotoErrorCode = "UserNotFoundError"
	GroupNotFoundError                 FamiPhotoErrorCode = "GroupNotFoundError"
	GroupAlreadyExistError             FamiPhotoErrorCode = "GroupAlreadyExistError"
	UserCreateFatal                    FamiPhotoErrorCode = "UserCreateFatal"
	UserAlreadyExists                  FamiPhotoErrorCode = "UserAlreadyExists"
	UserUnauthorizedError              FamiPhotoErrorCode = "UserUnauthorizedError"
	PasswordWeakError                  FamiPhotoErrorCode = "PasswordWeakError"
	UserPasswordNotFoundError          FamiPhotoErrorCode = "UserPasswordNotFoundError"
	OAuthClientNotFoundError           FamiPhotoErrorCode = "OAuthClientNotFoundError"
	OAuthClientAlreadyExist            FamiPhotoErrorCode = "OAuthClientAlreadyExist"
	OAuthClientCreateFatal             FamiPhotoErrorCode = "OAuthClientCreateFatal"
	OAuthClientUnauthorizedError       FamiPhotoErrorCode = "OAuthClientUnauthorizedError"
	OAuthClientInvalidRedirectURLError FamiPhotoErrorCode = "OAuthClientInvalidRedirectURLError"
	OAuthAccessTokenNotFoundError      FamiPhotoErrorCode = "OAuthAccessTokenNotFoundError"
	PhotoUploadSignNotFoundError       FamiPhotoErrorCode = "PhotoUploadSignNotFoundError"
	UserAuthNotFoundError              FamiPhotoErrorCode = "UserAuthNotFoundError"
	UnauthorizedError                  FamiPhotoErrorCode = "UnauthorizedError"
	DBRowNotFoundError                 FamiPhotoErrorCode = "DBRowNotFoundError"
	InvalidFilePathFatal               FamiPhotoErrorCode = "InvalidFilePathFatal"
	ForbiddenError                     FamiPhotoErrorCode = "ForbiddenError"
	TxnRollbackFatal                   FamiPhotoErrorCode = "TxnRollbackFatal"
	TxnBeginFatal                      FamiPhotoErrorCode = "TxnBeginFatal"
	HashPasswordFatal                  FamiPhotoErrorCode = "HashPasswordFatal"
	MatchPasswordFatal                 FamiPhotoErrorCode = "MatchPasswordFatal"
	RedisKeyNotFoundError              FamiPhotoErrorCode = "RedisKeyNotFoundError"
	RedisFatal                         FamiPhotoErrorCode = "RedisFatal"
	ContextValueNotFoundFatal          FamiPhotoErrorCode = "ContextValueNotFoundFatal"
	UnExpectedFileAlreadyExistError    FamiPhotoErrorCode = "UnExpectedFileAlreadyExistError"
	FileNotFoundError                  FamiPhotoErrorCode = "FileNotFoundError"
	PhotoNotFoundError                 FamiPhotoErrorCode = "PhotoNotFoundError"
	FileAlreadyExistError              FamiPhotoErrorCode = "FileAlreadyExistError"
	ElasticSearchFatal                 FamiPhotoErrorCode = "ElasticSearchFatal"
	InvalidTimezoneFatal               FamiPhotoErrorCode = "InvalidTimezoneFatal"
	NoExifError                        FamiPhotoErrorCode = "NoExifError"
)

func New(errCode FamiPhotoErrorCode, baseError error) error {
	return &FamiPhotoError{
		errorCode: errCode,
		baseError: baseError,
	}
}

func UnwrapFPError(err error) *FamiPhotoError {
	var dst *FamiPhotoError
	if ok := native.As(err, &dst); ok {
		return dst
	}
	return nil
}

func GetFPErrorCode(err error) FamiPhotoErrorCode {
	appError := UnwrapFPError(err)
	if appError == nil {
		return Unknown
	}
	return appError.ErrorCode()
}

func Is(err, target error) bool {
	return native.Is(err, target)
}

func IsErrCode(err error, errCode FamiPhotoErrorCode) bool {
	if err == nil {
		return false
	}
	return GetFPErrorCode(err) == errCode
}
