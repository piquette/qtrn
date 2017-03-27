package finance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetHistory(t *testing.T) {

	s := startTestServer("history_fixture.csv")
	defer s.Close()
	HistoryURL = s.URL

	bars, err := GetHistory("TWTR", Datetime{}, Datetime{}, Day)
	assert.Nil(t, err)

	// result should be a TWTR bar.
	assert.Equal(t, "TWTR", bars[4].Symbol)
}

func Test_GetEventHistory(t *testing.T) {

	s := startTestServer("events_fixture.csv")
	defer s.Close()
	EventURL = s.URL

	events, err := GetEventHistory("TWTR", Datetime{}, Datetime{})
	assert.Nil(t, err)

	// result should be a TWTR event.
	assert.Equal(t, "TWTR", events[4].Symbol)
}
