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
	return fmt.Sprintf("code: %s, %s", e.errorCode, e.baseError.Error())
}

func (e *FamiPhotoError) ErrorCode() FamiPhotoErrorCode {
	return e.errorCode
}

type FamiPhotoErrorCode string

const (
	Unknown           FamiPhotoErrorCode = "Unknown"
	UserNotFoundError FamiPhotoErrorCode = "UserNotFoundError"
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
