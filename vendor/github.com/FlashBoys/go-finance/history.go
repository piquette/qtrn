package finance

import (
	"strconv"

	"github.com/shopspring/decimal"
)

const (
	// Day interval.
	Day = "d"
	// Week interval.
	Week = "w"
	// Month interval.
	Month = "m"

	// Dividend constant.
	Dividend = "DIVIDEND"
	// Split constant.
	Split = "SPLIT"
)

type (
	// Interval is the duration of the bars returned from the query.
	Interval string
	// Bar represents a single bar(candle) in time-series of quotes.
	Bar struct {
		Date     Datetime
		Open     decimal.Decimal
		High     decimal.Decimal
		Low      decimal.Decimal
		Close    decimal.Decimal
		Volume   int
		AdjClose decimal.Decimal
		Symbol   string `yfin:"-"`
	}
	// Event contains one historical event (either a split or a dividend).
	Event struct {
		EventType string
		Date      Datetime
		Val       Value
		Symbol    string `yfin:"-"`
	}
	// Value is an event object that contains either a div amt or a split ratio.
	Value struct {
		Dividend decimal.Decimal
		Ratio    string
	}
)

// GetHistory fetches a single symbol's quote history from Yahoo Finance.
func GetHistory(symbol string, start Datetime, end Datetime, interval Interval) (b []Bar, err error) {

	// time range:
	// start |- | | [bars..] | | -| end

	params := map[string]string{
		"s":      symbol,
		"a":      strconv.Itoa(start.Month),
		"b":      strconv.Itoa(start.Day),
		"c":      strconv.Itoa(start.Year),
		"d":      strconv.Itoa(end.Month),
		"e":      strconv.Itoa(end.Day),
		"f":      strconv.Itoa(end.Year),
		"g":      string(interval),
		"ignore": ".csv",
	}

	t, err := fetchCSV(buildURL(HistoryURL, params))
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
func GetEventHistory(symbol string, start Datetime, end Datetime) (e []Event, err error) {

	params := map[string]string{
		"s":      symbol,
		"a":      strconv.Itoa(start.Month),
		"b":      strconv.Itoa(start.Day),
		"c":      strconv.Itoa(start.Year),
		"d":      strconv.Itoa(end.Month),
		"e":      strconv.Itoa(end.Day),
		"f":      strconv.Itoa(end.Year),
		"g":      "v",
		"y":      "0",
		"ignore": ".csv",
	}

	t, err := fetchCSV(buildURL(EventURL, params))
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

		isEvent := (row[0] == Dividend || row[0] == Split)
		if isEvent {
			mapFields(row, c, &ne)
			ne.Symbol = symbol
			e = append(e, ne)
		}
	}

	return
}
