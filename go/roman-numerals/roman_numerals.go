package romannumerals

import (
	"errors"
	"fmt"
)

// ToRomanNumeral provides the Roman Numeral equivalent of a given number
func ToRomanNumeral(number int) (result string, err error) {
	if number < 1 || number > 3000 {
		err = errors.New(fmt.Sprintf("Invalid input: %d must be between 1 and 3000", number))
	}
	thousands := number / 1000
	number -= thousands * 1000
	hundreds := number / 100
	number -= hundreds * 100
	tens := number / 10
	number -= tens * 10
	ones := number
	result += convertDigitToRomanNumeral(thousands, "M", "", "")
	result += convertDigitToRomanNumeral(hundreds, "C", "D", "M")
	result += convertDigitToRomanNumeral(tens, "X", "L", "C")
	result += convertDigitToRomanNumeral(ones, "I", "V", "X")

	return
}

func convertDigitToRomanNumeral(digit int, oneNumeral, fiveNumeral, tenNumeral string) string {
	switch digit {
	case 1:
		return oneNumeral
	case 2:
		return oneNumeral + oneNumeral
	case 3:
		return oneNumeral + oneNumeral + oneNumeral
	case 4:
		return oneNumeral + fiveNumeral
	case 5:
		return fiveNumeral
	case 6:
		return fiveNumeral + oneNumeral
	case 7:
		return fiveNumeral + oneNumeral + oneNumeral
	case 8:
		return fiveNumeral + oneNumeral + oneNumeral + oneNumeral
	case 9:
		return oneNumeral + tenNumeral
	default:
		return ""
	}
}
