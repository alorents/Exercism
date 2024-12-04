package transpose

import (
	"strings"
)

func Transpose(input []string) []string {
	result := []string{}
	if len(input) == 0 {
		return result
	}

	max := 0
	for _, row := range input {
		if len(row) > max {
			max = len(row)
		}
	}

	for i := 0; i < max; i++ {
		row := ""
		for j := range input {
			if i < len(input[j]) {
				row += string(input[j][i])
			} else {
				row += "#"
			}
		}
		result = append(result, strings.ReplaceAll(strings.TrimRight(row, "#"), "#", " "))
	}

	return result
}
