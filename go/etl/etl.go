package etl

import (
	"strings"
)

// Transforms a legacy scrabble score valueToLetter map to a letterToValue map
func Transform(valueToLetter map[int][]string) map[string]int {
	letterToValue := make(map[string]int)
	for value, letterArray := range valueToLetter {
		for _, letter := range letterArray {
			letterToValue[strings.ToLower(letter)] = value

		}
	}
	return letterToValue
}
