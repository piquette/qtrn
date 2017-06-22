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
	"fmt"
	"os"

	"github.com/FlashBoys/go-finance"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

const (
	quoteUsage     = "quote [symbols..]"
	quoteShortDesc = "Print stock quote table to the current shell"
	quoteLongDesc  = "Print stock quote table to the current shell. For more than just OHLC data, use --full or -f for a full quote."
)

var (
	// quote command.
	quoteCmd = &cobra.Command{
		Use:     quoteUsage,
		Short:   quoteShortDesc,
		Long:    quoteLongDesc,
		Aliases: []string{"q"},
		Example: "$ qtrn quote AAPL GOOG FB",
		Run:     quoteFunc,
	}
	// flagFullOutput set flag for a more informative quote.
	flagFullOutput bool
)

func init() {
	quoteCmd.Flags().BoolVarP(&flagFullOutput, "full", "f", false, "Set `--full` or `-f` for a more informative quote for each symbol")
}

// quoteFunc implements the quote command
func quoteFunc(cmd *cobra.Command, args []string) {

	// Print regular
	quotes, err := finance.GetQuotes(args[:])
	if err != nil {
		fmt.Println(err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)

	if flagFullOutput {
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.AppendBulk(setFullQuote(quotes))

	} else {

		table.SetHeader([]string{"Symbol", "Time", "Last", "Change", "Vol", "Bid", "Size", "Ask", "Size", "Open", "Prev Close", "Company"})
		table.AppendBulk(setTopOfBook(quotes))

	}

	table.Render()

}

func setTopOfBook(quotes []finance.Quote) (data [][]string) {

	for _, q := range quotes {

		timestamp := getFormattedDate(q)
		change := fmt.Sprintf("%s%s [%s%%]", getPrefix(q), q.ChangeNominal.String(), q.ChangePercent.String())
		data = append(data,
			[]string{
				q.Symbol,
				timestamp,
				q.LastTradePrice.String(),
				change,
				toString(q.Volume),
				q.Bid.String(),
				toString(q.BidSize),
				q.Ask.String(),
				toString(q.AskSize),
				q.Open.String(),
				q.PreviousClose.String(),
				q.Name,
			})

	}
	return
}

func setFullQuote(quotes []finance.Quote) (data [][]string) {

	for i, q := range quotes {

		timestamp := getFormattedDate(q)
		change := fmt.Sprintf("%s%s [%s%%]", getPrefix(q), q.ChangeNominal.String(), q.ChangePercent.String())

		data = append(data,
			[]string{"Symbol", q.Symbol},
			[]string{"Company", q.Name},
			[]string{"Time", timestamp},
			[]string{"Last", q.LastTradePrice.String()},
			[]string{"Change", change},
			[]string{"Vol", toString(q.Volume)},
			[]string{"Bid", q.Bid.String()},
			[]string{"Size", toString(q.BidSize)},
			[]string{"Ask", q.Ask.String()},
			[]string{"Size", toString(q.AskSize)},
			[]string{"Open", q.Open.String()},
			[]string{"Prev Close", q.PreviousClose.String()},
			[]string{"Exchange", q.Exchange},
			[]string{"Day High", q.DayHigh.String()},
			[]string{"Day Low", q.DayLow.String()},
			[]string{"52wk High", q.FiftyTwoWeekHigh.String()},
			[]string{"52wk Low", q.FiftyTwoWeekLow.String()},
			[]string{"Mkt Cap", q.MarketCap},
			[]string{"50D MA", q.FiftyDayMA.String()},
			[]string{"200D MA", q.TwoHundredDayMA.String()},
			[]string{"Avg Daily Vol", toString(q.AvgDailyVolume)},
			[]string{"EPS", q.EPS.String()},
			[]string{"P/E", q.PERatio.String()},
			[]string{"PEG Ratio", q.PEGRatio.String()},
			[]string{"P/S", q.PriceSales.String()},
			[]string{"P/B", q.PriceBook.String()},
			[]string{"Div", q.DivPerShare.String()},
			[]string{"Div Yield", q.DivYield.String()},
			[]string{"EPS Est Next Qtr", q.EPSEstNextQuarter.String()},
			[]string{"EPS Est Yr", q.EPSEstCurrentYear.String()},
			[]string{"Short Ratio", q.ShortRatio.String()},
			[]string{"Book Value", q.BookValue.String()},
			[]string{"EBITDA", q.EBITDA},
		)
		l := len(quotes)
		if l > 1 && i != l-1 {
			data = append(data, []string{"", ""})
		}
	}

	return
}
