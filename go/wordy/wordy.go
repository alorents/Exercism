package wordy

import (
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	var val1 int
	var op string
	question = strings.TrimLeft(question, "What is ")
	question = strings.TrimRight(question, "?")
	question = strings.ReplaceAll(question, "multiplied by", "multipliedby")
	question = strings.ReplaceAll(question, "divided by", "dividedby")
	strings := strings.Split(question, " ")
	val1, err := strconv.Atoi(strings[0])
	if err != nil {
		// value not found
		return 0, false
	}
	for i := 1; i < len(strings); i++ {
		op = strings[i]
		if i+1 >= len(strings) {
			// expected value not found
			return 0, false
		}
		val2, err := strconv.Atoi(strings[i+1])
		if err != nil {
			// expected value not found
			return 0, false
		}
		switch op {
		case "plus":
			val1 = val1 + val2
		case "minus":
			val1 = val1 - val2
		case "multipliedby":
			val1 = val1 * val2
		case "dividedby":
			val1 = val1 / val2
		default:
			// unexpected operator
			return 0, false
		}
		i++
	}

	// default error case
	return val1, true
}
