package prime

import "math"

// Nth finds the nth prime number
func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}
	// list of discovered primes
	primes := []int{2}
	// count of discovered primes
	primeCount := 1
	// next candidate to check for primality
	candidate := 3
	for primeCount < n {
		primeFound := true
		// check candidate for primality by checking smaller primes for divisors
		candidateSqrt := int(math.Sqrt(float64(candidate)))
		for _, prime := range primes {
			if prime > candidateSqrt {
				primeFound = true
				break
			}
			if candidate%prime == 0 {
				primeFound = false
				break
			}
		}
		if primeFound {
			primes = append(primes, candidate)
			primeCount++
		}
		candidate += 2
	}
	return primes[primeCount-1], true
}
