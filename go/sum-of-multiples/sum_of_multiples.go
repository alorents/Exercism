package summultiples

// SumMultiples return the sum of all unique multiples matching the given divisors under a limit
func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	for i := 1; i < limit; i++ {
		for _, divisor := range divisors {
			if divisor > 0 && i%divisor == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
