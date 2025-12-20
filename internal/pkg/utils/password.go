package utils

import (
	"regexp"

	dto "github.com/PakaiWA/PakaiWA/internal/app/pakaiwa/delivery/http/dto"
)

// Regex rules
var (
	upperCaseRegex = regexp.MustCompile(`[A-Z]`)
	lowerCaseRegex = regexp.MustCompile(`[a-z]`)
	numberRegex    = regexp.MustCompile(`[0-9]`)
	symbolRegex    = regexp.MustCompile(`[\W_]`)
)

type PasswordValidationError struct {
	Errors []dto.ValidationError
}

func (e *PasswordValidationError) Error() string {
	return "password does not meet complexity requirements"
}

func ValidateStrongPassword(password string) error {
	var validationErrors []dto.ValidationError

	if !upperCaseRegex.MatchString(password) {
		validationErrors = append(validationErrors, dto.ValidationError{
			Field:   "password",
			Tag:     "uppercase",
			Message: "not contain at least one uppercase letter",
		})
	}

	if !lowerCaseRegex.MatchString(password) {
		validationErrors = append(validationErrors, dto.ValidationError{
			Field:   "password",
			Tag:     "lowercase",
			Message: "not contain at least one lowercase letter",
		})
	}

	if !numberRegex.MatchString(password) {
		validationErrors = append(validationErrors, dto.ValidationError{
			Field:   "password",
			Tag:     "number",
			Message: "not contain at least one number",
		})
	}

	if !symbolRegex.MatchString(password) {
		validationErrors = append(validationErrors, dto.ValidationError{
			Field:   "password",
			Tag:     "symbol",
			Message: "not contain at least one special character",
		})
	}

	if len(validationErrors) > 0 {
		return &PasswordValidationError{Errors: validationErrors}
	}

	return nil
}
