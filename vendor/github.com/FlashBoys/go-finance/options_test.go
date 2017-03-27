package finance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewCycle(t *testing.T) {

	s := startTestServer("options_fixture.txt")
	defer s.Close()
	OptionsURL = s.URL

	c, err := NewCycle("TWTR")
	assert.Nil(t, err)

	// result should be a TWTR cycle
	assert.Equal(t, "TWTR", c.Symbol)

	// result should have expirations.
	assert.NotEmpty(t, c.Expirations)
}

func Test_GetChainForExpiration(t *testing.T) {

	s := startTestServer("options_fixture.txt")
	defer s.Close()
	OptionsURL = s.URL

	c, err := NewCycle("TWTR")
	assert.Nil(t, err)

	calls, puts, err := c.GetChainForExpiration(Datetime{})

	// result chain should be empty for datetime zero-value.
	assert.NotNil(t, err)
	assert.Empty(t, puts)
	assert.Empty(t, calls)

	calls, puts, err = c.GetChainForExpiration(c.Expirations[0])

	// result chain should not be empty.
	assert.Nil(t, err)
	assert.NotEmpty(t, puts)
	assert.NotEmpty(t, calls)
}

func Test_GetFrontMonth(t *testing.T) {

	s := startTestServer("options_fixture.txt")
	defer s.Close()
	OptionsURL = s.URL

	c, err := NewCycle("TWTR")
	assert.Nil(t, err)

	calls, puts, err := c.GetFrontMonth()

	// result chain should not be empty.
	assert.Nil(t, err)
	assert.NotEmpty(t, puts)
	assert.NotEmpty(t, calls)
}

func Test_GetCallsForExpiration(t *testing.T) {

	s := startTestServer("options_fixture.txt")
	defer s.Close()
	OptionsURL = s.URL

	c, err := NewCycle("TWTR")
	assert.Nil(t, err)

	calls, err := c.GetCallsForExpiration(Datetime{})

	// result calls should be empty for datetime zero-value.
	assert.NotNil(t, err)
	assert.Empty(t, calls)

	calls, err = c.GetCallsForExpiration(c.Expirations[0])

	// result calls should not be empty.
	assert.Nil(t, err)
	assert.NotEmpty(t, calls)
}

func Test_GetPutsForExpiration(t *testing.T) {

	s := startTestServer("options_fixture.txt")
	defer s.Close()
	OptionsURL = s.URL

	c, err := NewCycle("TWTR")
	assert.Nil(t, err)

	puts, err := c.GetPutsForExpiration(Datetime{})

	// result puts should be empty for datetime zero-value.
	assert.NotNil(t, err)
	assert.Empty(t, puts)

	puts, err = c.GetPutsForExpiration(c.Expirations[0])

	// result puts should not be empty.
	assert.Nil(t, err)
	assert.NotEmpty(t, puts)
}
