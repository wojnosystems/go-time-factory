package timeFactory

import (
	. "github.com/onsi/gomega"
	"testing"
	"time"
)

func Test_NowGet(t *testing.T) {
	cases := map[string]struct {
		input    Now
		expected time.Time
	}{
		"default: current time": {
			expected: time.Now(),
		},
		"override": {
			input: func() time.Time {
				return time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
			},
			expected: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	for caseName, td := range cases {
		t.Run(caseName, func(t *testing.T) {
			g := NewWithT(t)
			actual := td.input.Get()
			g.Expect(actual).Should(BeTemporally("~", td.expected, 1*time.Second))
		})
	}
}
