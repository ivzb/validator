package validator

import "testing"

type mockEmail struct {
	Email string `validate:"email"`
}

func instEmail(email string) mockEmail {
	return mockEmail{
		Email: email,
	}
}

func TestEmail_Blank(t *testing.T) {
	m := instEmail("")
	expected := "Email cannot be blank"

	testValidationFail(t, m, expected)
}

func TestString_InvalidEmail(t *testing.T) {
	m := instEmail("a")
	expected := "Email is not a valid"

	testValidationFail(t, m, expected)
}

func TestString_ValidEmail(t *testing.T) {
	m := instEmail("test@mail.com")

	testValidationPass(t, m)
}
