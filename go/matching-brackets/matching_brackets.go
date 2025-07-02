package brackets

func Bracket(input string) bool {
	var brackets []rune
	for _, r := range input {
		if r == '(' || r == '[' || r == '{' {
			// push rune to stack to check for a match later
			brackets = append(brackets, r)
		} else if r == ')' || r == ']' || r == '}' {
			if len(brackets) < 1 {
				// closing rune found without any opening rune on the stack
				return false
			}
			openingRune := brackets[len(brackets)-1]
			brackets = brackets[:len(brackets)-1]
			if r == ')' && openingRune != '(' || r == ']' && openingRune != '[' || r == '}' && openingRune != '{' {
				// check closing rune matches popped opening rune
				return false
			}
		} else {
			// non-bracket char, do nothing
		}
	}
	// if len > 0; opening runes still on the stack with no closing matches; else all matches found
	return len(brackets) == 0
}