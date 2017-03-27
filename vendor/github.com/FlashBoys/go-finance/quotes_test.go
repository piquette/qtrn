package finance

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getFixtureAsTable(filename string) [][]string {

	path := "./fixtures/" + filename
	f, err := os.Open(path)
	r := csv.NewReader(bufio.NewReader(f))
	r.FieldsPerRecord = -1
	table, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	return table
}

func getFixtureAsString(filename string) string {

	path := "./fixtures/" + filename
	jsonData, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(jsonData)
}

func startTestServer(fixtureFile string) *httptest.Server {

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if fixtureFile != "" {
			fmt.Fprintln(w, getFixtureAsString(fixtureFile))
		} else {
			fmt.Fprintln(w, "")
		}
	}))
}

func Test_GetQuote(t *testing.T) {

	s := startTestServer("quote_fixture.csv")
	defer s.Close()
	QuoteURL = s.URL

	q, err := GetQuote("AAPL")
	assert.Nil(t, err)

	// result should be a an Apple quote.
	assert.Equal(t, "AAPL", q.Symbol)

	s = startTestServer("quote_not_fixture.csv")
	defer s.Close()
	QuoteURL = s.URL

	_, err = GetQuote("AAPL")
	assert.NotNil(t, err)

}

func Test_GetQuotes(t *testing.T) {

	s := startTestServer("quotes_fixture.csv")
	QuoteURL = s.URL

	quotes, err := GetQuotes([]string{"AAPL", "TWTR"})
	assert.Nil(t, err)

	// result should be a an Apple quote and a Twitter quote.
	assert.Equal(t, "AAPL", quotes[0].Symbol)
	assert.Equal(t, "TWTR", quotes[1].Symbol)
	s.Close()

	s = startTestServer("")
	QuoteURL = s.URL

	_, err = GetQuotes([]string{"AAPL", "TWTR"})
	assert.NotNil(t, err)
	s.Close()

	s = startTestServer("quote_not_fixture.csv")
	defer s.Close()
	QuoteURL = s.URL

	_, err = GetQuotes([]string{"AAPL", "TWTR"})
	assert.NotNil(t, err)

}
