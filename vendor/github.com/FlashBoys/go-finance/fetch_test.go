package finance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FetchCSV(t *testing.T) {

	// Given that we want to download a csv
	ts := startTestServer("request_csv_fixture.csv")
	defer ts.Close()

	// When we request the csv,
	table, err := fetchCSV(ts.URL)
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

func Test_Fetch(t *testing.T) {
	// Given that we want to download options data
	// ts := startTestServer("request_fixture.txt")
	// defer ts.Close()
	//
	// // When we request the malformed json text,
	// response, err := fetchOptions(ts.URL)
	// assert.Nil(t, err)

	// Then the returned string should be-
	//assert.Equal(t, "{\"test\":{\"foo\":bar}}\n\n", string(response))

}
