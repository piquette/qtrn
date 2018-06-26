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

package cli

import (
	"fmt"
	"os"

	tw "github.com/olekukonko/tablewriter"
	finance "github.com/piquette/finance-go"
	"github.com/piquette/finance-go/equity"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	equityUsage     = "equity [symbols..]"
	equityShortDesc = "Print equity quote table to the current shell"
	equityLongDesc  = "Print equity quote table to the current shell. For more than just OHLC data, use --full or -f for a full quote."
)

var (
	// equity command.
	equityCmd = &cobra.Command{
		Use:     equityUsage,
		Short:   equityShortDesc,
		Long:    equityLongDesc,
		Aliases: []string{"e"},
		Example: "$ qtrn equity AAPL GOOG FB",
		Run:     equityFunc,
	}
	// flagFullOutput set flag for a more informative equity quote.
	flagFullOutput bool
)

func init() {
	equityCmd.Flags().BoolVarP(&flagFullOutput, "full", "f", false, "Set `--full` or `-f` for a more informative equity quote for each symbol")
}

// equityFunc implements the equity command
func equityFunc(cmd *cobra.Command, args []string) {

	// Iter.
	i := equity.List(args)

	var equities []*finance.Equity
	for i.Next() {
		e := i.Equity()
		equities = append(equities, e)
	}
	if i.Err() != nil {
		log.WithFields(log.Fields{
			"message": i.Err().Error(),
		}).Fatal("an error occured")
		return
	}

	// Create table writer.
	table := tw.NewWriter(os.Stdout)

	// Determine full or not.
	if flagFullOutput {
		table.SetAlignment(tw.ALIGN_LEFT)
		table.AppendBulk(full(equities))
	} else {
		table.SetHeader([]string{"Symbol", "Time", "Last", "Change", "Vol", "Bid", "Size", "Ask", "Size", "Open", "Prev Close", "Company"})
		table.AppendBulk(topOfBook(equities))
	}

	table.Render()
}

// topOfBook creates a table with basic quote information.
func topOfBook(equities []*finance.Equity) (data [][]string) {

	for _, e := range equities {
		timestamp := getFormattedDate(e)
		change := fmt.Sprintf("%s%s [%s%%]", getPrefix(e), toStringF(e.Quote.RegularMarketChange), toStringF(e.RegularMarketChangePercent))
		data = append(data,
			[]string{
				e.Symbol,
				timestamp,
				toStringF(e.RegularMarketPrice),
				change,
				toString(e.RegularMarketVolume),
				toStringF(e.Bid),
				toString(e.BidSize),
				toStringF(e.Ask),
				toString(e.AskSize),
				toStringF(e.RegularMarketOpen),
				toStringF(e.RegularMarketPreviousClose),
				parseString(e.LongName),
			})

	}
	return
}

// full creates a table with a lot of quote information.
func full(equities []*finance.Equity) (data [][]string) {

	for i, e := range equities {
		timestamp := getFormattedDate(e)
		change := fmt.Sprintf("%s%s [%s%%]", getPrefix(e), toStringF(e.Quote.RegularMarketChange), toStringF(e.RegularMarketChangePercent))
		data = append(data,
			[]string{"Symbol", e.Symbol},
			[]string{"Company", parseString(e.LongName)},
			[]string{"Time", timestamp},
			[]string{"Last", toStringF(e.RegularMarketPrice)},
			[]string{"Change", change},
			[]string{"Vol", toString(e.RegularMarketVolume)},
			[]string{"Bid", toStringF(e.Bid)},
			[]string{"Size", toString(e.BidSize)},
			[]string{"Ask", toStringF(e.Ask)},
			[]string{"Size", toString(e.AskSize)},
			[]string{"Open", toStringF(e.RegularMarketOpen)},
			[]string{"Prev Close", toStringF(e.RegularMarketPreviousClose)},
			[]string{"Exchange", e.ExchangeID},
			[]string{"Day High", toStringF(e.RegularMarketDayHigh)},
			[]string{"Day Low", toStringF(e.RegularMarketDayLow)},
			[]string{"52wk High", toStringF(e.FiftyTwoWeekHigh)},
			[]string{"52wk Low", toStringF(e.FiftyTwoWeekLow)},
			[]string{"Mkt Cap", toString(int(e.MarketCap))},
			[]string{"50D MA", toStringF(e.FiftyDayAverage)},
			[]string{"200D MA", toStringF(e.TwoHundredDayAverage)},
			[]string{"Avg Daily Vol", toString(e.AverageDailyVolume10Day)},
			[]string{"EPS", toStringF(e.EpsForward)},
			[]string{"P/E", toStringF(e.ForwardPE)},
			[]string{"P/B", toStringF(e.PriceToBook)},
			[]string{"Div", toStringF(e.TrailingAnnualDividendRate)},
			[]string{"Div Yield", toStringF(e.TrailingAnnualDividendYield)},
			[]string{"Short Ratio", toStringF(e.EpsForward)},
		)

		l := len(equities)
		if l > 1 && i != l-1 {
			data = append(data, []string{"", ""})
		}
	}
	return
}
