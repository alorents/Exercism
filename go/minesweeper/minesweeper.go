package minesweeper

import (
	"strconv"
)

// Annotate returns an annotated board
func Annotate(board []string) []string {
	height := len(board)
	if height < 1 {
		return board
	}
	width := len(board[0])
	if width < 1 {
		return board
	}
	annotatedBoard := make([]string, height)
	for y, row := range board {
		annotatedRow := ""
		for x, char := range row {
			if char == '*' {
				annotatedRow += "*"
			} else {
				count := 0
				for y2 := y - 1; y2 <= y+1; y2++ {
					if y2 < 0 || y2 >= height {
						continue
					}
					for x2 := x - 1; x2 <= x+1; x2++ {
						if x2 < 0 || x2 >= width {
							// edge case -- off the board
							continue
						}
						if board[y2][x2] == '*' {
							count++
						}
					}
				}
				if count == 0 {
					annotatedRow += " "
				} else {
					annotatedRow += strconv.Itoa(count)
				}
			}
		}
		annotatedBoard[y] = annotatedRow
	}
	return annotatedBoard
}
