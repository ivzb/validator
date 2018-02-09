package validator

import "testing"

type mockString struct {
	ID string `validate:"string.(min=2,max=5)"`
}

func instString(id string) mockString {
	return mockString{
		ID: id,
	}
}

func TestString_Blank(t *testing.T) {
	m := instString("")
	expected := "ID cannot be blank"

	testValidationFail(t, m, expected)
}

func TestString_LessThanMinLength(t *testing.T) {
	m := instString("a")
	expected := "ID should be at least 2 chars long"

	testValidationFail(t, m, expected)
}

func TestString_MinLength(t *testing.T) {
	m := instString("aa")

	testValidationPass(t, m)
}

func TestString_GreaterThanMaxLength(t *testing.T) {
	m := instString("aaaaaa")
	expected := "ID should be less than 5 chars long"

	testValidationFail(t, m, expected)
}

func TestString_MaxLength(t *testing.T) {
	m := instString("aaaaa")

	testValidationPass(t, m)
}
