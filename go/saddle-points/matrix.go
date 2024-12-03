package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix [][]int

func New(s string) (*Matrix, error) {

	rows := strings.Split(s, "\n")
	result := make(Matrix, len(rows))
	for i, row := range rows {
		if row == "" {
			return &result, nil
		}
		cols := strings.Split(strings.TrimSpace(row), " ")
		result[i] = make([]int, len(cols))
		for j, col := range cols {
			num, err := strconv.Atoi(col)
			if err != nil {
				return nil, err
			}
			result[i][j] = num
		}
	}

	// Check if the matrix is valid
	for i := 1; i < len(result); i++ {
		if len(result[i]) != len(result[i-1]) {
			return nil, errors.New("uneven rows")
		}
	}

	return &result, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	cols := make([][]int, len(m[0]))
	for i := 0; i < len(cols); i++ {
		cols[i] = make([]int, len(m))
		for j := 0; j < len(cols[i]); j++ {
			cols[i][j] = m[j][i]
		}
	}
	return cols
}

func (m Matrix) Rows() [][]int {
	rows := make([][]int, len(m))
	for i := 0; i < len(rows); i++ {
		rows[i] = make([]int, len(m[i]))
		for j := 0; j < len(rows[i]); j++ {
			rows[i][j] = m[i][j]
		}
	}
	return rows
}

func (m Matrix) Set(row, col, val int) (ok bool) {
	if ok = row >= 0 && row < len(m) && col >= 0 && col < len(m[0]); ok {
		m[row][col] = val
	}

	return ok
}
