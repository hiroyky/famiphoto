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
	UserAuthNotFoundError              FamiPhotoErrorCode = "UserAuthNotFoundError"
	UnauthorizedError                  FamiPhotoErrorCode = "UnauthorizedError"
	DBColumnNotFoundError              FamiPhotoErrorCode = "DBColumnNotFoundError"
	InvalidFilePathFatal               FamiPhotoErrorCode = "InvalidFilePathFatal"
	ForbiddenError                     FamiPhotoErrorCode = "ForbiddenError"
	TxnRollbackFatal                   FamiPhotoErrorCode = "TxnRollbackFatal"
	TxnBeginFatal                      FamiPhotoErrorCode = "TxnBeginFatal"
	HashPasswordFatal                  FamiPhotoErrorCode = "HashPasswordFatal"
	MatchPasswordFatal                 FamiPhotoErrorCode = "MatchPasswordFatal"
	RedisKeyNotFoundError              FamiPhotoErrorCode = "RedisKeyNotFoundError"
	RedisFatal                         FamiPhotoErrorCode = "RedisFatal"
	ContextValueNotFoundFatal          FamiPhotoErrorCode = "ContextValueNotFoundFatal"
	FileNotFoundError                  FamiPhotoErrorCode = "FileNotFoundError"
	SambaConnectFatal                  FamiPhotoErrorCode = "SambaConnectFatal"
	SambaCreateFatal                   FamiPhotoErrorCode = "SambaCreateFatal"
	SambaReadFatal                     FamiPhotoErrorCode = "SambaReadFatal"
	SambaCreateDirFatal                FamiPhotoErrorCode = "SambaCreateDirFatal"
	SambaRenameFatal                   FamiPhotoErrorCode = "SambaRenameFatal"
	SambaDeleteFatal                   FamiPhotoErrorCode = "SambaDeleteFatal"
	SambaDeleteAllFatal                FamiPhotoErrorCode = "SambaDeleteAllFatal"
	SambaReadDirFatal                  FamiPhotoErrorCode = "SambaReadDirFatal"
	SambaGlobFatal                     FamiPhotoErrorCode = "SambaGlobFatal"
	ElasticSearchFatal                 FamiPhotoErrorCode = "ElasticSearchFatal"
	InvalidTimezoneFatal               FamiPhotoErrorCode = "InvalidTimezoneFatal"
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
