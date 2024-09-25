package robotname

import (
	"fmt"
	"math/rand"
)

// Define the Robot type here.
type Robot struct {
	name string
}

const maxIdx = 26 * 26 * 10 * 10 * 10

var namePool = generateNames()

var nameIdx = 0

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		err := r.Reset()
		if err != nil {
			return "", err
		}
	}
	return r.name, nil
}

func generateNames() []string {
	names := make([]string, maxIdx)
	pos := 0
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			for k := 0; k < 1000; k++ {
				names[pos] = fmt.Sprintf("%s%s%03d",
					string(rune(i+65)),
					string(rune(j+65)),
					k,
				)
				pos++
			}
		}
	}
	rand.Shuffle(len(names), func(i, j int) {
		names[i], names[j] = names[j], names[i]
	})
	return names
}

func (r *Robot) Reset() error {
	if nameIdx >= maxIdx {
		return fmt.Errorf("ran out of available names")
	}
	r.name = namePool[nameIdx]
	nameIdx++
	return nil
}
