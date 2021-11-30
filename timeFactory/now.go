package timeFactory

import "time"

// Now is a type used in structs. The default is to call time.Now(), but you can override this behavior by replacing
// this instance value of the NowGetter with any method you wish that returns a NowGetter.
// Do not use with pointers of this type, it will do strange things
type Now func() time.Time

// Get will return the current time if not set, or will call the NowGetter if set
// provides a nice, safe default while allowing developers to override the current time.
func (n Now) Get() time.Time {
	if n != nil {
		return n()
	}
	return time.Now()
}
