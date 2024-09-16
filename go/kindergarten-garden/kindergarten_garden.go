package kindergarten

import (
	"fmt"
	"sort"
	"strings"
)

// Define the Garden type here.
type Garden map[string][]string

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func NewGarden(diagram string, children []string) (*Garden, error) {
	garden := make(Garden)

	rows := strings.Split(diagram, "\n")
	if len(rows) != 3 || len(rows[1]) != len(rows[2]) || len(rows[1])%2 != 0 {
		return nil, fmt.Errorf("unexpected diagram input")
	}
	if len(rows[1]) != 2*len(children) {
		return nil, fmt.Errorf("diagram - children mismatch")
	}

	plantsLookup := map[rune]string{
		'C': "clover",
		'G': "grass",
		'R': "radishes",
		'V': "violets",
	}

	sortedChildren := make([]string, len(children))
	copy(sortedChildren, children)
	sort.Strings(sortedChildren)

	for i, child := range sortedChildren {
		p0, ok0 := plantsLookup[rune(rows[1][2*i])]
		p1, ok1 := plantsLookup[rune(rows[1][2*i+1])]
		p2, ok2 := plantsLookup[rune(rows[2][2*i])]
		p3, ok3 := plantsLookup[rune(rows[2][2*i+1])]
		if !(ok0 && ok1 && ok2 && ok3) {
			return nil, fmt.Errorf("invalid plant code")
		}
		garden[child] = []string{p0, p1, p2, p3}
	}

	if len(garden) != len(sortedChildren) {
		return nil, fmt.Errorf("duplicate child name")
	}

	return &garden, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	plants, ok := (*g)[child]
	return plants, ok
}
