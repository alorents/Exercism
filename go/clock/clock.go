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
	clock = clock.Add(hour * 60)
	clock = clock.Add(min)
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
	if min < 0 {
		return clock.Subtract(-min)
	}
	clock.minute += min % 60
	if clock.minute >= 60 {
		clock.minute -= 60
		clock.hour++
	}
	clock.hour += min / 60
	clock.hour = clock.hour % 24
	return clock
}

// Subtract minutes from a Clock type
func (clock Clock) Subtract(min int) Clock {
	if min == 0 {
		return clock
	}
	if min < 0 {
		return clock.Add(-min)
	}
	clock.minute -= min % 60
	if clock.minute < 0 {
		clock.minute += 60
		clock.hour--
	}
	clock.hour -= min / 60
	clock.hour = clock.hour % 24
	if clock.hour < 0 {
		clock.hour += 24
	}

	return clock
}
