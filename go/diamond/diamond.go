package diamond

import (
	"bytes"
	"fmt"
	"strings"
)

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", fmt.Errorf("invalid character, expected input A..Z")
	}
	if char == 'A' {
		return "A", nil
	}

	length := int(char-'A')*2 + 1
	rows := make([]string, length)

	for c := byte('A'); c <= char; c++ {
		row := bytes.Repeat([]byte{' '}, length)
		row[char-c], row[length-1-int(char-c)] = c, c
		rows[c-'A'], rows[length-1-int(c-'A')] = string(row), string(row)
	}

	return strings.Join(rows, "\n"), nil
}
