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

func TestStartTimeOfMonth(t *testing.T) {
	assert := assert.New(t)

	time, err := time.Parse(time.RFC3339, "2021-02-03T04:05:06+08:00")
	assert.Nil(err)

	startTimeOfMonth := StartTimeOfMonth(time, CST)
	assert.Equal(2021, startTimeOfMonth.Year())
	assert.EqualValues(2, startTimeOfMonth.Month())
	assert.Equal(1, startTimeOfMonth.Day())
	assert.Zero(startTimeOfMonth.Hour())
	assert.Zero(startTimeOfMonth.Minute())
	assert.Zero(startTimeOfMonth.Second())
	assert.Zero(startTimeOfMonth.Nanosecond())
}

func TestStartTimeOfQuarter(t *testing.T) {
	assert := assert.New(t)

	time, err := time.Parse(time.RFC3339, "2021-02-03T04:05:06+08:00")
	assert.Nil(err)

	startTimeOfQuarter := StartTimeOfQuarter(time, CST)
	assert.Equal(2021, startTimeOfQuarter.Year())
	assert.EqualValues(1, startTimeOfQuarter.Month())
	assert.Equal(1, startTimeOfQuarter.Day())
	assert.Zero(startTimeOfQuarter.Hour())
	assert.Zero(startTimeOfQuarter.Minute())
	assert.Zero(startTimeOfQuarter.Second())
	assert.Zero(startTimeOfQuarter.Nanosecond())
}
