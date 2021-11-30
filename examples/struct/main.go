package main

import (
	"github.com/wojnosystems/go-time-factory/timeFactory"
	"log"
	"time"
)

func main() {
	var nowFactory timeFactory.Now
	// default, uninitialized timeFactory.Now returns the current time, as though you called time.Now()
	log.Println(nowFactory.Get())

	// nowFactory is over-ridden with a custom time factory
	nowFactory = func() time.Time {
		return time.Date(1999, 12, 31, 0, 0, 0, 0, time.UTC)
	}
	// calling it returns the time returned from the custom time factory
	log.Println(nowFactory.Get())

	// More complicated methods can be used to set up a series of fake times
	nowFactory = timeFactory.ReturnTimeWithDurations(
		time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		24*time.Hour,
		1*time.Hour,
		2*time.Hour,
	)
	for i := 0; i < 5; i++ {
		log.Println(nowFactory.Get())
	}
}
