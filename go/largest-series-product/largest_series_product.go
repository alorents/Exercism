package lsproduct

import (
	"errors"
)

// LargestSeriesProduct calculates the the largest product of n(span) consecutive digits with the given series
func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span > len(digits) || span < 0 {
		return 0, errors.New("span must be positive and cannot exceed the series length")
	}
	largestProduct := int64(0)
	for i := 0; i < len(digits)-span+1; i++ {
		localProduct := int64(1)
		for j := i; j < i+span; j++ {
			if digits[j] > '9' || digits[j] <'0' {
				return 0, errors.New("series must contain only digits 0-9")
			}
			digit := int64(digits[j]-'0')
			localProduct *= int64(digit)
		}
		if localProduct > largestProduct {
			largestProduct = localProduct
		}
	}
	return largestProduct, nil
}
