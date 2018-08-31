// Copyright © 2018 Piquette Capital, LLC
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

package write

import (
	"fmt"
	"strings"
	"time"

	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
	"github.com/piquette/qtrn/utils"

	"github.com/piquette/finance-go/quote"

	"github.com/spf13/cobra"
)

const (
	usage        = "write"
	short        = "Writes a csv of stock market data"
	long         = "Writes a csv of stock market data into the current directory."
	quoteShort   = "Writes a csv of stock quotes"
	quoteLong    = "Writes a csv of stock quotes and can accept multiple symbols as arguments"
	historyShort = "Writes a csv of historical data"
	historyLong  = "Writes a csv of historical data, can only accept one symbol at a time"
)

var (
	// Cmd is the write command.
	Cmd = &cobra.Command{
		Use:     usage,
		Short:   short,
		Long:    long,
		Aliases: []string{"w"},
		Example: "qtrn write -r quote -f AAPL GOOG FB AMZN",
	}
	// quote subcommand
	quoteCmd = &cobra.Command{
		Use:     "quote",
		Short:   quoteShort,
		Long:    quoteLong,
		Example: "qtrn write quote AAPL",
		RunE:    executeQuote,
	}
	// history subcommand
	historyCmd = &cobra.Command{
		Use:     "history",
		Short:   historyShort,
		Long:    historyLong,
		Example: "qtrn write history [flags]",
		RunE:    executeHistory,
	}
	// removeHeaderF set flag to specify whether to remove the header in the file.
	removeHeaderF bool
	// startF set flag to specify the start time of the csv frame.
	startF string
	// endF set flag to specify the end time of the csv frame.
	endF string
	// aggregationF set flag to specify time interval of each OHLC point.
	aggregationF string
	//historyFields csv header for history results.
	historyFields = []string{"date", "open", "high", "low", "close", "adj-close", "volume", "symbol"}
	//quoteFields csv header for quote results.
	quoteFields = []string{"symbol", "date", "last", "change", "change-percent", "volume", "bid", "bid-size", "ask", "ask-size", "open", "prev-close", "name"}
)

func init() {
	Cmd.AddCommand(quoteCmd)
	Cmd.AddCommand(historyCmd)
	Cmd.Flags().BoolVarP(&removeHeaderF, "remove", "r", false, "remove the header in the csv. default is false.")
	historyCmd.Flags().StringVarP(&startF, "start", "s", "", "set a date (formatted yyyy-mm-dd) to specify the start of the historical time frame")
	historyCmd.Flags().StringVarP(&endF, "end", "e", "", "set a date (formatted yyyy-mm-dd) to specify the end of the historical time frame")
	historyCmd.Flags().StringVarP(&aggregationF, "agg", "a", "1d", "set a candle aggregation interval ( 1d | 5d | 1mo | 1y )")
}

// executeQuote implements the quote data writer.
func executeQuote(cmd *cobra.Command, args []string) error {
	// check symbols.
	symbols := args
	if len(symbols) == 0 {
		return fmt.Errorf("no symbols provided")
	}

	// get quote data.
	i := quote.List(symbols)

	// format as strings.
	quotes, err := format(i)
	if err != nil {
		return err
	}

	// write csv.
	err = write(quoteFields, strings.Join(symbols, "_"), "quote", quotes)
	if err != nil {
		return fmt.Errorf("error writing data")
	}

	return nil
}

// executeHistory implements the history data writer.
func executeHistory(cmd *cobra.Command, args []string) error {
	// check symbols.
	symbols := args
	if len(symbols) == 0 {
		return fmt.Errorf("no symbols provided")
	}

	// format params.
	// start.
	var start *datetime.Datetime
	if startF == "" {
		// provide default start time of YTD.
		start = &datetime.Datetime{Month: 1, Day: 1, Year: time.Now().Year()}
	} else {
		// parse start time.
		dt, err := time.Parse("2006-01-02", startF)
		if err != nil {
			return fmt.Errorf("could not parse start time- correct format is yyyy-mm-dd")
		}
		start = datetime.New(&dt)
	}
	// end.
	var end *datetime.Datetime
	if endF == "" {
		t := time.Now()
		end = datetime.New(&t)
	} else {
		// parse start time.
		dt, err := time.Parse("2006-01-02", endF)
		if err != nil {
			return fmt.Errorf("could not parse end time- correct format is yyyy-mm-dd")
		}
		end = datetime.New(&dt)
	}

	// validate aggregation periods.
	if aggregationF != "1d" && aggregationF != "5d" && aggregationF != "1mo" && aggregationF != "1y" {
		return fmt.Errorf("invalid aggregation period")
	}

	p := &chart.Params{
		Symbol:   symbols[0],
		Start:    start,
		End:      end,
		Interval: datetime.Interval(aggregationF),
	}

	// get chart.
	iter := chart.Get(p)

	// format.
	chartdata, err := formatC(iter)
	if err != nil {
		return err
	}

	// write.
	err = write(historyFields, symbols[0], "history", chartdata)
	if err != nil {
		fmt.Println("error writing data")
	}

	return nil
}

// format formats quotes into a writeable format.
func format(iter *quote.Iter) (data [][]string, err error) {
	for iter.Next() {
		q := iter.Quote()
		d := datetime.FromUnix(q.RegularMarketTime)
		date := fmt.Sprintf("%02d-%02d-%02d", d.Year, d.Month, d.Day)
		fq := []string{
			q.Symbol,
			date,
			utils.ToStringF(q.RegularMarketPrice),
			utils.ToStringF(q.RegularMarketChange),
			utils.ToStringF(q.RegularMarketChangePercent),
			utils.ToString(q.RegularMarketVolume),
			utils.ToStringF(q.Bid),
			utils.ToString(q.BidSize),
			utils.ToStringF(q.Ask),
			utils.ToString(q.AskSize),
			utils.ToStringF(q.RegularMarketOpen),
			utils.ToStringF(q.RegularMarketPreviousClose),
			q.ShortName,
		}
		data = append(data, fq)
	}
	return data, iter.Err()
}

// formatC formats chart data into a writeable format.
func formatC(iter *chart.Iter) (data [][]string, err error) {
	for iter.Next() {
		b := iter.Bar()
		d := datetime.FromUnix(b.Timestamp)
		date := fmt.Sprintf("%02d-%02d-%02d", d.Year, d.Month, d.Day)
		p := []string{
			date,
			b.Open.StringFixed(2),
			b.High.StringFixed(2),
			b.Low.StringFixed(2),
			b.Close.StringFixed(2),
			b.AdjClose.StringFixed(2),
			utils.ToString(b.Volume),
			iter.Meta().Symbol,
		}
		data = append(data, p)
	}
	return data, iter.Err()
}
