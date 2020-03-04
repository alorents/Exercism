package space

type Planet string

const secondsInEarthYears = 31557600
var secondsInPlanetYears = map[Planet]float64{
	"Mercury":	secondsInEarthYears * 0.2408467,
	"Venus": 	secondsInEarthYears * 0.61519726,
	"Earth": 	secondsInEarthYears,
	"Mars": 	secondsInEarthYears * 1.8808158,
	"Jupiter": 	secondsInEarthYears * 11.862615,
	"Saturn": 	secondsInEarthYears * 29.447498,
	"Uranus": 	secondsInEarthYears * 84.016846,
	"Neptune": 	secondsInEarthYears * 164.79132,
}

func Age(ageInSeconds float64, planet Planet) float64 {
	return ageInSeconds / secondsInPlanetYears[planet]
}
