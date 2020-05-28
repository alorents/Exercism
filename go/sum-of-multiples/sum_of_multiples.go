package summultiples

// SumMultiples return the sum of all unique multiples matching the given divisors under a limit
func SumMultiples(limit int, divisors ...int) int {
	if limit < 2 {
		return 0
	}
	sum := 0
	for i := 1; i < limit; i++ {
		isMultiple := func() bool {
			for _, divisor := range divisors {
				if divisor == 0 {
					continue
				}
				if i%divisor == 0 {
					return true
				}
			}
			return false
		}
		if isMultiple() {
			sum += i
		}
	}
	return sum
}
