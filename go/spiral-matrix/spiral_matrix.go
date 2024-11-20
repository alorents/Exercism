package spiralmatrix

func SpiralMatrix(size int) [][]int {
	if size < 1 {
		return [][]int{}
	}
	if size == 1 {
		return [][]int{{1}}
	}

	matrix := make([][]int, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
	}

	row, col := 0, 0
	rowMin, colMin := 0, 0
	rowMax, colMax := size-1, size-1
	direction := "right"
	for i := 0; i < size*size; i++ {
		matrix[row][col] = i + 1
		switch direction {
		case "right":
			if col < colMax {
				col++
			} else if col == colMax {
				direction = "down"
				row++
				rowMin++
			}
		case "down":
			if row < rowMax {
				row++
			} else if row == rowMax {
				direction = "left"
				col--
				colMax--
			}
		case "left":
			if col > colMin {
				col--
			} else if col == colMin {
				direction = "up"
				row--
				rowMax--
			}
		case "up":
			if row > rowMin {
				row--
			} else if row == rowMin {
				direction = "right"
				col++
				colMin++
			}
		}
	}

	return matrix
}
