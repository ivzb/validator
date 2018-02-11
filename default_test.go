package validator

import "testing"

type mockDefault struct {
	ID string `json:"id" validate:"non_existing_validation_tag"`
}

func instDefault(id string) mockDefault {
	return mockDefault{
		ID: id,
	}
}

// all fields that are not part of validation
// should be marked empty, '-' or non-existing validation tag
func TestDefault(t *testing.T) {
	m := instDefault("mock_id")

	testValidationPass(t, m)
}
