package allergies

/*
Allergy score mapping:
	eggs (1)
	peanuts (2)
	shellfish (4)
	strawberries (8)
	tomatoes (16)
	chocolate (32)
	pollen (64)
	cats (128)
*/
var allergenList = []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}

//Allergies returns a list of allergens based on a person's allergy score
func Allergies(score uint) []string {
	results := []string{}
	if score <= 0 {
		return results
	}
	for _, substance := range allergenList {
		if AllergicTo(score, substance) == true {
			results = append(results, substance)
		}
	}
	return results
}

//AllergicTo determines if a person is allergic to a specific substance based on their allergy score
func AllergicTo(score uint, substance string) bool {
	if score <= 0 {
		return false
	}
	switch substance {
	case "eggs":
		return score&1 == 1
	case "peanuts":
		return score&2 == 2
	case "shellfish":
		return score&4 == 4
	case "strawberries":
		return score&8 == 8
	case "tomatoes":
		return score&16 == 16
	case "chocolate":
		return score&32 == 32
	case "pollen":
		return score&64 == 64
	case "cats":
		return score&128 == 128
	}
	return true
}
