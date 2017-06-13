package finance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetHistory(t *testing.T) {

	s := startTestServer("history_fixture.csv")
	defer s.Close()
	HistoryURL = s.URL

	cs := startCookieServer("yahoo_appl.html", true)
	defer cs.Close()
	sessionURL = cs.URL

	bars, err := GetHistory("", Datetime{}, Datetime{}, Day)
	assert.Nil(t, err)

	// result should correspond to the requested symbol bar.
	assert.Equal(t, "", bars[4].Symbol)
}

func Test_GetEventHistory(t *testing.T) {

	s := startTestServer("events_fixture.csv")
	defer s.Close()
	HistoryURL = s.URL

	cs := startCookieServer("yahoo_appl.html", true)
	defer cs.Close()
	sessionURL = cs.URL

	events, err := GetEventHistory("", Datetime{}, Datetime{}, Dividends)
	assert.Nil(t, err)

	// result should be a dividend event.
	assert.Equal(t, "", events[4].Val.Ratio)
}
