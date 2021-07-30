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

func TestStartTimeOfWeek(t *testing.T) {
	assert := assert.New(t)

	time, err := time.Parse(time.RFC3339, "2021-02-03T04:05:06+08:00")
	assert.Nil(err)

	startTimeOfWeek := StartTimeOfWeek(time, CST)
	assert.Equal(2021, startTimeOfWeek.Year())
	assert.EqualValues(2, startTimeOfWeek.Month())
	assert.Equal(1, startTimeOfWeek.Day())
	assert.Zero(startTimeOfWeek.Hour())
	assert.Zero(startTimeOfWeek.Minute())
	assert.Zero(startTimeOfWeek.Second())
	assert.Zero(startTimeOfWeek.Nanosecond())
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

func TestStartTimeOfYear(t *testing.T) {
	assert := assert.New(t)

	time, err := time.Parse(time.RFC3339, "2021-02-03T04:05:06+08:00")
	assert.Nil(err)

	startTimeOfYear := StartTimeOfYear(time, CST)
	assert.Equal(2021, startTimeOfYear.Year())
	assert.EqualValues(1, startTimeOfYear.Month())
	assert.Equal(1, startTimeOfYear.Day())
	assert.Zero(startTimeOfYear.Hour())
	assert.Zero(startTimeOfYear.Minute())
	assert.Zero(startTimeOfYear.Second())
	assert.Zero(startTimeOfYear.Nanosecond())
}
