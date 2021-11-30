# Overview

Helper package/module that provides ways of simulating calls to `time.Now()`. Useful in testing your applications or packages that need to simulate time. Provides an initialization-safe struct that allows it to be overwritten

# How to use

## Install

`go get -u github.com/wojnosystems/go-time-factory`

## Example Struct

```go
package main

import (
	"github.com/wojnosystems/go-time-factory/timeFactory"
	"log"
	"time"
)

func main() {
	var nowFactory timeFactory.Now
	// default, uninitialized timeFactory.Now returns the current time
	log.Println(nowFactory.Get())

	// nowFactory is over-ridden with a custom time factory
	nowFactory = func()time.Time {
		return time.Date(1999, 12, 31, 0, 0, 0, 0, time.UTC)
	}
	// calling it returns the time returned from the custom time factory
	log.Println(nowFactory.Get())

	// More complicated methods can be used to setup a series of fake times
	nowFactory = timeFactory.MockTimeWithDurations(
		time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		24 * time.Hour,
		1 * time.Hour,
		2 * time.Hour,
	)
	for i := 0; i < 5; i++ {
		log.Println(nowFactory.Get())
	}
}
```

The above example prints:

```text
2021/11/29 22:54:00 2021-11-29 22:54:00.000860785 -0800 PST m=+0.000096911
2021/11/29 22:54:00 1999-12-31 00:00:00 +0000 UTC
2021/11/29 22:54:00 2021-01-01 00:00:00 +0000 UTC
2021/11/29 22:54:00 2021-01-02 00:00:00 +0000 UTC
2021/11/29 22:54:00 2021-01-02 01:00:00 +0000 UTC
2021/11/29 22:54:00 2021-01-02 03:00:00 +0000 UTC
2021/11/29 22:54:00 0001-01-01 00:00:00 +0000 UTC
```

What's going on here? We're creating a struct (you don't have to, though) that contains our timeFactory.Now struct type. By default, this value is nil. When called on a nil-valued struct, the Get method will return the default time.Now() result. That's why the first Println returns a non-utc time.

However, we're overriding the nil value and replacing it with a custom generator and setting it to the year, 1999.

Then, using the helper methods provided by timeFactory, we generate a series of times that Get will return. The time starts at the beginning of 2021, then each subsequent call returns the last time, plus the duration. Once all the durations are exhausted, it returns the time.Time{} default, which is the zero-value of Time, Jan 1, of year 1.
