package chart

import (
	"fmt"
	"time"

	"github.com/guptarohit/asciigraph"
	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
	"github.com/spf13/cobra"
)

const (
	usage = "chart"
	short = "Print sparkline chart to the current shell"
	long  = "Print sparkline chart to the current shell."
)

var (
	// Cmd is the options command.
	Cmd = &cobra.Command{
		Use:     usage,
		Short:   short,
		Long:    long,
		Aliases: []string{"c"},
		Example: "qtrn chart AAPL",
		RunE:    execute,
	}

	// periodF set flag to specify chart period for the supplied symbol.
	periodF string
)

func init() {
	Cmd.Flags().StringVarP(&periodF, "period", "p", "1mo", "set flag to specify chart period for the supplied symbol. (1d | 1wk | 1mo | 6mo | 1yr | 5yr)")
}

// execute implements the options command
func execute(cmd *cobra.Command, args []string) error {
	//check symbol.
	symbols := args
	if len(symbols) == 0 {
		return fmt.Errorf("no symbol provided")
	}

	var start int
	var intrval datetime.Interval
	end := int(time.Now().Unix())

	switch periodF {
	case "1d":
		end = 0
		intrval = datetime.FifteenMins
	case "1wk":
		start = int(time.Now().Add(-7 * 24 * time.Hour).Unix())
		intrval = datetime.OneDay
	case "6mo":
		start = int(time.Now().Add(-6 * 30 * 24 * time.Hour).Unix())
		intrval = datetime.OneDay
	case "1yr":
		start = int(time.Now().Add(-12 * 30 * 24 * time.Hour).Unix())
		intrval = datetime.FiveDay
	case "5yr":
		start = int(time.Now().Add(-5 * 12 * 30 * 24 * time.Hour).Unix())
		intrval = datetime.OneMonth
	default:
		start = int(time.Now().Add(-30 * 24 * time.Hour).Unix())
		intrval = datetime.OneDay
	}

	// fetch chart bars.
	params := &chart.Params{
		Symbol:   symbols[0],
		Interval: intrval,
	}
	if start != 0 {
		params.Start = datetime.FromUnix(start)
	}
	if end != 0 {
		params.End = datetime.FromUnix(end)
	}

	iter := chart.Get(params)

	data := []float64{}
	for iter.Next() {
		b := iter.Bar()
		fl, _ := b.Close.Round(2).Float64()
		data = append(data, fl)
	}
	if iter.Err() != nil {
		return iter.Err()
	}
	if len(data) == 0 {
		return fmt.Errorf("could not find chart data")
	}

	graph := asciigraph.Plot(data)
	fmt.Println(graph)
	return nil
}
