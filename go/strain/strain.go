package strain

// Ints is an array of integers
type Ints []int

// Lists is an array of integer arrays
type Lists [][]int

// Strings is an array of strings
type Strings []string

// Keep strains a collection of Ints to only include the members that match a given predicate function
func (ints Ints) Keep(predicate func(int) bool) Ints {
	var result Ints
	for _, num := range ints {
		if predicate(num) {
			result = append(result, num)
		}
	}
	return result
}

// Discard strains a collection of Ints to only include the members that don't match a given predicate function
// This uses the Keep function to form a compliment set
func (ints Ints) Discard(predicate func(int) bool) Ints {
	return ints.Keep(func(n int) bool { return !predicate(n) })
}

// Keep strains a collection of Strings to only include the members that match a given predicate function
func (strs Strings) Keep(predicate func(string) bool) Strings {
	var result Strings
	for _, str := range strs {
		if predicate(str) {
			result = append(result, str)
		}
	}
	return result
}

// Keep strains a collection of Lists to only include the members that match a given predicate function
func (lists Lists) Keep(predicate func([]int) bool) Lists {
	var result Lists
	for _, list := range lists {
		if predicate(list) {
			result = append(result, list)
		}
	}
	return result
}
