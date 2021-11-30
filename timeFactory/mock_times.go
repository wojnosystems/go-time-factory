package timeFactory

import "time"

var zeroTime time.Time

// MockTimes will return each time provided, in order, whenever the returned function is called.
// Should the returned function be called more than the number of times provided, returns time.Zero
func MockTimes(times ...time.Time) func() time.Time {
	i := 0
	return func() time.Time {
		if i >= len(times) {
			return zeroTime
		}
		now := times[i]
		i++
		return now
	}
}
