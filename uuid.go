package validator

import (
	"fmt"

	"github.com/google/uuid"
)

// UUIDValidator checks if string is a valid uuid,
type UUIDValidator struct {
}

func (v UUIDValidator) Validate(val interface{}) (bool, error) {
	value := val.(string)
	length := len(value)

	if length == 0 {
		return false, fmt.Errorf("cannot be blank")
	}

	if length != 36 {
		return false, fmt.Errorf("is not a valid uuid")
	}

	_, err := uuid.Parse(value)

	if err != nil {
		return false, fmt.Errorf("is not a valid uuid")
	}

	return true, nil
}
