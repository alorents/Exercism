package strand

import (
	"unicode"
)

// ToRNA provides the complimentary RNA strand for the given DNA strand
func ToRNA(dna string) string {
	rna := ""
	for _, nucleotide := range dna {
		switch unicode.ToUpper(nucleotide) {
		case 'G':
			rna += "C"
		case 'C':
			rna += "G"
		case 'T':
			rna += "A"
		case 'A':
			rna += "U"
		}
	}
	return rna
}
