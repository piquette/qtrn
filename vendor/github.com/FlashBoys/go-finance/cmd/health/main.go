package main

import (
	"fmt"
	"time"

	"github.com/FlashBoys/go-finance"
)

// Checks the health of various endpoints through go-finance funcs.
func main() {

	/*
		Check snapshot quote.
	*/
	q, err := finance.GetQuote("TWTR")
	if err != nil {
		fmt.Print("Problem fetching quote snapshot:\n", err)
	} else {
		fmt.Print("Success fetching quote snapshot:\n", q)
	}

	/*
		Check multiple snapshot quotes.
	*/
	symbols := []string{"AAPL", "TWTR", "FB"}
	quotes, err := finance.GetQuotes(symbols)
	if err != nil {
		fmt.Print("\nProblem fetching snapshot quotes:\n", err)
	} else {
		fmt.Print("\nSuccess fetching snapshot quotes:\n", quotes)
	}

	/*
		Check history.
	*/
	start := finance.ParseDatetime("1/1/2017")
	end := finance.ParseDatetime("2/1/2017")
	bars, err := finance.GetHistory("TWTR", start, end, finance.Day)
	if err != nil {
		fmt.Print("\nProblem fetching quote history:\n", err)
	} else {
		fmt.Print("\nSuccess fetching quote history:\n", bars)
	}

	/*
		Check currency pair.
	*/
	pairquote, err := finance.GetCurrencyPairQuote(finance.USDJPY)
	if err != nil {
		fmt.Print("\nProblem fetching currency pair:\n", err)
	} else {
		fmt.Print("\nSuccess fetching currency pair:\n", pairquote)
	}

	/*
		Check events history.
	*/
	start = finance.ParseDatetime("1/1/2010")
	end = finance.NewDatetime(time.Now())
	events, err := finance.GetEventHistory("AAPL", start, end, finance.Splits)
	if err != nil {
		fmt.Print("\nProblem fetching event history:\n", err)
	} else {
		fmt.Print("\nSuccess fetching event history:\n", events)
	}

	/*
		Check symbols list.
	*/
	symbols, err = finance.GetUSEquitySymbols()
	if err != nil {
		fmt.Print("\nProblem fetching symbol list:\n", err)
	} else {
		fmt.Print("\nSuccess fetching symbol list:\n", symbols[:10])
	}

	/*
		Check options.
	*/
	c, err := finance.NewCycle("AAPL")
	if err != nil {
		fmt.Print("\nProblem fetching option cycle:\n", err)
	} else {
		fmt.Print("\nSuccess fetching option cycle:\n", c)
	}

	calls, puts, err := c.GetFrontMonth()
	if err != nil {
		fmt.Print("\nProblem fetching option chain:\n", err)
	} else {
		fmt.Print("\nSuccess fetching option chain:\n", calls[:2], "\n", puts[:2])
	}

}
