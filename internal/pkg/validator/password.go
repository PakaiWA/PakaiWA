package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Regex rules
var (
	upperCaseRegex = regexp.MustCompile(`[A-Z]`)
	lowerCaseRegex = regexp.MustCompile(`[a-z]`)
	numberRegex    = regexp.MustCompile(`[0-9]`)
	symbolRegex    = regexp.MustCompile(`[\W_]`)
)

func PasswordStrength(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	if len(password) < 8 {
		return false
	}

	if !upperCaseRegex.MatchString(password) {
		return false
	}

	if !lowerCaseRegex.MatchString(password) {
		return false
	}

	if !numberRegex.MatchString(password) {
		return false
	}

	if !symbolRegex.MatchString(password) {
		return false
	}

	return true
}
