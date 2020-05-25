package protein

import "errors"

// ErrStop signals the end of a protein sequence
var ErrStop = errors.New("stop codon encountered")

// ErrInvalidBase indicates the encounters nucleotide sequence does not match any known codons
var ErrInvalidBase = errors.New("invalid codon base")

// FromRNA converts a nucleotide sequence to an array of its proteins
func FromRNA(input string) ([]string, error) {
	var output []string
	for i := 0; i < len(input)/3; i++ {
		codon := input[i*3 : i*3+3]
		protein, err := FromCodon(codon)
		if err != nil {
			if err == ErrStop {
				return output, nil
			}
			return output, err
		}
		output = append(output, protein)
	}

	return output, nil
}

// FromCodon converts a codon string into its protein string
func FromCodon(input string) (string, error) {
	switch input {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
