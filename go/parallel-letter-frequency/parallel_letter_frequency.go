package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency takes a []string and runs Frequency on each string concurrently
func ConcurrentFrequency(texts []string) FreqMap {
	freqMapChan := make(chan FreqMap)

	for _, text := range texts {
		go func(text string) {
			freqMapChan <- Frequency(text)
		}(text)
	}

	freqMap := FreqMap{}
	for i := 0; i < len(texts); i++ {
		for k, v := range <-freqMapChan {
			freqMap[k] += v
		}
	}

	return freqMap
}
