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

	// Print regular
	// quotes, err := finance.GetQuotes(args[:])
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

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

	table := tw.NewWriter(os.Stdout)

	// if flagFullOutput {
	// 	//table.SetAlignment(tablewriter.ALIGN_LEFT)
	// 	//table.AppendBulk(setFullQuote(quotes))
	// } else {
	// }

	table.SetHeader([]string{"Symbol", "Time", "Last", "Change", "Vol", "Bid", "Size", "Ask", "Size", "Open", "Prev Close", "Company"})
	table.AppendBulk(setTopOfBook(equities))

	table.Render()
}

func setTopOfBook(equities []*finance.Equity) (data [][]string) {

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
				e.LongName,
			})

	}
	return

}

// func setTopOfBook(quotes []finance.Quote) (data [][]string) {
//
// 	for _, q := range quotes {
//
// 		timestamp := getFormattedDate(q)
// 		change := fmt.Sprintf("%s%s [%s%%]", getPrefix(q), q.ChangeNominal.String(), q.ChangePercent.String())
// 		data = append(data,
// 			[]string{
// 				q.Symbol,
// 				timestamp,
// 				q.LastTradePrice.String(),
// 				change,
// 				toString(q.Volume),
// 				q.Bid.String(),
// 				toString(q.BidSize),
// 				q.Ask.String(),
// 				toString(q.AskSize),
// 				q.Open.String(),
// 				q.PreviousClose.String(),
// 				q.Name,
// 			})
//
// 	}
// 	return
// }
//
// func setFullQuote(quotes []finance.Quote) (data [][]string) {
//
// 	for i, q := range quotes {
//
// 		timestamp := getFormattedDate(q)
// 		change := fmt.Sprintf("%s%s [%s%%]", getPrefix(q), q.ChangeNominal.String(), q.ChangePercent.String())
//
// 		data = append(data,
// 			[]string{"Symbol", q.Symbol},
// 			[]string{"Company", q.Name},
// 			[]string{"Time", timestamp},
// 			[]string{"Last", q.LastTradePrice.String()},
// 			[]string{"Change", change},
// 			[]string{"Vol", toString(q.Volume)},
// 			[]string{"Bid", q.Bid.String()},
// 			[]string{"Size", toString(q.BidSize)},
// 			[]string{"Ask", q.Ask.String()},
// 			[]string{"Size", toString(q.AskSize)},
// 			[]string{"Open", q.Open.String()},
// 			[]string{"Prev Close", q.PreviousClose.String()},
// 			[]string{"Exchange", q.Exchange},
// 			[]string{"Day High", q.DayHigh.String()},
// 			[]string{"Day Low", q.DayLow.String()},
// 			[]string{"52wk High", q.FiftyTwoWeekHigh.String()},
// 			[]string{"52wk Low", q.FiftyTwoWeekLow.String()},
// 			[]string{"Mkt Cap", q.MarketCap},
// 			[]string{"50D MA", q.FiftyDayMA.String()},
// 			[]string{"200D MA", q.TwoHundredDayMA.String()},
// 			[]string{"Avg Daily Vol", toString(q.AvgDailyVolume)},
// 			[]string{"EPS", q.EPS.String()},
// 			[]string{"P/E", q.PERatio.String()},
// 			[]string{"PEG Ratio", q.PEGRatio.String()},
// 			[]string{"P/S", q.PriceSales.String()},
// 			[]string{"P/B", q.PriceBook.String()},
// 			[]string{"Div", q.DivPerShare.String()},
// 			[]string{"Div Yield", q.DivYield.String()},
// 			[]string{"EPS Est Next Qtr", q.EPSEstNextQuarter.String()},
// 			[]string{"EPS Est Yr", q.EPSEstCurrentYear.String()},
// 			[]string{"Short Ratio", q.ShortRatio.String()},
// 			[]string{"Book Value", q.BookValue.String()},
// 			[]string{"EBITDA", q.EBITDA},
// 		)
// 		l := len(quotes)
// 		if l > 1 && i != l-1 {
// 			data = append(data, []string{"", ""})
// 		}
// 	}
//
// 	return
// }
