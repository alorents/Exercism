package sieve

func Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}
	var primes []int
	marked := map[int]bool{}
	for i := 2; i <= limit; i++ {
		_, ok := marked[i]
		if !ok {
			primes = append(primes, i)
			for j := i; j <= limit; j += i {
				marked[j] = true
			}
		}
	}

	return primes
}
