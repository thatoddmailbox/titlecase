package titlecase_test

import (
	"testing"

	"github.com/thatoddmailbox/titlecase"
)

func individualTest(t *testing.T, input string, expected string) {
	result := titlecase.Title(input)
	if result != expected {
		t.Errorf(
			"For input '%s': expected '%s', got '%s'",
			input,
			expected,
			result,
		)
	}
}

func TestBasic(t *testing.T) {
	individualTest(t, "word", "Word")
	individualTest(t, "WORD", "Word")
	individualTest(t, "interesting phrase", "Interesting Phrase")
	individualTest(t, "testing yay", "Testing Yay")
}
