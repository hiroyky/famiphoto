package validators

import (
	"github.com/go-playground/validator"
	"github.com/hiroyky/famiphoto/errors"
	"github.com/labstack/echo"
	"reflect"
	"strings"
)

func NewValidator() echo.Validator {
	v := validator.New()

	v.RegisterTagNameFunc(getTagName)
	return &appValidator{
		validator: v,
	}
}

type appValidator struct {
	validator *validator.Validate
}

func (v *appValidator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return errors.New(errors.InvalidRequestError, err)
	}
	return nil
}

func getTagName(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}
