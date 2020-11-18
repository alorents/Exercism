package rotationalcipher

import (
	"strings"
	"unicode"
)

// RotationalCipher encodes the provided plainText by rotating individual letters according to the shiftKey
func RotationalCipher(plainText string, shiftKey int) string {
	shiftKey %= 26
	if shiftKey == 0 {
		return plainText
	}
	var encodedText strings.Builder
	for _, r := range plainText {
		encodedRune := r
		if unicode.IsLetter(r) {
			encodedRune += int32(shiftKey)
			if (unicode.IsUpper(r) && encodedRune > 'Z') || encodedRune > 'z' {
				encodedRune -= 26
			}
		}
		encodedText.WriteRune(encodedRune)
	}

	return encodedText.String()
}
