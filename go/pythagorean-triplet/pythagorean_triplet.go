package pythagorean

import "math"

// Triplet represent a pythagorean triplet a < b < c && a^2 + b^2 = c^2
type Triplet [3]int

// Range returns all pythagorean triplets within the provided range
func Range(min, max int) []Triplet {
	var triplets []Triplet
	for a := min; a < max; a++ {
		for b := a + 1; b <= max; b++ {
			c2 := a*a + b*b
			c := int(math.Sqrt(float64(a*a + b*b)))
			if c > max {
				break
			}
			if c*c == c2 {
				triplets = append(triplets, Triplet{a, b, c})
			}
		}
	}
	return triplets
}

// Sum returns all pythagorean triplets that also satifies a + b + c = p

func Sum(p int) []Triplet {
	var triplets []Triplet
	max := p / 2
	for a := 1; a < max; a++ {
		for b := a + 1; b <= max; b++ {
			c2 := a*a + b*b
			c := int(math.Sqrt(float64(a*a + b*b)))
			if c > max {
				break
			}
			if c*c == c2 && a+b+c == p {
				triplets = append(triplets, Triplet{a, b, c})
			}
		}
	}
	return triplets
}
