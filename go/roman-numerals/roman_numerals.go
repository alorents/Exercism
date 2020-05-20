package romannumerals

import (
	"errors"
	"fmt"
)

var arabicToRomanNumeral = []struct {
	arabic       int
	romanNumeral string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ToRomanNumeral provides the Roman Numeral equivalent of a given number
func ToRomanNumeral(number int) (result string, err error) {
	if number < 1 || number > 3000 {
		err = errors.New(fmt.Sprintf("Invalid input: %d must be between 1 and 3000", number))
	}
	for _, values := range arabicToRomanNumeral {
		for ; number >= values.arabic; number -= values.arabic {
			result += values.romanNumeral
		}
	}
	return
}

