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

package quote

import (
	"fmt"
	"os"

	tw "github.com/olekukonko/tablewriter"
	finance "github.com/piquette/finance-go"
	"github.com/piquette/finance-go/quote"
	"github.com/piquette/qtrn/utils"
	"github.com/spf13/cobra"
)

const (
	usage = "quote"
	short = "Print quote table to the current shell"
	long  = "Print quote table to the current shell. For more than just OHLC data, use --full or -f for a full quote."
)

var (
	// Cmd is the quote command.
	Cmd = &cobra.Command{
		Use:     usage,
		Short:   short,
		Long:    long,
		Aliases: []string{"q"},
		Example: "$ qtrn quote AAPL GOOG FB",
		RunE:    execute,
	}
	// fullF set flag for a more informative quote.
	fullF bool
)

func init() {
	Cmd.Flags().BoolVarP(&fullF, "full", "f", false, "Set `--full` or `-f` for a more informative quote for each symbol")
}

// execute implements the quote command
func execute(cmd *cobra.Command, args []string) error {
	var qs []*finance.Quote

	if len(args) == 3 {
		return fmt.Errorf("this sucks")
	}

	// Get quotes.
	iter := quote.List(args)
	for iter.Next() {
		qs = append(qs, iter.Quote())
	}
	// Check error.
	if iter.Err() != nil {
		return iter.Err()
	}
	// Create table writer.
	table := tw.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetAlignment(tw.ALIGN_LEFT)
	table.SetCenterSeparator("*")
	table.SetColumnSeparator("|")
	table.AppendBulk(build(qs))
	table.Render()

	return nil
}

// build builds table lines.
func build(qs []*finance.Quote) (tbl [][]string) {

	// Get fields.
	var fs []string
	if fullF {
		fs = FieldsFullQ()
	} else {
		fs = FieldsQ()
	}

	// Append fields and values.
	for _, f := range fs {
		line := []string{utils.Bold(f)}
		for _, q := range qs {
			cell := MapQ(f, q)
			line = append(line, cell)
		}
		tbl = append(tbl, line)
	}
	return tbl
}

// []string{"Mkt Cap", utils.ToString(int(e.MarketCap))},
// []string{"EPS", utils.ToStringF(e.EpsForward)},
// []string{"P/E", utils.ToStringF(e.ForwardPE)},
// []string{"P/B", utils.ToStringF(e.PriceToBook)},
// []string{"Div", utils.ToStringF(e.TrailingAnnualDividendRate)},
// []string{"Div Yield", utils.ToStringF(e.TrailingAnnualDividendYield)},
// []string{"Short Ratio", utils.ToString(e.EpsForward)},
