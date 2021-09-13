package titlecase

import (
	"strings"
	"unicode"
)

// Title converts the input string into title case.
func Title(s string) string {
	s = strings.ToLower(s)

	result := ""
	word := ""
	for _, nextRune := range s {
		if unicode.IsSpace(nextRune) {
			// we hit a separator, we have a complete word
			// decide what to do to the word and append it
			result += strings.Title(word)

			// append this separator to the result
			result += string(nextRune)

			// reset the word
			word = ""

			continue
		}

		word += string(nextRune)
	}
	if len(word) != 0 {
		result += strings.Title(word)
	}

	return strings.Title(s)
}
