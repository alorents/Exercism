package raindrops

import "strconv"

// Convert function takes an int input and converts it to a raindrop sound based on the value
// The rules of raindrops are that if a given number:
//		has 3 as a factor, add 'Pling' to the result.
//		has 5 as a factor, add 'Plang' to the result.
//		has 7 as a factor, add 'Plong' to the result.
//		does not have any of 3, 5, or 7 as a factor, the result should be the digits of the number.
func Convert(number int) string {
	result := ""

	if number%3 == 0 {
		result += "Pling"
	}
	if number%5 == 0 {
		result += "Plang"
	}
	if number%7 == 0 {
		result += "Plong"
	}
	if result == "" {
		result = strconv.Itoa(number)
	}

	return result
}
