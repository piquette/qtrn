// Copyright © 2017 Michael Ackley <ackleymi@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	finance "github.com/FlashBoys/go-finance"
	"github.com/spf13/cobra"
)

const (
	writeUsage            = "write [subcommand]"
	writeShortDesc        = "Writes a csv of stock market data"
	writeLongDesc         = "Writes a csv of stock market data into the current directory using a subcommand `quote` for quotes or `history` for historical prices"
	writeQuoteShortDesc   = "Writes a csv of a stock quote"
	writeQuoteLongDesc    = "Writes a csv of a stock quote and can accomodate multiple symbols as arguments"
	writeHistoryShortDesc = "Writes a csv of a historical data"
	writeHistoryLongDesc  = "Writes a csv of a historical data, can only accept one symbol at a time"
)

var (
	// write command.
	writeCmd = &cobra.Command{
		Use:     writeUsage,
		Short:   writeShortDesc,
		Long:    writeLongDesc,
		Aliases: []string{"w"},
		Example: "$ qtrn write -h quote -f AAPL GOOG FB AMZN",
		Run: func(cmd *cobra.Command, args []string) {
			// Stub.
			fmt.Printf("\nSubcommand not specified, use either ( quote | history )\n\n")
		},
	}
	writeQuoteCmd = &cobra.Command{
		Use:     "quote",
		Short:   writeQuoteShortDesc,
		Long:    writeQuoteLongDesc,
		Aliases: []string{"q"},
		Example: "$ qtrn write quote AAPL",
		Run:     writeQuoteFunc,
	}
	writeHistoryCmd = &cobra.Command{
		Use:     "history",
		Short:   writeHistoryShortDesc,
		Long:    writeHistoryLongDesc,
		Aliases: []string{"h"},
		Example: "$ qtrn write history",
		Run:     writeHistoryFunc,
	}
	// flagRemoveHeader set flag to specify whether to remove the header in the file.
	flagRemoveHeader bool
	// flagFullOutput set flag to write a more informative quote.
	flagWriteFullOutput bool
	// flagStartTime set flag to specify the start time of the csv frame.
	flagWriteStartTime string
	// flagEndTime set flag to specify the end time of the csv frame.
	flagWriteEndTime string
	// flagInterval set flag to specify time interval of each OHLC point.
	flagWriteInterval string
	//historyHeader csv header for history results.
	historyHeader = []string{"date", "open", "high", "low", "close", "adj-close", "volume", "symbol"}
	//quoteHeader csv header for quote results.
	quoteHeader = []string{"symbol", "date", "last", "change", "change-percent", "volume", "bid", "bid-size", "ask", "ask-size", "open", "prev-close", "name"}
	//quoteFullHeader csv header for quote results.
	quoteFullHeader = []string{"symbol", "date", "name", "last", "change", "change-percent", "volume", "bid", "bid-size",
		"ask", "ask-size", "open", "prev-close", "exchange", "day-high", "day-low", "52wk-high", "52wk-low", "mkt-cap",
		"50d-ma", "200d-ma", "avg-volume", "eps", "pe-ratio", "peg-ratio", "price-sales", "price-book", "div-per-share",
		"div-yield", "eps-est-next-qtr", "eps-est-year", "short-ratio", "book-value", "ebitda"}
)

func init() {
	writeCmd.AddCommand(writeQuoteCmd)
	writeCmd.AddCommand(writeHistoryCmd)
	writeCmd.Flags().BoolVarP(&flagRemoveHeader, "remove", "r", false, "Set `--remove` or `-r` to remove the header in the csv. Default is FALSE.")
	writeQuoteCmd.Flags().BoolVarP(&flagWriteFullOutput, "full", "f", false, "Set `--full` or `-f` to write a more informative quote for each symbol")
	writeHistoryCmd.Flags().StringVarP(&flagWriteStartTime, "start", "s", "", "Set a date (formatted YYYY-MM-DD) using `--start` or `-s` to specify the start of the csv time frame")
	writeHistoryCmd.Flags().StringVarP(&flagWriteEndTime, "end", "e", "", "Set a date (formatted YYYY-MM-DD) using `--start` or `-s` to specify the start of the csv time frame")
	writeHistoryCmd.Flags().StringVarP(&flagWriteInterval, "interval", "i", finance.Day, "Set an interval ( 1d | 1wk | 1mo ) using `--interval` or `-i` to specify the time interval of each OHLC point")
}

// writeQuoteFunc implements the quote write command.
func writeQuoteFunc(cmd *cobra.Command, args []string) {

	symbols := args[:]
	if len(symbols) == 0 {
		fmt.Println("No symbols provided.")
		return
	}

	quotes, err := finance.GetQuotes(symbols)
	if err != nil {
		fmt.Println("Error fetching data, please try again")
		return
	}

	data := formatQuoteData(quotes)
	err = writeData(headerForQuote(), combine(symbols, "+"), "quote", data)
	if err != nil {
		fmt.Println("Error writing data")
	}
}

