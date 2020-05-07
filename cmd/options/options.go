package options

import (
	"fmt"
	"os"
	"time"

	tw "github.com/olekukonko/tablewriter"
	finance "github.com/piquette/finance-go"
	"github.com/piquette/finance-go/datetime"
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

	// listExpirationsF set flag to list the available expiration dates for the supplied symbol.
	listExpirationsF bool
	// expirationF set flag to specify expiration date for the supplied symbol.
	expirationF string
	//openInterestF set flag to filter results by greater than or equal to OI for supplied symbol.
	openInterestF int
	//volume set flag to filter results by greater than or equal to volume for supplied symbol.
	volumeF int
)

func init() {
	Cmd.Flags().BoolVarP(&listExpirationsF, "list", "l", false, "list the available expiration dates for the supplied symbol. default is false.")
	Cmd.Flags().StringVarP(&expirationF, "exp", "e", "", "set flag to specify expiration date for the supplied symbol. (formatted yyyy-mm-dd)")
	Cmd.Flags().IntVarP(&openInterestF, "openInterest", "o", -1, "set flag to specify showing only results >= (int)")
	Cmd.Flags().IntVarP(&volumeF, "volume", "v", -1, "set flag to specify showing only results >= (int)")
}

// execute implements the options command
func execute(cmd *cobra.Command, args []string) error {
	// check symbol.
	symbols := args
	if len(symbols) == 0 {
		return fmt.Errorf("no symbols provided")
	}
	// fetch options.
	p := &options.Params{
		UnderlyingSymbol: symbols[0],
	}
	// add expiration.
	if expirationF != "" {
		dt, err := time.Parse("2006-01-02", expirationF)
		if err != nil {
			return fmt.Errorf("could not parse expiration- correct format is yyyy-mm-dd")
		}
		p.Expiration = datetime.New(&dt)
	}
	iter := options.GetStraddleP(p)

	if listExpirationsF {
		return writeE(iter)
	}

	// write straddles.
	return write(iter)
}

// write writes the straddle table.
func write(iter *options.StraddleIter) error {
	// iterate.
	straddles := []*finance.Straddle{}
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
	table.SetHeader([]string{"", "", "Calls", "", "", utils.DateFS(iter.Meta().ExpirationDate + 86400), "", "", "Puts", "", ""})
	table.AppendBulk(build(straddles))
	table.Render()

	return nil
}

// writeE writes the expiration dates table.
func writeE(iter *options.StraddleIter) error {
	// iterate.
	meta := iter.Meta()
	if meta == nil {
		return fmt.Errorf("could not retrieve dates")
	}

	dates := [][]string{}
	for _, stamp := range meta.AllExpirationDates {
		// set the day to friday instead of EOD thursday..
		// weird math here..
		stamp = stamp + 86400
		t := time.Unix(int64(stamp), 0)
		dates = append(dates, []string{t.Format("2006-01-02")})
	}

	// Create table writer.
	table := tw.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	table.SetAlignment(tw.ALIGN_LEFT)
	table.SetCenterSeparator("*")
	table.SetColumnSeparator("|")
	table.SetHeader([]string{"Exp. Dates"})
	table.AppendBulk(dates)
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
		if call != nil && call.OpenInterest >= openInterestF && call.Volume >= volumeF{
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
		if put != nil && put.OpenInterest >= openInterestF && put.Volume >= volumeF{
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
