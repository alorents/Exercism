package foodchain

import (
	"fmt"
	"strings"
)

var animals = []string{
	"fly",
	"spider",
	"bird",
	"cat",
	"dog",
	"goat",
	"cow",
	"horse",
}

var introLines = map[string]string{
	"fly":    "",
	"spider": "It wriggled and jiggled and tickled inside her.\n",
	"bird":   "How absurd to swallow a bird!\n",
	"cat":    "Imagine that, to swallow a cat!\n",
	"dog":    "What a hog, to swallow a dog!\n",
	"goat":   "Just opened her throat and swallowed a goat!\n",
	"cow":    "I don't know how she swallowed a cow!\n",
	"horse":  "She's dead, of course!",
}

var uniqueLines = map[string]string{
	"fly":  "I don't know why she swallowed the fly. Perhaps she'll die.\n",
	"bird": "She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.\n",
}

func Verse(v int) string {
	verses := strings.Split(Song(), "\n\n")
	return verses[v-1]
}

func Verses(start, end int) string {
	verses := strings.Split(Song(), "\n\n")
	return strings.Join(verses[start-1:end], "\n\n")
}

func Song() string {
	song := ""
	tail := ""
	previousAnimal := ""
	for _, animal := range animals {
		song += fmt.Sprintf("I know an old lady who swallowed a %s.\n", animal)
		song += introLines[animal]

		if line, ok := uniqueLines[animal]; ok {
			tail = line + tail
		} else {
			tail = fmt.Sprintf("She swallowed the %s to catch the %s.\n", animal, previousAnimal) + tail
		}
		if animal != "horse" {
			song += tail + "\n"
			previousAnimal = animal
		}
	}
	return song
}
