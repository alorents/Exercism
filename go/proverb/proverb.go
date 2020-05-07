package proverb

import "fmt"

const (
	line = "For want of a %s the %s was lost."
	finalLine = "And all for the want of a %s."
)

/*
Proverb generates a relevant proverb given the provided inputs

For example, given the list ["nail", "shoe", "horse", "rider", "message", "battle", "kingdom"]
the resulting proverbial rhyme will be:

For want of a nail the shoe was lost.
For want of a shoe the horse was lost.
For want of a horse the rider was lost.
For want of a rider the message was lost.
For want of a message the battle was lost.
For want of a battle the kingdom was lost.
And all for the want of a nail.
*/
func Proverb(rhyme []string) []string {
	if len(rhyme) <= 0 {
		return nil
	}
	proverb := make([]string, len(rhyme))
	for i := range rhyme[1:] {
			proverb[i] = fmt.Sprintf(line, rhyme[i], rhyme[i+1])
	}
	proverb[len(rhyme)-1] = fmt.Sprintf(finalLine, rhyme[0])

	return proverb
}
