package wordcount

import (
	"strings"
	"unicode"
)

// Frequency maps a word to its count
type Frequency map[string]int

// WordCount returns a Frequency map of the word counts within the given string
func WordCount(text string) Frequency {
	frequencyMap := make(Frequency)
	words := strings.FieldsFunc(strings.ToLower(text), func(r rune) bool {
		return !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '\'')
	})
	for _, word := range words {
		frequencyMap[strings.Trim(word, "'")]++
	}
	return frequencyMap
}
