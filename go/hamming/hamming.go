package hamming

import "errors"

// Distance calculates the Hamming Distance between two DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("sequences are of unequal length")
	}

	hammingDistance := 0
	for pos := range a {
		if a[pos] != b[pos] {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}
