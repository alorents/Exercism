package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

// Encode uses the crypto-square method to encode plaintext
func Encode(plaintext string) string {
	if len(plaintext) < 2 {
		return plaintext
	}
	text := normalizeText(plaintext)
	rows, columns := calculateDimensions(len(text))
	// add padding if needed to make a perfect rectangle
	paddingSize := rows*columns - len(text)
	text += strings.Repeat(" ", paddingSize)

	// build encoded text "column" x "row"
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

// normalize text
func normalizeText(plaintext string) string {
	text := strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, plaintext)
	return text
}

// calculateDimensions determines the correct row and column size
func calculateDimensions(length int) (int, int) {
	rows := int(math.Sqrt(float64(length)))
	columns := rows
	if rows*columns < length {
		columns++
		if rows*columns < length {
			rows++
		}
	}
	return rows, columns
}
