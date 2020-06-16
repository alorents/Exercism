package grains

import (
	"errors"
)

const (
	minSquare = 1
	maxSquare = 64
)

// Square returns the number of grains on a given square of a chessboard
// assuming there is 1 grain on square 1 each square following doubles the amount on the previous square
// this can be expressed as: x = 2^(n-1)
func Square(square int) (uint64, error) {
	if square < minSquare || square > maxSquare {
		return 0, errors.New("input must match a square on a chessboard")
	}
	return 1 << (square - 1), nil
}

// Total calculates the total number of grains on the board following the premise described in the square function
// the sum of 2^0 + 2^1 + 2^2... + 2^n can be expressed as: x = 2^(n+1)-1
// For a chess board this is 2^64 - 1
func Total() uint64 {
	return 1<<64 - 1
}
