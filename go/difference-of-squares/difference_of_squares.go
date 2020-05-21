package diffsquares

// SquareOfSum returns the square of the sum of integers 1..n
func SquareOfSum(n int) (result int) {
	for i := 1; i <= n; i++ {
		result += i
	}
	return result * result
}

// SumOfSquares returns the sum of the squares of integers 1..n
func SumOfSquares(n int) (result int) {
	for i := 1; i <= n; i++ {
		result += i * i
	}
	return result
}

// Difference returns the difference of between the square of sums and the sum of squares of integers 1..n
func Difference(n int) (result int) {
	return SquareOfSum(n) - SumOfSquares(n)
}
