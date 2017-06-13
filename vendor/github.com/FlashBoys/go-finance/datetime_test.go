package finance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UnixTime(t *testing.T) {

	// Given have have a time,
	start := ParseDatetime("1/1/2017")

	// When we convert it to a unix timestamp,
	timestamp := start.unixTime()

	// Then it should equal a string of the number of secs since Jan 1, 1970-
	assert.Equal(t, "1483228800", timestamp)
}
