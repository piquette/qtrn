package finance

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

var (
	firstRegex  = regexp.MustCompile(`(\w+:)(\d+\.?\d*)`)
	secondRegex = regexp.MustCompile(`(\w+):`)
)

var (
	// OptionsURL option chains
	OptionsURL = "http://www.google.com/finance/option_chain?"
	// HistoryURL quote history
	HistoryURL = "https://query1.finance.yahoo.com/v7/finance/download/"
	// SymbolsURL symbols list
	SymbolsURL = "http://www.batstrading.com/market_data/symbol_listing/csv/"
	// QuoteURL stock quotes
	QuoteURL = "http://download.finance.yahoo.com/d/quotes.csv"
	// sessionURL cookie parsing
	sessionURL = "https://finance.yahoo.com/quote/AAPL/history"
)

type optionsResponse struct {
	Expiry      json.RawMessage     `json:"expiry"`
	Expirations []Datetime          `json:"expirations"`
	Underlying  json.RawMessage     `json:"underlying_id"`
	Price       string              `json:"underlying_price"`
	Calls       []map[string]string `json:"calls,omitempty"`
	Puts        []map[string]string `json:"puts,omitempty"`
}

func fetchCSV(url string, cookie *http.Cookie) (table [][]string, err error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	if cookie != nil {
		req.AddCookie(cookie)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	r := csv.NewReader(resp.Body)
	r.FieldsPerRecord = -1
	table, err = r.ReadAll()
	return
}

// buildURL takes a base URL and parameters returns the full URL.
func buildURL(base string, params map[string]string) string {

	url, _ := url.ParseRequestURI(base)
	q := url.Query()

	for k, v := range params {
		q.Set(k, v)
	}
	url.RawQuery = q.Encode()

	return url.String()
}

// getsession retrieves a session cookie and crumb to validate the yhoo request.
func getsession() (*http.Cookie, string, error) {

	req, err := http.NewRequest("GET", sessionURL, nil)
	if err != nil {
		return nil, "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	// Get cookies.
	cookies := resp.Cookies()
	if len(cookies) <= 0 {
		return nil, "", errors.New("cookie unavailable")
	}

	// Get crumb.
	b, err := ioutil.ReadAll(resp.Body)
	match := rp.FindStringSubmatch(string(b))
	if len(match) <= 1 {
		return nil, "", errors.New("crumb unavailable")
	}

	return cookies[0], match[1], nil
}

// buildURL takes a base URL and parameters returns the full URL.
func buildOptionsURL(base string, symbol string, d Datetime) string {
	return buildURL(base, map[string]string{
		"q":      symbol,
		"expd":   strconv.Itoa(d.Day),
		"expm":   strconv.Itoa(d.Month),
		"expy":   strconv.Itoa(d.Year),
		"output": "json",
	})
}

// fetchOptions retrieves options data.
func fetchOptions(url string) (or *optionsResponse, err error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result := string(contents)
	j := []byte(secondRegex.ReplaceAllString(firstRegex.ReplaceAllString(result, "$1\"$2\""), "\"$1\":"))

	err = json.Unmarshal(j, &or)
	if err != nil {
		return nil, fmt.Errorf("options format error:  (error was: %s)\n", err.Error())
	}

	return
}
