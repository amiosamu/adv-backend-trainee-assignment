package validator

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

const (
	picLinksMinLimit    = 1
	picLinksMaxLimit    = 3
	descriptionMinLimit = 100
	descriptionMaxLimit = 1000
	nameMinLimit        = 10
	nameLimit           = 200
)

type Validator struct {
	v              *validator.Validate
	pictureErr     error
	nameErr        error
	descriptionErr error
}

func NewValidator() *Validator {
	validateInstance := validator.New()
	v := &Validator{v: validateInstance}
	validateInstance.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})
	return v
}
