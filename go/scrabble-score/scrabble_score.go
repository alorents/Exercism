package scrabble

import (
	"unicode"
)

// Score computes the scrabble score of a given word according to the values defined in letterValues
func Score(input string) int {
	totalScore := 0
	letterValue := 0
	for _, letter := range input {
		switch unicode.ToUpper(letter) {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			letterValue = 1
		case 'D', 'G':
			letterValue = 2
		case 'B', 'C', 'M', 'P':
			letterValue = 3
		case 'F', 'H', 'V', 'W', 'Y':
			letterValue = 4
		case 'K':
			letterValue = 5
		case 'J', 'X':
			letterValue = 8
		case 'Q', 'Z':
			letterValue = 10
		}

		totalScore += letterValue
	}
	return totalScore
}
