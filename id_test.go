package validator

import "testing"

type mockID struct {
	ID int `json:"id" validate:"id"`
}

func instID(id int) mockID {
	return mockID{
		ID: id,
	}
}

func TestID_LessThanMin(t *testing.T) {
	m := instID(0)
	expected := "id should be greather than 0"

	testValidationFail(t, m, expected)
}

func TestID_Min(t *testing.T) {
	m := instID(1)

	testValidationPass(t, m)
}
