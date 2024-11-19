package yacht

func Score(dice []int, category string) int {
	switch category {
	case "ones":
		return count(dice, 1)
	case "twos":
		return count(dice, 2)
	case "threes":
		return count(dice, 3)
	case "fours":
		return count(dice, 4)
	case "fives":
		return count(dice, 5)
	case "sixes":
		return count(dice, 6)
	case "full house":
		return fullHouse(dice)
	case "four of a kind":
		return fourOfAKind(dice)
	case "little straight":
		return straight(dice, category)
	case "big straight":
		return straight(dice, category)
	case "choice":
		return choice(dice)
	case "yacht":
		return yacht(dice)
	}

	return 0
}

func count(dice []int, val int) int {
	result := 0
	for _, die := range dice {
		if die == val {
			result += val
		}
	}
	return result
}

func fullHouse(dice []int) int {
	val1, val2 := dice[0], 0
	count1, count2 := 0, 0

	for _, die := range dice {
		if die != val1 && val2 == 0 {
			val2 = die
		}
		if die == val1 {
			count1++
		} else if die == val2 {
			count2++
		} else {
			return 0
		}
	}

	if count1 == 2 && count2 == 3 || count1 == 3 && count2 == 2 {
		return val1*count1 + val2*count2
	}

	return 0
}

func fourOfAKind(dice []int) int {
	val := dice[0]
	count := 0
	for _, die := range dice {
		if die == val {
			count++
		}
	}
	if count >= 4 {
		return 4 * val
	} else if count == 1 {
		return 4 * dice[1]
	}
	return 0
}

func straight(dice []int, category string) int {
	vals := make([]bool, 6)
	for _, die := range dice {
		if vals[die-1] {
			return 0
		} else {
			vals[die-1] = true
		}
	}
	if vals[0] && vals[5] {
		return 0
	}
	if vals[0] && category == "little straight" ||
		vals[5] && category == "big straight" {
		return 30
	}

	return 0
}

func choice(dice []int) int {
	total := 0
	for _, die := range dice {
		total += die
	}
	return total
}

func yacht(dice []int) int {
	for i := 1; i < len(dice)-1; i++ {
		if dice[i] != dice[0] {
			return 0
		}
	}
	return 50
}
