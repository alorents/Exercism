package meetup

import (
	"time"
)

// Define the WeekSchedule type here.
type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	d := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	switch wSched {
	case Second:
		d = d.AddDate(0, 0, 7)
	case Third:
		d = d.AddDate(0, 0, 14)
	case Fourth:
		d = d.AddDate(0, 0, 21)
	case Teenth:
		d = d.AddDate(0, 0, 12)
	case Last:
		d = d.AddDate(0, 1, -7)
	}
	for d.Weekday() != wDay {
		d = d.AddDate(0, 0, 1)
	}
	return d.Day()
}
