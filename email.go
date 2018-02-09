package validator

import (
	"fmt"
	"regexp"
)

// Regular expression to validate email address.
var mailRe = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)

// EmailValidator checks if string is a valid email address.
type EmailValidator struct {
}

func (v EmailValidator) Validate(val interface{}) (bool, error) {
	value := val.(string)

	if len(value) == 0 {
		return false, fmt.Errorf("cannot be blank")
	}

	if !mailRe.MatchString(value) {
		return false, fmt.Errorf("is not a valid")
	}

	return true, nil
}
