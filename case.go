package titlecase

import (
	"strings"
	"unicode"
)

var stopwords = map[string]bool{
	"a":   true,
	"an":  true,
	"the": true,
}

func titleWord(word string) string {
	if stopwords[word] {
		return word
	}

	return strings.Title(word)
}

// Title converts the input string into title case.
func Title(s string) string {
	s = strings.ToLower(s)

	result := ""
	word := ""
	for _, nextRune := range s {
		if unicode.IsSpace(nextRune) {
			// we hit a separator, we have a complete word
			// decide what to do to the word and append it
			result += titleWord(word)

			// append this separator to the result
			result += string(nextRune)

			// reset the word
			word = ""

			continue
		}

		word += string(nextRune)
	}
	if len(word) != 0 {
		result += titleWord(word)
	}

	return result
}
