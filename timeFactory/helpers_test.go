package timeFactory

import "time"

// getTimesUntilEmpty continuously calls the NowGetter until a zeroTime is returned.
// Appends each returned time to returned slice.
func getTimesUntilEmpty(nowFactory func() time.Time) []time.Time {
	actual := make([]time.Time, 0, 10)
	for {
		item := nowFactory()
		if item == zeroTime {
			break
		}
		actual = append(actual, item)
	}
	return actual
}
