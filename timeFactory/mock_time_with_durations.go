package timeFactory

import "time"

// MockTimeWithDurations simulates times, but allows you to start with a time,
// then add durations to each subsequent returned value, pushing the times further into the future.
func MockTimeWithDurations(startTime time.Time, addToEach ...time.Duration) func() time.Time {
	times := make([]time.Time, len(addToEach)+1)
	times[0] = startTime
	for i, each := range addToEach {
		times[i+1] = times[i].Add(each)
	}
	return MockTimes(times...)
}
