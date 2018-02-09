package validator

import "testing"

type mockNumber struct {
	ID int `validate:"number.(min=1,max=10)"`
}

func instNumber(id int) mockNumber {
	return mockNumber{
		ID: id,
	}
}

func TestNumber_LessThanMin(t *testing.T) {
	m := instNumber(0)
	expected := "ID should be greater than 1"

	testValidationFail(t, m, expected)
}

func TestNumber_Min(t *testing.T) {
	m := instNumber(1)

	testValidationPass(t, m)
}

func TestNumber_GreaterThanMax(t *testing.T) {
	m := instNumber(11)
	expected := "ID should be less than 10"

	testValidationFail(t, m, expected)
}

func TestNumber_Max(t *testing.T) {
	m := instNumber(10)

	testValidationPass(t, m)
}
