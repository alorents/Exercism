package resistorcolortrio

import "fmt"

var colorMap = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	if len(colors) < 3 {
		return "unexpected input"
	}
	for _, v := range colors {
		_, ok := colorMap[v]
		if !ok {
			return "unexpected input"
		}
	}

	unit := "ohms"
	ohms := colorMap[colors[0]]*10 + colorMap[colors[1]]*1
	numZeros := colorMap[colors[2]]
	if colorMap[colors[1]] == 0 {
		ohms /= 10
		numZeros++
	}
	switch numZeros {
	case 0:
		unit = "ohms"
		ohms *= 1
	case 1:
		unit = "ohms"
		ohms *= 10
	case 2:
		unit = "ohms"
		ohms *= 100
	case 3:
		unit = "kiloohms"
		ohms *= 1
	case 4:
		unit = "kiloohms"
		ohms *= 10
	case 5:
		unit = "kiloohms"
		ohms *= 100
	case 6:
		unit = "megaohms"
		ohms *= 1
	case 7:
		unit = "megaohms"
		ohms *= 10
	case 8:
		unit = "megaohms"
		ohms *= 100
	case 9:
		unit = "gigaohms"
		ohms *= 1
	case 10:
		unit = "gigaohms"
		ohms *= 10
	}
	return fmt.Sprintf("%v %s", ohms, unit)
}
