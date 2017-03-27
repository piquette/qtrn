package qtrnmain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_toString(t *testing.T) {

	// result should be string of 1.
	assert.Equal(t, "1", toString(1))

	// result should be a string of 0
	assert.Equal(t, "0", toString(0))
}

func Test_toInt(t *testing.T) {

	// result should be 34.
	assert.Equal(t, 34, toInt("34"))

	// result should be 0.
	assert.Equal(t, 0, toInt("-"))
}
