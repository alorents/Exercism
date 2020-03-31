package hamming

import "errors"

// Distance calculates the Hamming Distance between two DNA strands.
func Distance(a, b string) (int, error) {
	ar, br := []rune(a), []rune(b)

	if len(ar) != len(br) {
		return 0, errors.New("sequences are of unequal length")
	}

	hammingDistance := 0
	for pos := range ar {
		if ar[pos] != br[pos] {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}
