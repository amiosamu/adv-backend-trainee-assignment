package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"regexp"
	"strings"
)

const (
	picLinksMinLimit    = 1
	picLinksMaxLimit    = 3
	descriptionMinLimit = 100
	descriptionMaxLimit = 1000
	nameMinLimit        = 10
	nameMaxLimit        = 200
	priceMinLimit       = 0
)

var (
	nameLengthRegexp = regexp.MustCompile(fmt.Sprintf(`^.{%d,%d}$`, nameMinLimit, nameMaxLimit))
	descLengthRegexp = regexp.MustCompile(fmt.Sprintf(`^.{%d,%d}$`, descriptionMinLimit, descriptionMaxLimit))
	priceRegexp      = regexp.MustCompile(fmt.Sprintf(`.{%d,}$`, priceMinLimit))
)

type Validator struct {
	v              *validator.Validate
	pictureErr     error
	nameErr        error
	descriptionErr error
	priceErr       error
}

func NewValidator() *Validator {
	validateInstance := validator.New()
	v := &Validator{v: validateInstance}
	validateInstance.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return v
}

func (v *Validator) newValidationError(field string, value interface{}, tag string, param string) error {
	switch tag {
	case "required":
		return fmt.Errorf("field %s is required", field)
	case "name":
		return v.nameErr
	case "description":
		return v.descriptionErr
	case "pictures":
		return v.pictureErr
	case "price":
		return v.priceErr
	case "min":
		return fmt.Errorf("field %s must be at least %s characters", field, param)
	case "max":
		return fmt.Errorf("field %s must be at most %s characters", field, param)
	default:
		return fmt.Errorf("field %s is invalid", field)
	}
}

func (v *Validator) nameValidate(fl validator.FieldLevel) bool {
	if fl.Field().Kind() != reflect.String {
		v.nameErr = fmt.Errorf("field %s must be a string", fl.FieldName())
		return false
	}
	fieldValue := fl.Field().String()

	if ok := nameLengthRegexp.MatchString(fieldValue); !ok {
		v.nameErr = fmt.Errorf("field %s must be between %d and %d characters ", fl.FieldName(), nameMinLimit, nameMaxLimit)
		return false
	}
	return true
}
