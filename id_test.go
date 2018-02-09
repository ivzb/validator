package validator

import "testing"

type mockID struct {
	ID int `validate:"id"`
}

func instID(id int) mockID {
	return mockID{
		ID: id,
	}
}

func TestID_LessThanMin(t *testing.T) {
	m := instID(0)
	expected := "ID should be greather than 0"

	testValidationFail(t, m, expected)
}

func TestID_Min(t *testing.T) {
	m := instID(1)

	testValidationPass(t, m)
}

//func TestNumber_GreaterThanMax(t *testing.T) {
//m := instNumber(11)
//expected := "ID should be less than 10"

//testValidationFail(t, m, expected)
//}

//func TestNumber_Max(t *testing.T) {
//m := instNumber(10)

//testValidationPass(t, m)
//}
