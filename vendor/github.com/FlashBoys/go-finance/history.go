package finance

import (
	"regexp"

	"github.com/shopspring/decimal"
)

var rp = regexp.MustCompile("\"CrumbStore\":{\"crumb\":\"([^\"]+)\"}")

const (
	// Day interval.
	Day = "1d"
	// Week interval.
	Week = "1wk"
	// Month interval.
	Month = "1mo"
	// Dividends event type.
	Dividends = "div"
	// Splits event type.
	Splits = "split"
)

type (
	// Interval is the duration of the bars returned from the query.
	Interval string
	// EventType is the type of history event, either divs or splits.
	EventType string
	// Bar represents a single bar(candle) in time-series of quotes.
	Bar struct {
		Date     Datetime
		Open     decimal.Decimal
		High     decimal.Decimal
		Low      decimal.Decimal
		Close    decimal.Decimal
		AdjClose decimal.Decimal
		Volume   int
		Symbol   string `yfin:"-"`
	}

	// Event contains one historical event (either a split or a dividend).
	Event struct {
		Date   Datetime
		Val    Value
		Type   EventType `yfin:"-"`
		Symbol string    `yfin:"-"`
	}
	// Value is an event object that contains either a div amt or a split ratio.
	Value struct {
		Dividend decimal.Decimal
		Ratio    string
	}
)

// GetHistory fetches a single symbol's quote history from Yahoo Finance.
func GetHistory(symbol string, start Datetime, end Datetime, interval Interval) (b []Bar, err error) {

	cookie, crumb, err := getsession()
	if err != nil {
		return
	}

	params := map[string]string{
		"period1":  start.unixTime(),
		"period2":  end.unixTime(),
		"interval": string(interval),
		"events":   "history",
		"crumb":    crumb,
	}

	t, err := fetchCSV(buildURL(HistoryURL+symbol, params), cookie)
	if err != nil {
		return
	}

	var nb Bar
	_, c := structFields(nb)
	for i, row := range t {

		// Skip the header.
		if i == 0 {
			continue
		}

		mapFields(row, c, &nb)
		nb.Symbol = symbol
		b = append(b, nb)
	}

	return
}

// GetEventHistory fetches a single symbol's dividend and split history from Yahoo Finance.
func GetEventHistory(symbol string, start Datetime, end Datetime, eventType EventType) (e []Event, err error) {

	cookie, crumb, err := getsession()
	if err != nil {
		return
	}

	params := map[string]string{
		"period1":  start.unixTime(),
		"period2":  end.unixTime(),
		"interval": "1d",
		"events":   string(eventType),
		"crumb":    crumb,
	}

	t, err := fetchCSV(buildURL(HistoryURL+symbol, params), cookie)
	if err != nil {
		return
	}

	var ne Event
	_, c := structFields(ne)
	for i, row := range t {

		// Skip the header.
		if i == 0 {
			continue
		}

		mapFields(row, c, &ne)
		ne.Symbol = symbol
		ne.Type = eventType
		e = append(e, ne)
	}

	return
}
