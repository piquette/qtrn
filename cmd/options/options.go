package options

import (
	"os"

	tw "github.com/olekukonko/tablewriter"
	finance "github.com/piquette/finance-go"
	"github.com/piquette/finance-go/options"
	"github.com/piquette/qtrn/utils"
	"github.com/spf13/cobra"
)

const (
	usage = "options"
	short = "Print options chain to the current shell"
	long  = "Print options chain to the current shell."
)

var (
	// Cmd is the options command.
	Cmd = &cobra.Command{
		Use:     usage,
		Short:   short,
		Long:    long,
		Aliases: []string{"o"},
		Example: "qtrn options AAPL",
		RunE:    execute,
	}
)

// execute implements the quote command
func execute(cmd *cobra.Command, args []string) error {

	symbol := args[0]
	straddles := []*finance.Straddle{}
	iter := options.GetStraddle(symbol)
	for iter.Next() {
		straddles = append(straddles, iter.Straddle())
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
	table.SetHeader([]string{"", "", "CALLS", "", "", utils.DateFS(iter.Meta().ExpirationDate), "", "", "PUTS", "", ""})
	table.AppendBulk(build(straddles))
	table.Render()

	return nil
}

// build builds table lines.
func build(ss []*finance.Straddle) (tbl [][]string) {
	// Get fields.
	fs := fields()
	fs = append(fs, utils.Bold("Strike"))
	fs = append(fs, fields()...)
	tbl = append(tbl, fs)

	for _, s := range ss {
		row := []string{}
		// Call
		call := s.Call
		if call != nil {
			row = append(row, utils.ToStringF(call.LastPrice))
			row = append(row, utils.ToStringF(call.Change))
			row = append(row, utils.ToStringF(call.PercentChange))
			row = append(row, utils.ToString(call.Volume))
			row = append(row, utils.ToString(call.OpenInterest))
		} else {
			row = append(row, "--")
			row = append(row, "--")
			row = append(row, "--")
			row = append(row, "--")
			row = append(row, "--")
		}
		// Strike.
		row = append(row, utils.Bold(utils.ToStringF(s.Strike)))
		// Put.
		put := s.Put
		if put != nil {
			row = append(row, utils.ToStringF(put.LastPrice))
			row = append(row, utils.ToStringF(put.Change))
			row = append(row, utils.ToStringF(put.PercentChange))
			row = append(row, utils.ToString(put.Volume))
			row = append(row, utils.ToString(put.OpenInterest))
		} else {
			row = append(row, "--")
			row = append(row, "--")
			row = append(row, "--")
			row = append(row, "--")
			row = append(row, "--")
		}

		tbl = append(tbl, row)
	}

	return tbl
}

func fields() []string {

	return []string{utils.Bold("Last Price"), utils.Bold("Change"), utils.Bold("% Change"), utils.Bold("Volume"), utils.Bold("Open Int")}
}
