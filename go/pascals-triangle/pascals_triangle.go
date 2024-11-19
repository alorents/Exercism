package pascal

func Triangle(n int) [][]int {
	if n < 1 {
		return [][]int{}
	}
	rows := make([][]int, n)
	for i := 0; i < n; i++ {
		row := make([]int, i+1)
		for j := 0; j < i+1; j++ {
			// left and right side have < 2 addends and are always 1
			if j == 0 || j == i {
				row[j] = 1
			} else {
				row[j] = rows[i-1][j-1] + rows[i-1][j]
			}
		}
		rows[i] = row
	}
	return rows
}
