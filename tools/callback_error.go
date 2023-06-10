package tools

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type CallbackError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func readErrorBasedOnTags(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "Data must be filled in"
	case "min":
		return "Enter at least " + fe.Param() + " characters"
	case "max":
		return "The maximum amount of data is " + fe.Param() + " characters"
	case "email":
		return "Invalid Email"
	case "uuid":
		return "Invalid UUID Format"
	case "eqfield":
		return "Value not matched"
	case "jwt":
		return "Invalid token value"
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
				out[i] = CallbackError{fe.Field(), readErrorBasedOnTags(fe)}
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
			out[i] = CallbackError{fe.Field(), readErrorBasedOnTags(fe)}
		}
		return out
	}

	return nil
}
