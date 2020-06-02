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
	doubleDigit := false
	var sum int
	for i := len(input) - 1; i >= 0; i-- {
		if !unicode.IsDigit([]rune(input)[i]) {
			return false
		}
		value := int([]rune(input)[i] - '0')
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
