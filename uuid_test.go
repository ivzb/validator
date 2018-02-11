package validator

import "testing"

type mockUUID struct {
	ID string `json:"id" validate:"uuid"`
}

func instUUID(id string) mockUUID {
	return mockUUID{
		ID: id,
	}
}

func TestUUID_Blank(t *testing.T) {
	m := instUUID("")
	expected := "id cannot be blank"

	testValidationFail(t, m, expected)
}

func TestUUID_InvalidLengthID(t *testing.T) {
	m := instUUID("a")
	expected := "id is not a valid uuid"

	testValidationFail(t, m, expected)
}

func TestUUID_InvalidID(t *testing.T) {
	m := instUUID("zxc3036f-0f1c-4e98-b71c-d4cd61213f90")
	expected := "id is not a valid uuid"

	testValidationFail(t, m, expected)
}

func TestUUID_ValidID(t *testing.T) {
	m := instUUID("fbd3036f-0f1c-4e98-b71c-d4cd61213f90")

	testValidationPass(t, m)
}
