package reverse

import "strings"

// Reverse takes a string as input and returns a string with the rune order reversed
func Reverse(str string) string {
	runeArray := []rune(str)
	var sb strings.Builder
	for i := len(runeArray) - 1; i >= 0; i-- {
		sb.WriteRune(runeArray[i])
	}
	return sb.String()
}
