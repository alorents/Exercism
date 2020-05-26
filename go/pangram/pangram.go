package pangram

import (
	"strings"
)

// IsPangram determines if a given sentence is a pangram (contains every letter of the alphabet at least once)
func IsPangram(sentence string) bool {
	sentence = strings.ToLower(sentence)
	for letter := 'a'; letter <= 'z'; letter++ {
		if !strings.ContainsRune(sentence, letter) {
			return false
		}
	}
	return true
}
