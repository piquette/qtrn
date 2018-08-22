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
	"os"

	tw "github.com/olekukonko/tablewriter"
	finance "github.com/piquette/finance-go"
	"github.com/piquette/finance-go/crypto"
	"github.com/piquette/finance-go/equity"
	"github.com/piquette/finance-go/etf"
	"github.com/piquette/finance-go/forex"
	"github.com/piquette/finance-go/future"
	"github.com/piquette/finance-go/index"
	"github.com/piquette/finance-go/mutualfund"
	"github.com/piquette/finance-go/option"
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
	// infoF set flag for a more informative quote.
	infoF bool
	// formatF set flag to request quotes to specific asset types.
	formatF string
)

func init() {
	Cmd.Flags().BoolVarP(&infoF, "info", "i", false, "set for a more informative quote for each symbol")
	Cmd.Flags().StringVarP(&formatF, "format", "f", "", "set (equity|etf|future|fx|crypto|fund|option|idx) to request formatted quotes to specific asset types ")
}

// execute implements the quote command
func execute(cmd *cobra.Command, args []string) error {
	//
	iter := quotes(args)
	//
	qs := []interface{}{}
	for iter.Next() {
		qs = append(qs, iter.Current())
	}
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
func build(qs []interface{}) (tbl [][]string) {
	// Get fields.
	fs := fields()

	// Append fields and values.
	for _, f := range fs {
		line := []string{utils.Bold(f)}
		for _, q := range qs {
			cell := value(f, q)
			line = append(line, cell)
		}
		tbl = append(tbl, line)
	}
	return tbl
}

// fields
func fields() []string {
	// Fields based on quote type.
	switch formatF {
	case "equity":
		return FieldsEquity()
	case "etf":
		return FieldsETF()
	case "option":
		return FieldsOption()
	case "future":
		return FieldsFuture()
	case "crypto":
		return FieldsCrypto()
	case "fx":
		return FieldsForex()
	case "idx":
		return FieldIndex()
	case "fund":
		return FieldsMutualFund()
	default:
		if infoF {
			return FieldsInfoQuote()
		}
		return FieldsQuote()
	}
}

func value(field string, quote interface{}) string {
	// Fields based on quote type.
	switch formatF {
	case "equity":
		return MapEquity(field, quote.(*finance.Equity))
	case "etf":
		return MapETF(field, quote.(*finance.ETF))
	case "option":
		return MapOption(field, quote.(*finance.Option))
	case "future":
		return MapFuture(field, quote.(*finance.Future))
	case "crypto":
		return MapCrypto(field, quote.(*finance.CryptoPair))
	case "fx":
		return MapForex(field, quote.(*finance.ForexPair))
	case "idx":
		return MapIndex(field, quote.(*finance.Index))
	case "fund":
		return MapMutualFund(field, quote.(*finance.MutualFund))
	default:
		if infoF {
			return MapQuote(field, quote.(*finance.Quote))
		}
		return MapQuote(field, quote.(*finance.Quote))
	}
}

// quotes
func quotes(symbols []string) quoteIter {
	// Request based on quote type.
	switch formatF {
	case "equity":
		return equity.List(symbols)
	case "etf":
		return etf.List(symbols)
	case "option":
		return option.List(symbols)
	case "future":
		return future.List(symbols)
	case "crypto":
		return crypto.List(symbols)
	case "fx":
		return forex.List(symbols)
	case "idx":
		return index.List(symbols)
	case "fund":
		return mutualfund.List(symbols)
	default:
		return quote.List(symbols)
	}
}

// generic iter interface.
type quoteIter interface {
	Next() bool
	Current() interface{}
	Err() error
}
