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
	text := replacer.Replace(plaintext)
	rows := int(math.Sqrt(float64(len(text))))
	columns := rows
	if rows*columns < len(text) {
		columns++
		if rows*columns < len(text) {
			rows++
		}
	}
	paddingSize := rows*columns - len(text)
	text += strings.Repeat(" ", paddingSize)
	var builder strings.Builder
	rowIndex, columnIndex := 0, 0
	for i := 0; i < len(text); i++ {
		index := rowIndex*columns + columnIndex
		builder.WriteRune(unicode.ToLower(rune(text[index])))
		rowIndex++
		if rowIndex == rows {
			rowIndex = 0
			columnIndex++
			if columnIndex < columns {
				builder.WriteRune(' ')
			}
		}
	}
	return builder.String()
}
