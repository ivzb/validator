package validator

import (
	"testing"
)

type mockEmptyTag struct {
	ID string `validate:"-"`
}

func instEmptyTag(id string) mockEmptyTag {
	return mockEmptyTag{
		ID: id,
	}
}

func testValidation(t *testing.T, m interface{}, length int, expected string) {
	errs := ValidateStruct(m)

	if len(errs) != length {
		t.Fatalf("expected %d validation error, got %d", length, len(errs))
	}

	if length > 0 {
		actual := errs[0].Error()

		if expected != actual {
			t.Fatalf("\nexpected: %s\ngot: %s", expected, actual)
		}
	}
}

func testValidationFail(t *testing.T, m interface{}, expected string) {
	testValidation(t, m, 1, expected)
}

func testValidationPass(t *testing.T, m interface{}) {
	testValidation(t, m, 0, "")
}

func TestEmptyTag(t *testing.T) {
	m := instEmptyTag("mock_id")

	testValidationPass(t, m)
}
