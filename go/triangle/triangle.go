// Package triangle contains methods to work with the triangle geometric shape
package triangle

import "math"

// Kind refers to the kind of triangle
type Kind int

// List of triangle Kinds
const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// Validate inputs
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) ||
		math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}

	// All sides must be greater than 0
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}
	// The sum of any 2 sides must be greater than the third
	if a+b < c || a+c < b || b+c < a {
		return NaT
	}

	// An equilateral triangle has all three sides the same length
	if a == b && a == c {
		return Equ
	}

	// An isosceles triangle has at least two sides the same length
	if a == b || a == c || b == c {
		return Iso
	}

	// A scalene triangle has all sides of different lengths
	if a != b && a != c && b != c {
		return Sca
	}

	// Fallback case
	return NaT
}
