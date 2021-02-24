package allergies

var allergenList = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

//Allergies returns a list of allergens based on a person's allergy score
func Allergies(score uint) []string {
	results := []string{}
	if score <= 0 {
		return results
	}
	for substance := range allergenList {
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
	var substanceScore = allergenList[substance]
	return score&substanceScore == substanceScore
}
