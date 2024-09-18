package allyourbase

import (
	"fmt"
)

/*
	Assumptions:

1. Zero is always represented in outputs as [0] instead of [].
2. In no other instances are leading zeroes present in any outputs.
3. Leading zeroes are accepted in inputs.
4. An empty sequence of input digits is considered zero, rather than an error.
5. input base must be >= 2
6. output base must be >= 2
7. all digits must satisfy 0 <= d < input base
*/
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) (outputDigits []int, err error) {
	// verify input
	if inputBase < 2 {
		return nil, fmt.Errorf("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, fmt.Errorf("output base must be >= 2")
	}
	for _, digit := range inputDigits {
		if digit < 0 || digit >= inputBase {
			return nil, fmt.Errorf("all digits must satisfy 0 <= d < input base")
		}
	}

	// edge cases
	if len(inputDigits) == 0 {
		return []int{0}, nil
	}
	if len(inputDigits) == 1 {
		if inputDigits[0] == 0 {
			return []int{0}, nil
		}
		if inputDigits[0] == 1 {
			return []int{1}, nil
		}
	}

	// standarize input to base10
	decimalNum := convertToDecimal(inputBase, inputDigits)
	if decimalNum == 0 {
		return []int{0}, nil
	}

	// perform conversion
	outputDigits = convertToTargetBase(decimalNum, outputBase)

	return outputDigits, nil
}

func convertToDecimal(inputBase int, inputDigits []int) int {
	num := 0
	for i := 0; i < len(inputDigits); i++ {
		digit := inputDigits[i]
		coefficient := powInts(inputBase, len(inputDigits)-1-i)
		num += digit * coefficient
	}
	return num
}

func convertToTargetBase(decimalNum int, targetBase int) []int {
	maxCoefficient := 0
	numDigits := 0
	for i := 0; powInts(targetBase, i) < decimalNum; i++ {
		maxCoefficient = powInts(targetBase, i)
		numDigits++
	}
	outputDigits := make([]int, numDigits)
	for i := 0; i < numDigits; i++ {
		outputDigits[i] = decimalNum / maxCoefficient
		decimalNum %= maxCoefficient
		maxCoefficient /= targetBase
	}

	return outputDigits
}

func powInts(base, power int) int {
	if power == 0 {
		return 1
	}

	if power == 1 {
		return base
	}

	result := base
	for i := 2; i <= power; i++ {
		result *= base
	}
	return result
}
