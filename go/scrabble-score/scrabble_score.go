package scrabble

import (
	"unicode"
)

// Score computes the scrabble score of a given word according to the values defined in letterValues
func Score(input string) int {
	totalScore := 0
	for _, letter := range input {
		switch unicode.ToUpper(letter) {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			totalScore += 1
		case 'D', 'G':
			totalScore += 2
		case 'B', 'C', 'M', 'P':
			totalScore += 3
		case 'F', 'H', 'V', 'W', 'Y':
			totalScore += 4
		case 'K':
			totalScore += 5
		case 'J', 'X':
			totalScore += 8
		case 'Q', 'Z':
			totalScore += 10
		}
	}
	return totalScore
}
