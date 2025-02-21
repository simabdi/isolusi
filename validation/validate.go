package validation

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

func Validate(data interface{}) (dataError string) {
	var errors []string
	validate := validator.New()

	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err.Field()+" : "+errorValidate(err.Tag(), err.Type().String(), err.Param()))
		}
	}

	dataError = strings.Join(errors, ", ")
	return dataError
}

func errorValidate(tag interface{}, typeName string, param string) string {
	switch tag {
	case "min":
		return "This field must is min " + param + " character"
	case "max":
		return "This field must is max " + param + " character"
	case "number":
		return "This field must is number"
	case "string":
		return "This field must is string"
	case "required":
		return "This field is required " + typeName
	case "email":
		return "Invalid email"
	case "oneof":
		return "Failed on the 'oneof'"
	case "eqfield":
		return "Password confirmation doesn't match"
	case "date":
		return "Format date must be Y-m-d"
	}

	return ""
}
