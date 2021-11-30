package timeFactory

import (
	. "github.com/onsi/gomega"
	"testing"
	"time"
)

func Test_simulateNows(t *testing.T) {
	cases := map[string]struct {
		input []time.Time
	}{
		"no times return zero-time": {
			input: []time.Time{},
		},
		"some times returned": {
			input: []time.Time{
				time.Date(2021, 11, 28, 21, 35, 15, 0, time.UTC),
				time.Date(2021, 11, 28, 21, 35, 15, 1, time.UTC),
				time.Date(2021, 11, 28, 21, 35, 15, 2, time.UTC),
			},
		},
		"another call to prove i is not shared among calls": {
			input: []time.Time{
				time.Date(2021, 11, 28, 21, 35, 15, 0, time.UTC),
				time.Date(2021, 11, 28, 21, 35, 15, 1, time.UTC),
				time.Date(2021, 11, 28, 21, 35, 15, 2, time.UTC),
			},
		},
	}

	for caseName, td := range cases {
		t.Run(caseName, func(t *testing.T) {
			g := NewWithT(t)
			actual := getTimesUntilEmpty(MockTimes(td.input...))
			g.Expect(actual).Should(ConsistOf(td.input))
		})
	}
}
