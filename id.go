package validator

import (
	"errors"
)

// IDValidator performs numerical value validation.
// Its limited to int type for simplicity.
type IDValidator struct {
}

func (v IDValidator) Validate(val interface{}) (bool, error) {
	num := val.(int)

	if num <= 0 {
		return false, errors.New("should be greather than 0")
	}

	return true, nil
}
