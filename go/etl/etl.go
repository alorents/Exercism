package etl

import (
	"strings"
)

func Transform(valueToLetter map[int][]string) map[string]int {
	letterToValue := make(map[string]int)
	for value, letterArray := range valueToLetter {
		for _, letter := range letterArray {
			letterToValue[strings.ToLower(letter)] = value

		}
	}
	return letterToValue
}