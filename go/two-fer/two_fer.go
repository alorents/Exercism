// Package twofer describes two-for-one deals
package twofer

// ShareWith outputs the name of someone to share the twofer deal with
// If blank will assume "you"
// Assumes the sharer is included
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