// writeHistoryFunc implements the history write command.
func writeHistoryFunc(cmd *cobra.Command, args []string) {

	symbol := args[0]
	if symbol == "" {
		fmt.Println("No symbol provided.")
		return
	}

	var start finance.Datetime
	var end finance.Datetime
	if flagWriteStartTime == "" {
		start = finance.ParseDatetime(fmt.Sprintf("%v-01-01", time.Now().Year()))
	} else {
		start = finance.ParseDatetime(flagWriteStartTime)
	}
	if flagWriteEndTime == "" {
		t := time.Now()
		end = finance.ParseDatetime(fmt.Sprintf("%d-%02d-%02d", t.Year(), int(t.Month()), t.Day()))
	} else {
		end = finance.ParseDatetime(flagWriteEndTime)
	}

	bars, err := finance.GetHistory(symbol, start, end, finance.Interval(flagWriteInterval))
	if err != nil {
		fmt.Println("Error fetching data, please try again")
		return
	}

	data := formatHistoricalData(bars)
	err = writeData(historyHeader, symbol, "history", data)
	if err != nil {
		fmt.Println("Error writing data")
	}
}

// writeData writes the downloaded and formatted data to a csv.
func writeData(header []string, prefix string, cmd string, data [][]string) error {

	t := time.Now()
	now := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d-00:00",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	fileTitle := fmt.Sprintf("%v_%v_%v.csv", prefix, cmd, now)
	file, err := os.Create(fileTitle)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if !flagRemoveHeader {
		err = writer.Write(header)
		if err != nil {
			return err
		}
	}

	for _, value := range data {
		err = writer.Write(value)
		if err != nil {
			return err
		}
	}
	fmt.Println(fileTitle)
	return nil
}

// formatHistoricalData formats a slice of historical prices into a writeable format.
func formatHistoricalData(bars []finance.Bar) (data [][]string) {
	for _, b := range bars {
		date := fmt.Sprintf("%v-%v-%v", b.Date.Year, b.Date.Month, b.Date.Day)
		point := []string{
			date,
			b.Open.StringFixed(2),
			b.High.StringFixed(2),
			b.Low.StringFixed(2),
			b.Close.StringFixed(2),
			b.AdjClose.StringFixed(2),
			toString(b.Volume),
			b.Symbol,
		}
		data = append(data, point)
	}
	return
}

// formatQuoteData formats a slice of quotes into a writeable format.
func formatQuoteData(quotes []finance.Quote) (data [][]string) {
	for _, q := range quotes {

		var formattedQuote []string
		date := fmt.Sprintf("%v-%v-%v", q.LastTradeDate.Year, q.LastTradeDate.Month, q.LastTradeDate.Day)
		if flagWriteFullOutput {
			formattedQuote = []string{
				q.Symbol,
				date,
				q.Name,
				q.LastTradePrice.StringFixed(2),
				q.ChangeNominal.StringFixed(2),
				q.ChangePercent.StringFixed(2),
				toString(q.Volume),
				q.Bid.StringFixed(2),
				toString(q.BidSize),
				q.Ask.StringFixed(2),
				toString(q.AskSize),
				q.Open.StringFixed(2),
				q.PreviousClose.StringFixed(2),
				q.Exchange,
				q.DayHigh.StringFixed(2),
				q.DayLow.StringFixed(2),
				q.FiftyTwoWeekHigh.StringFixed(2),
				q.FiftyTwoWeekLow.StringFixed(2),
				q.MarketCap,
				q.FiftyDayMA.StringFixed(2),
				q.TwoHundredDayMA.StringFixed(2),
				toString(q.AvgDailyVolume),
				q.EPS.StringFixed(2),
				q.PERatio.StringFixed(2),
				q.PEGRatio.StringFixed(2),
				q.PriceSales.StringFixed(2),
				q.PriceBook.StringFixed(2),
				q.DivPerShare.StringFixed(2),
				q.DivYield.StringFixed(2),
				q.EPSEstNextQuarter.StringFixed(2),
				q.EPSEstCurrentYear.StringFixed(2),
				q.ShortRatio.String(),
				q.BookValue.String(),
				q.EBITDA,
			}
		} else {
			formattedQuote = []string{
				q.Symbol,
				date,
				q.LastTradePrice.StringFixed(2),
				q.ChangeNominal.StringFixed(2),
				q.ChangePercent.StringFixed(2),
				toString(q.Volume),
				q.Bid.StringFixed(2),
				toString(q.BidSize),
				q.Ask.StringFixed(2),
				toString(q.AskSize),
				q.Open.StringFixed(2),
				q.PreviousClose.StringFixed(2),
				q.Name,
			}
		}
		data = append(data, formattedQuote)
	}

	return
}

// headerForQuote determines which header to write.
func headerForQuote() []string {
	if flagWriteFullOutput {
		return quoteFullHeader
	}
	return quoteHeader
}
