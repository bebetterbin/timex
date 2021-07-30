package timex

import (
	"time"
)

// CST represents the time zone of beijing.
var CST = time.FixedZone("CST", 8*60*60)

// BeijingNow return the current beijing time.
func BeijingNow() time.Time {
	return time.Now().In(CST)
}
