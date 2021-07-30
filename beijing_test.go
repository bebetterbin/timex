package timex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCST(t *testing.T) {
	assert.Equal(t, "CST", CST.String())
}

func TestBeijingNow(t *testing.T) {
	_, offset := BeijingNow().Zone()
	assert.Equal(t, 8*60*60, offset)
}
