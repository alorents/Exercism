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
	// if doubledValue > 9, subtract 9
	var sum int
	for i := len(input) - 1; i >= 0; i-- {
		if !unicode.IsDigit([]rune(input)[i]) {
			return false
		}
		value := int([]rune(input)[i] - '0')
		if (len(input)-i)%2 != 1 {
			value = value * 2
			if value > 9 {
				value -= 9
			}
		}
		sum += value
	}
	return sum%10 == 0
}
