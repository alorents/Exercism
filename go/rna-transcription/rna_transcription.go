package strand

import (
	"strings"
)

var (
	replacer = strings.NewReplacer(
		"G", "C",
		"C", "G",
		"T", "A",
		"A", "U",
	)
)

// ToRNA provides the complimentary RNA strand for the given DNA strand
func ToRNA(dna string) string {
	return replacer.Replace(dna)
}
