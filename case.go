package titlecase

import (
	"strings"
	"unicode"
)

var stopwords = map[string]bool{
	"a":    true,
	"an":   true,
	"and":  true,
	"as":   true,
	"but":  true,
	"for":  true,
	"from": true,
	"in":   true,
	"into": true,
	"like": true,
	"nor":  true,
	"of":   true,
	"on":   true,
	"or":   true,
	"over": true,
	"so":   true,
	"the":  true,
	"to":   true,
	"upon": true,
	"with": true,
	"yet":  true,
}

func titleWord(word string, wordIndex int) string {
	if stopwords[word] && wordIndex != 0 {
		return word
	}

	return strings.Title(word)
}

func isSeparator(r rune) bool {
	// based off of the go strings package's implementation of strings.Title

	if r <= 0x7F {
		switch {
		case '0' <= r && r <= '9':
			return false
		case 'a' <= r && r <= 'z':
			return false
		case 'A' <= r && r <= 'Z':
			return false
		case r == '_':
			return false
		case r == '\'':
			return false
		}
		return true
	}

	return unicode.IsSpace(r)
}

// Title converts the input string into title case.
func Title(s string) string {
	s = strings.ToLower(s)

	result := ""
	word := ""
	wordIndex := 0
	for _, nextRune := range s {
		if isSeparator(nextRune) {
			// we hit a separator, we have a complete word
			// decide what to do to the word and append it
			result += titleWord(word, wordIndex)

			// append this separator to the result
			result += string(nextRune)

			// reset the word
			word = ""
			wordIndex++

			continue
		}

		word += string(nextRune)
	}
	if len(word) != 0 {
		result += titleWord(word, wordIndex)
	}

	return result
}
