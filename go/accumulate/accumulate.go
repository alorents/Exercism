package accumulate

// Accumulate preforms the converter function on the given string array
func Accumulate(given []string, converter func(string) string) (output []string) {
	output = make([]string, len(given))
	for i, s := range given {
		output[i] = converter(s)
	}
	return
}
