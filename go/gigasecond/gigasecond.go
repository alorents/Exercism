package gigasecond

import "time"

// AddGigasecond returns a time 1 gigasecond (1e9) seconds after the provided time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
