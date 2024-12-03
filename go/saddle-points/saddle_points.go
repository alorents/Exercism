package matrix

import "math"

type Pair [2]int

func (m *Matrix) Saddle() (pairs []Pair) {
	trees := m.Rows()
	if len(*m) == 0 {
		return pairs
	}

	rowMax, colMin := make([]int, len(trees)), make([]int, len(trees[0]))
	for i := range colMin {
		colMin[i] = math.MaxInt
	}
	
	for i, row := range trees {
		for j, tree := range row {
			if tree > rowMax[i] {
				rowMax[i] = tree
			}
			if tree < colMin[j] {
				colMin[j] = tree
			}
		}
	}

	for i, row := range trees {
		for j, tree := range row {
			if tree == rowMax[i] && tree == colMin[j] {
				pairs = append(pairs, Pair{i+1,j+1})
			}
		}
	}

	return pairs
}
