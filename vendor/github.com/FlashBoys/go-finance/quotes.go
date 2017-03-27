// Package finance

package finance

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

// Quote is the object that is returned for a quote inquiry.
type Quote struct {
	Symbol             string          `yfin:"s"`
	Name               string          `yfin:"n"`
	LastTradeTime      Datetime        `yfin:"t1"`
	LastTradeDate      Datetime        `yfin:"d1"`
	LastTradePrice     decimal.Decimal `yfin:"l1"`
	LastTradeSize      int             `yfin:"k3"`
	Ask                decimal.Decimal `yfin:"a"`
	AskSize            int             `yfin:"a5"`
	Bid                decimal.Decimal `yfin:"b"`
	BidSize            int             `yfin:"b6"`
	Volume             int             `yfin:"v"`
	ChangeNominal      decimal.Decimal `yfin:"c1"`
	ChangePercent      decimal.Decimal `yfin:"p2"`
	Open               decimal.Decimal `yfin:"o"`
	PreviousClose      decimal.Decimal `yfin:"p"`
	Exchange           string          `yfin:"x"`
	DayLow             decimal.Decimal `yfin:"g"`
	DayHigh            decimal.Decimal `yfin:"h"`
	FiftyTwoWeekLow    decimal.Decimal `yfin:"j"`
	FiftyTwoWeekHigh   decimal.Decimal `yfin:"k"`
	Currency           string          `yfin:"c4"`
	MarketCap          string          `yfin:"j1"`
	FiftyDayMA         decimal.Decimal `yfin:"m3"`
	TwoHundredDayMA    decimal.Decimal `yfin:"m4"`
	AvgDailyVolume     int             `yfin:"a2"`
	FiftyTwoWeekTarget decimal.Decimal `yfin:"t8"`
	ShortRatio         decimal.Decimal `yfin:"s7"`
	BookValue          decimal.Decimal `yfin:"b4"`
	EBITDA             string          `yfin:"j4"`
	PriceSales         decimal.Decimal `yfin:"p5"`
	PriceBook          decimal.Decimal `yfin:"p6"`
	PERatio            decimal.Decimal `yfin:"r"`
	PEGRatio           decimal.Decimal `yfin:"r5"`
	DivYield           decimal.Decimal `yfin:"y"`
	DivPerShare        decimal.Decimal `yfin:"d"`
	DivExDate          Datetime        `yfin:"q"`
	DivPayDate         Datetime        `yfin:"r1"`
	EPS                decimal.Decimal `yfin:"e"`
	EPSEstCurrentYear  decimal.Decimal `yfin:"e7"`
	EPSEstNextYear     decimal.Decimal `yfin:"e8"`
	EPSEstNextQuarter  decimal.Decimal `yfin:"e9"`
}

// GetQuote fetches a single symbol's quote from Yahoo Finance.
func GetQuote(symbol string) (q Quote, err error) {

	f, c := structFields(q)
	params := map[string]string{
		"s": symbol,
		"f": f,
		"e": ".csv",
	}

	t, err := fetchCSV(buildURL(QuoteURL, params))
	if err != nil {
		return
	}

	mapFields(t[0], c, &q)

	return
}

// GetQuotes fetches multiple symbol's quotes from Yahoo Finance.
func GetQuotes(symbols []string) (q []Quote, err error) {

	var nq Quote
	f, c := structFields(nq)
	params := map[string]string{
		"s": strings.Join(symbols[:], ","),
		"f": f,
		"e": ".csv",
	}

	t, err := fetchCSV(buildURL(QuoteURL, params))
	if err != nil {
		return
	}

	if len(t) == 0 {
		return nil, fmt.Errorf("symbol does not exist")
	}

	for _, row := range t {
		mapFields(row, c, &nq)
		q = append(q, nq)
	}

	return
}
