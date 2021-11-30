package timeFactory

import (
	. "github.com/onsi/gomega"
	"testing"
	"time"
)

func Test_startAtAndAddDurations(t *testing.T) {
	cases := map[string]struct {
		start    time.Time
		adds     []time.Duration
		expected []time.Time
	}{
		"no times return zero-time": {
			start: time.Date(2021, 11, 28, 21, 35, 15, 0, time.UTC),
			adds: []time.Duration{
				1 * time.Second,
				4 * time.Second,
				1 * time.Second,
			},
			expected: []time.Time{
				time.Date(2021, 11, 28, 21, 35, 15, 0, time.UTC),
				time.Date(2021, 11, 28, 21, 35, 16, 0, time.UTC),
				time.Date(2021, 11, 28, 21, 35, 20, 0, time.UTC),
				time.Date(2021, 11, 28, 21, 35, 21, 0, time.UTC),
			},
		},
	}

	for caseName, td := range cases {
		t.Run(caseName, func(t *testing.T) {
			g := NewWithT(t)
			actual := getTimesUntilEmpty(MockTimeWithDurations(td.start, td.adds...))
			g.Expect(actual).Should(ConsistOf(td.expected))
		})
	}
}
