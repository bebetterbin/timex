package timex

import "time"

// StartTimeOfDay return the 00:00 time of the t in loc.
func StartTimeOfDay(t time.Time, loc *time.Location) time.Time {
	tInLoc := t.In(loc)
	return time.Date(tInLoc.Year(), tInLoc.Month(), tInLoc.Day(), 0, 0, 0, 0, loc)
}
