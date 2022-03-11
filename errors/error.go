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
	Unknown            FamiPhotoErrorCode = "Unknown"
	UserNotFoundError  FamiPhotoErrorCode = "UserNotFoundError"
	UserCreateFatal    FamiPhotoErrorCode = "UserCreateFatal"
	UserAlreadyExists  FamiPhotoErrorCode = "UserAlreadyExists"
	PasswordWeakError  FamiPhotoErrorCode = "PasswordWeakError"
	TxnRollbackFatal   FamiPhotoErrorCode = "TxnRollbackFatal"
	TxnBeginFatal      FamiPhotoErrorCode = "TxnBeginFatal"
	HashPasswordFatal  FamiPhotoErrorCode = "HashPasswordFatal"
	MatchPasswordFatal FamiPhotoErrorCode = "MatchPasswordFatal"
)

func New(errCode FamiPhotoErrorCode, baseError error) error {
	return &FamiPhotoError{
		errorCode: errCode,
		baseError: baseError,
	}
}

func UnwrapFPError(err error) *FamiPhotoError {
	var dst FamiPhotoError
	if ok := native.As(err, &dst); ok {
		return &dst
	}
	return nil
}

func GetFPErrorCode(err error) FamiPhotoErrorCode {
	srError := UnwrapFPError(err)
	if srError == nil {
		return Unknown
	}
	return srError.ErrorCode()
}
