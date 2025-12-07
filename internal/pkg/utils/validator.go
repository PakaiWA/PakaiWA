package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ValidationErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ExtractValidationErrors(err error) []ValidationErrorDetail {
	var ve validator.ValidationErrors
	if !errors.As(err, &ve) {
		return nil
	}

	out := make([]ValidationErrorDetail, 0, len(ve))
	for _, fe := range ve {
		out = append(out, ValidationErrorDetail{
			Field:   fe.Field(),
			Message: validationMessage(fe),
		})
	}

	return out
}

func validationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "email":
		return fe.Field() + " must be a valid email address"
	case "min":
		return fe.Field() + " must be at least " + fe.Param() + " characters"
	case "strongpwd":
		return "password must contain uppercase, lowercase, number, and symbol"
	}
	return fe.Field() + " is invalid"
}
