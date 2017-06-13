package finance

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FetchCSV(t *testing.T) {

	// Given that we want to download a csv
	ts := startTestServer("request_csv_fixture.csv")
	defer ts.Close()

	// When we request the csv,
	table, err := fetchCSV(ts.URL, &http.Cookie{})
	assert.Nil(t, err)

	// Then the returned table should have 1 row.
	assert.Len(t, table, 1)
	// And the first cell should be
	assert.Equal(t, "foo", table[0][0])
	// Then the second cell should be
	assert.Equal(t, "bar", table[0][1])
}

func Test_BuildURL(t *testing.T) {

	//Given we have a base url and a set of query params,
	base := "http://example.com/d/quotes.csv"
	params := map[string]string{
		"s": "AAPL",
	}

	// When we convert it to a url,
	url := buildURL(base, params)

	// Then the url should equal-
	assert.Equal(t, "http://example.com/d/quotes.csv?s=AAPL", url)
}

func Test_GetSession(t *testing.T) {

	// Given we need a valid yhoo session,
	cs := startCookieServer("yahoo_appl.html", true)
	sessionURL = cs.URL

	// When we make a request,
	cookie, crumb, err := getsession()

	// We should get valid params back.
	assert.Nil(t, err)
	assert.NotNil(t, cookie)
	assert.Equal(t, "j\\u002FlphNGEHaA", crumb)
	cs.Close()

	// Test bad crumb handling-
	cs = startCookieServer("", true)
	sessionURL = cs.URL
	cookie, crumb, err = getsession()
	assert.Nil(t, cookie)
	assert.Equal(t, "crumb unavailable", err.Error())
	cs.Close()

	// Test bad cookie handling-
	cs = startCookieServer("yahoo_appl.html", false)
	sessionURL = cs.URL
	defer cs.Close()

	cookie, crumb, err = getsession()
	assert.Equal(t, "cookie unavailable", err.Error())

}

func Test_BuildOptionsURL(t *testing.T) {

	// Given we have a set of params and a base url,
	baseURL := "http://example.org"
	sym := "TWTR"
	dt := Datetime{Month: 5, Day: 30, Year: 2017}

	// When we construct the options url,
	optionsURL := buildOptionsURL(baseURL, sym, dt)

	// It should equal-
	assert.Equal(t, "http://example.org?expd=30&expm=5&expy=2017&output=json&q=TWTR", optionsURL)
}

func Test_FetchOptions(t *testing.T) {
	// Given that we want to download options data
	ts := startTestServer("options_fixture.txt")

	// When we request the malformed json text,
	or, err := fetchOptions(ts.URL)
	assert.Nil(t, err)

	// Then the returned response should be-
	assert.Equal(t, "110.43", or.Price)
	ts.Close()

	// Given that we want to download options data
	ts = startTestServer("request_fixture.txt")
	defer ts.Close()

	// When we request the malformed json text,
	or, err = fetchOptions(ts.URL)
	assert.NotNil(t, err)
}
