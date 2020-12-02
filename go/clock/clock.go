package clock

import (
	"fmt"
)

// Clock defines a time using hours and minutes using 24-hour time without dates
type Clock struct {
	hour   int
	minute int
}

// New is the clock constructor
func New(hour, min int) Clock {
	// normalize time in minutes
	minutes := min + hour*60
	minutes %= 24 * 60
	if minutes < 0 {
		minutes += 24 * 60
	}
	return Clock{minutes / 60, minutes % 60}
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}

// Add minutes to a Clock type
func (clock Clock) Add(min int) Clock {
	return New(clock.hour, clock.minute+min)
}

// Subtract minutes from a Clock type
func (clock Clock) Subtract(min int) Clock {
	return New(clock.hour, clock.minute-min)
}
