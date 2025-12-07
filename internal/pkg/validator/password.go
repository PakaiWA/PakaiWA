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

func ValidateStrongPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	if len(password) < 8 {
		return false
	}

	return upperCaseRegex.MatchString(password) &&
		lowerCaseRegex.MatchString(password) &&
		numberRegex.MatchString(password) &&
		symbolRegex.MatchString(password)
}
