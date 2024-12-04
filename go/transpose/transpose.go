package transpose

func Transpose(input []string) []string {
	max := 0
	for _, row := range input {
		if len(row) > max {
			max = len(row)
		}
	}
	result := make([]string, max)

	for i, row := range input {
		for j, char := range row {
			for len(result[j]) < i {
				result[j] += " "
			}
			result[j] += string(char)
		}

	}
	return result
}
