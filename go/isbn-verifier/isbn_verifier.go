package isbn

import (
	"strings"
	"unicode"
)

// IsValidISBN takes an isbn-10 string and returns its validity as a boolean
// to check validity it follows this formula:
// (x1 * 10 + x2 * 9 + x3 * 8 + x4 * 7 + x5 * 6 + x6 * 5 + x7 * 4 + x8 * 3 + x9 * 2 + x10 * 1) mod 11 == 0
func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}
	sum := 0
	for i, r := range isbn {
		// check digit
		if i == 9 && r == 'X' {
			sum += 10
			continue
		}
		if !unicode.IsDigit(r) {
			return false
		}
		digit := int(r - '0')
		sum += digit * (10 - i)
	}
	return sum%11 == 0
}
