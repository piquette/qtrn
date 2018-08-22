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

package chart

import (
	"fmt"

	ui "github.com/gizak/termui"
	"github.com/piquette/finance-go/datetime"
	"github.com/spf13/cobra"
)

const (
	usage = "chart [symbol]"
	short = "Print stock chart to the current shell (still in beta)"
	long  = "Print stock chart to the current shell using a symbol, time frame, and interval (still in beta)"
)

var (
	// Cmd is the chart command.
	Cmd = &cobra.Command{
		Use:     usage,
		Short:   short,
		Long:    long,
		Args:    cobra.ExactArgs(1),
		Aliases: []string{"c"},
		Example: "$ qtrn chart AAPL -s 2016-12-01 -e 2017-06-20 -i 1d",
		// Args: func(cmd *cobra.Command, args []string) error {
		//
		// 	// Fmt interval here..
		// 	return nil
		// 	//fmt.Errorf("invalid color specified: %s", args[0])
		// },
		RunE: execute,
	}
	// startF set flag to specify the start time of the chart frame.
	startF string
	// endF set flag to specify the end time of the chart frame.
	endF string
	// intervalF set flag to specify time interval of each chart point.
	intervalF string
)

func init() {
	// time frame, interval.
	Cmd.Flags().StringVarP(&startF, "start", "s", "", "Set a date (formatted YYYY-MM-DD) using `--start` or `-s` to specify the start of the chart's time frame")
	Cmd.Flags().StringVarP(&endF, "end", "e", "", "Set a date (formatted YYYY-MM-DD) using `--start` or `-s` to specify the start of the chart's time frame")
	Cmd.Flags().StringVarP(&intervalF, "interval", "i", string(datetime.OneDay), "Set an interval ( 1d | 1wk | 1mo ) using `--interval` or `-i` to specify the time interval of each chart point")
	Cmd.MarkFlagRequired("start")
	Cmd.MarkFlagRequired("end")
}

func execute(cmd *cobra.Command, args []string) error {

	symbol := args[0]
	points, dates, err := fetch(symbol, intervalF)
	if err != nil || len(points) == 0 {
		fmt.Printf("\nError fetching chart data, please try again\n\n")
		return err
	}

	err = ui.Init()
	if err != nil {
		fmt.Printf("\nCannot render chart\n\n")
		return err
	}
	defer ui.Close()

	draw(symbol, points, dates)

	return nil
}

func fetch(symbol string, interval string) (points []float64, dates []string, err error) {
	//
	// var start datetime.Datetime
	// var end datetime.Datetime
	//
	// // if flagStartTime == "" {
	// 	start = finance.ParseDatetime(fmt.Sprintf("%v-01-01", time.Now().Year()))
	// } else {
	// 	start = finance.ParseDatetime(flagStartTime)
	// }
	// if flagEndTime == "" {
	// 	t := time.Now()
	// 	end = finance.ParseDatetime(fmt.Sprintf("%d-%02d-%02d", t.Year(), int(t.Month()), t.Day()))
	// } else {
	// 	end = finance.ParseDatetime(flagEndTime)
	// }

	// bars, err := finance.GetHistory(symbol, start, end, finance.Interval(interval))
	// if err != nil {
	// 	return
	// }

	// for _, b := range bars {
	// 	close, _ := b.AdjClose.Round(2).Float64()
	// 	datetime := fmt.Sprintf("%v/%v/%v", b.Date.Month, b.Date.Day, b.Date.Year)
	// 	points = append(points, close)
	// 	dates = append(dates, datetime)
	// }
	return
}

func draw(symbol string, points []float64, dates []string) {

	chartPane := ui.NewLineChart()
	chartPane.Mode = "dot"
	chartPane.DotStyle = '+'
	chartPane.BorderLabel = fmt.Sprintf("  %+v Daily Chart (%+v - %+v)  ", symbol, dates[0], dates[len(dates)-1])
	chartPane.Data = points
	chartPane.DataLabels = dates

	chartPane.Width = len(points) + (len(points) / 10)
	chartPane.Height = 20
	chartPane.X = 0
	chartPane.Y = 0
	chartPane.AxesColor = ui.ColorWhite
	chartPane.LineColor = ui.ColorGreen | ui.AttrBold

	ui.Render(chartPane)
	ui.Handle("/sys/kbd", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Loop()

}
