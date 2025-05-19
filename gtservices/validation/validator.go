package validation

import (
	"regexp"
	"strings"
)

// IsValidEmail checks if the email is in a valid format
func IsValidEmail(email string) bool {
	email = strings.TrimSpace(email)
	// simple regex for email validation
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
