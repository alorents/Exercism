package strain

type Ints []int
type Lists [][]int
type Strings []string

func (ints Ints) Keep(keepfunc func(int) bool) Ints {
	var result []int
	for _, num := range ints {
		if keepfunc(num) {
			result = append(result, num)
		}
	}
	return result
}

func (ints Ints) Discard(discardfunc func(int) bool) Ints {
	var result []int
	for _, num := range ints {
		if !discardfunc(num) {
			result = append(result, num)
		}
	}
	return result
}

func (strs Strings) Keep(keepfunc func(string) bool) Strings {
	var result []string
	for _, str := range strs {
		if keepfunc(str) {
			result = append(result, str)
		}
	}
	return result
}

func (strs Strings) Discard(discardfunc func(string) bool) Strings {
	var result []string
	for _, str := range strs {
		if !discardfunc(str) {
			result = append(result, str)
		}
	}
	return result
}

func (lists Lists) Keep(keepfunc func([]int) bool) Lists {
	var result [][]int
	for _, list := range lists {
		if keepfunc(list) {
			result = append(result, list)
		}
	}
	return result
}

func (lists Lists) Discard(discardfunc func([]int) bool) Lists {
	var result [][]int
	for _, list := range lists {
		if !discardfunc(list) {
			result = append(result, list)
		}
	}
	return result
}