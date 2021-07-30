package timex

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStartTimeOfDay(t *testing.T) {
	assert := assert.New(t)

	time, err := time.Parse(time.RFC3339, "2021-02-03T04:05:06+08:00")
	assert.Nil(err)

	startTimeOfDay := StartTimeOfDay(time, CST)
	assert.Equal(2021, startTimeOfDay.Year())
	assert.EqualValues(2, startTimeOfDay.Month())
	assert.Equal(3, startTimeOfDay.Day())
	assert.Zero(startTimeOfDay.Hour())
	assert.Zero(startTimeOfDay.Minute())
	assert.Zero(startTimeOfDay.Second())
	assert.Zero(startTimeOfDay.Nanosecond())
}
