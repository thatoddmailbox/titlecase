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

	individualTest(t, "This is a test", "This Is a Test")
	individualTest(t, "give up the ghost", "Give Up the Ghost")
	individualTest(t, "time out of mind", "Time Out of Mind")

	// this one doesn't work because phrasal verbs are not handled correctly
	// individualTest(t, "puttin' on the ritz", "Puttin' On the Ritz")

	individualTest(t, "INTRO TO PSYCHOLOGY", "Intro to Psychology")
}

func TestExceptions(t *testing.T) {
	// these start with a stopword, but it's capitalized anyways
	individualTest(t, "a weird story", "A Weird Story")
	individualTest(t, "THE TEST TITLE", "The Test Title")
}

func TestSeparators(t *testing.T) {
	// hyphenation is tricky
	// in general, you're supposed to leave the word after the hyphen lowercase
	// however, this was originally designed for class names, which should have the next word capitalized
	individualTest(t, "trans-siberian pipeline", "Trans-Siberian Pipeline")
	individualTest(t, "INTRO-COMPUT SCI/PROG IN JAVA", "Intro-Comput Sci/Prog in Java")

	individualTest(t, "ADVANCED PROGRAMMING&HEXING", "Advanced Programming&Hexing")
	individualTest(t, "INT'L AFFAIR", "Int'l Affair")
}

func TestReadme(t *testing.T) {
	individualTest(t, "INTRODUCTION TO PROGRAMMING", "Introduction to Programming")
	individualTest(t, "some cool book", "Some Cool Book")
	individualTest(t, "an interesting story", "An Interesting Story")
}
