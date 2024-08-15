package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("n must be a positive integer")
	}
	count := 0
	for {
		if n == 1 {
			return count, nil
		}
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		count++
	}
}
