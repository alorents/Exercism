package clock

import "strconv"

// Clock defines a time using hours and minutes using 24-hour time without dates
type Clock struct {
	hour   int
	minute int
}

// New is the clock constructor
func New(hour, min int) Clock {
	clock := Clock{hour: 0, minute: 0}

	// Adjust for over/underflow minutes
	hour += min / 60
	min = min % 60
	if min < 0 {
		min += 60
		hour--
	}

	// Adjust for over/underflow hours
	hour = hour % 24
	if hour < 0 {
		hour += 24
	}
	clock.hour = hour
	clock.minute = min

	return clock
}

func (clock Clock) String() string {
	hourString := strconv.Itoa(clock.hour)
	if clock.hour < 10 {
		hourString = "0" + hourString
	}
	minString := strconv.Itoa(clock.minute)
	if clock.minute < 10 {
		minString = "0" + minString
	}
	return hourString + ":" + minString
}

// Add minutes to a Clock type
func (clock Clock) Add(min int) Clock {
	if min == 0 {
		return clock
	}

	return New(clock.hour, clock.minute+min)
}

// Subtract minutes from a Clock type
func (clock Clock) Subtract(min int) Clock {
	if min == 0 {
		return clock
	}

	return clock.Add(-min)
}
