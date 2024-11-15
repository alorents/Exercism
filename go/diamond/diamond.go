package diamond

import (
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

	var top string
	var bottom string
	outerSpaces := int(char - 'A')
	innerSpaces := 0

	for c := 'A'; c <= rune(char); c++ {
		line := ""
		for i := 0; i < outerSpaces; i++ {
			line += " "
		}
		line += string(c)
		if innerSpaces > 0 {
			for i := 0; i < innerSpaces; i++ {
				line += " "
			}
			line += string(c)
		}
		for i := 0; i < outerSpaces; i++ {
			line += " "
		}

		outerSpaces--
		if innerSpaces == 0 {
			innerSpaces++
		} else {
			innerSpaces += 2
		}
		if c == rune(char) {
			top += line + "\n"
		} else {
			top += line + "\n"
			bottom = line + "\n" + bottom
		}
	}

	s := top + bottom
	s, _ = strings.CutSuffix(s, "\n")
	return s, nil
}
