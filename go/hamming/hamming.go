package hamming

import "errors"

func Distance(a, b string) (int, error) {
	hammingDistance := 0

	if len(a) != len(b) {
		return -1, errors.New("sequences are of unequal length")
	}

	for pos, _ := range a {
		if a[pos] != b[pos] {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}
