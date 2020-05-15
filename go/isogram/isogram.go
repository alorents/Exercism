package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns a boolean indicating if the provided string is an isogram
func IsIsogram(input string) bool {
	letters := ""
	for _, letter := range input {
		letter = unicode.ToLower(letter)
		if letter == ' ' || letter == '-' {
			continue
		} else if strings.ContainsRune(letters, letter) {
			return false
		} else {
			letters += string(letter)
		}
	}
	return true
}
