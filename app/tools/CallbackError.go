package tools

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type CallbackError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ReadErrorBasedOnTags(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "Data Harus diisi"
	case "min":
		return "Data Minimal " + fe.Param() + " Karakter"
	case "max":
		return "Data Maximal " + fe.Param() + " Karakter"
	case "email":
		return "Invalid Email"
	case "uuid":
		return "Invalid UUID Format"
	case "eqfield":
		return "Value not matched"
	}
	return fe.Error() // default error
}

func GenerateErrorMessage(i interface{}) []CallbackError {
	v := validator.New()
	if err := v.Struct(i); err != nil {
		var va validator.ValidationErrors
		if errors.As(err, &va) {
			out := make([]CallbackError, len(va))
			for i, fe := range va {
				out[i] = CallbackError{fe.Field(), ReadErrorBasedOnTags(fe)}
			}
			return out
		}
	}

	return nil
}

func GenerateErrorMessageV2(err error) []CallbackError {
	var va validator.ValidationErrors
	if errors.As(err, &va) {
		out := make([]CallbackError, len(va))
		for i, fe := range va {
			out[i] = CallbackError{fe.Field(), ReadErrorBasedOnTags(fe)}
		}
		return out
	}

	return nil
}
