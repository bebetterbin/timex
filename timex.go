package timex

import "time"

// StartTimeOfDay return the 00:00 time of the t in loc.
func StartTimeOfDay(t time.Time, loc *time.Location) time.Time {
	tInLoc := t.In(loc)
	return time.Date(tInLoc.Year(), tInLoc.Month(), tInLoc.Day(), 0, 0, 0, 0, loc)
}

// StartTimeOfMonth return the 00:00 time of the first day of the month of the t in loc.
func StartTimeOfMonth(t time.Time, loc *time.Location) time.Time {
	tInLoc := t.In(loc)
	return time.Date(tInLoc.Year(), tInLoc.Month(), 1, 0, 0, 0, 0, loc)
}

// StartTimeOfQuarter return the 00:00 time of the first day of the quarter of the t in loc.
func StartTimeOfQuarter(t time.Time, loc *time.Location) time.Time {
	tInLoc := t.In(loc)
	month := tInLoc.Month()
	return time.Date(tInLoc.Year(), ((month-1)/3)*3+1, 1, 0, 0, 0, 0, loc)
}
