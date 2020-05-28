package encode

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// RunLengthEncode compress plaintext using run length encoding
func RunLengthEncode(plaintext string) string {
	if len(plaintext) == 0 {
		return ""
	}
	count := 0
	var sb strings.Builder
	currentRune := []rune(plaintext)[0]
	for _, r := range plaintext {
		if r == currentRune {
			count++
		} else {
			if count > 1 {
				sb.WriteString(strconv.Itoa(count))
			}
			sb.WriteRune(currentRune)
			count = 1
			currentRune = r
		}
	}
	if count > 1 {
		sb.WriteString(fmt.Sprintf("%d", count))
	}
	sb.WriteRune(currentRune)

	return sb.String()
}

// RunLengthDecode decompress an encoded string using run length encoding
func RunLengthDecode(encoded string) string {
	count := 1
	var db strings.Builder
	var sb strings.Builder
	for _, r := range encoded {
		if unicode.IsDigit(r) {
			db.WriteRune(r)
		} else {
			if db.Len() > 0 {
				count, _ = strconv.Atoi(db.String())
				db.Reset()
			}
			for i := 0; i < count; i++ {
				sb.WriteRune(r)
			}
			count = 1
		}
	}
	return sb.String()
}
