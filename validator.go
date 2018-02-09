package validator

import (
	"fmt"
	"reflect"
	"strings"
)

// Name of the struct tag used in examples.
const tagName = "validate"

// Generic data validator.
type Validator interface {
	// Validate method performs validation and returns result and optional error.
	Validate(interface{}) (bool, error)
}

// Returns validator struct corresponding to validation type
func getValidatorsFromTag(tag string) []Validator {
	validations := strings.Split(tag, "&")
	validators := make([]Validator, len(validations))

	for i, validation := range validations {
		validators[i] = mapToValidator(validation)
	}

	return validators
}

// Returns validator struct corresponding to validation type
func mapToValidator(validation string) Validator {
	args := strings.Split(validation, ".")

	switch args[0] {
	case "id":
		return IDValidator{}
	case "uuid":
		return UUIDValidator{}
	case "number":
		validator := NumberValidator{}
		fmt.Sscanf(strings.Join(args[1:], ","), "(min=%d,max=%d)", &validator.Min, &validator.Max)

		return validator
	case "string":
		validator := StringValidator{}
		fmt.Sscanf(strings.Join(args[1:], ","), "(min=%d,max=%d)", &validator.Min, &validator.Max)

		return validator
	case "email":
		return EmailValidator{}
	}

	return DefaultValidator{}
}

// Performs actual data validation using validator definitions on the struct
func validateStruct(s interface{}) []error {
	errs := []error{}

	// ValueOf returns a Value representing the run-time data
	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {
		// Get the field tag value
		tag := v.Type().Field(i).Tag.Get(tagName)

		// Skip if tag is not defined or ignored
		if tag == "" || tag == "-" {
			continue
		}

		// Get a validators that corresponds to a tag
		validators := getValidatorsFromTag(tag)
		field := v.Field(i).Interface()
		fieldName := v.Type().Field(i).Name

		for _, validator := range validators {
			// Perform validation
			valid, err := validator.Validate(field)

			// Append error to results
			if !valid && err != nil {
				errs = append(errs, fmt.Errorf("%s %s", fieldName, err.Error()))
			}
		}
	}

	return errs
}
