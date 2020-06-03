package luhn

import (
	"strings"
	"unicode"
)

// Valid determines if the given number is valid according to the Luhn Algorithm
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	if len(input) <= 1 {
		return false
	}
	// double every second digit from the right
	doubleDigit := len(input)%2 == 0
	var sum int
	for _, r := range input {
		if !unicode.IsDigit(r) {
			return false
		}
		value := int(r - '0')
		if doubleDigit {
			value *= 2
			if value > 9 {
				value -= 9
			}
		}
		doubleDigit = !doubleDigit
		sum += value
	}
	return sum%10 == 0
}
