package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(plaintext string) string {
	if len(plaintext) < 1 {
		return ""
	}
	replacer := strings.NewReplacer(
		" ", "",
		"!", "",
		"@", "",
		"#", "",
		"$", "",
		"%", "",
		"^", "",
		"&", "",
		"*", "",
		"(", "",
		")", "",
		",", "",
		".", "",
	)
	encodedText := replacer.Replace(plaintext)
	rows := int(math.Sqrt(float64(len(encodedText))))
	columns := rows
	if rows*columns < len(encodedText) {
		columns++
		if rows*columns < len(encodedText) {
			rows++
		}
	}
	paddingSize := rows*columns - len(encodedText)
	encodedText += strings.Repeat(" ", paddingSize)
	squareText := make([]string, rows)
	for i := 0; i < rows-1; i++ {
		squareText[i] = encodedText[i*columns : (i+1)*columns]
	}
	// handle final row which may be shorter than the others
	squareText[rows-1] = encodedText[(rows-1)*columns:]

	var builder strings.Builder
	for i := 0; i < columns; i++ {
		for j := 0; j < rows; j++ {
			if j < rows {
				builder.WriteRune(unicode.ToLower(rune(squareText[j][i])))
			}
		}
		if i < columns-1 {
			builder.WriteRune(' ')
		}
	}
	encodedText = builder.String()
	return encodedText
}
