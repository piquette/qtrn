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

package cmd

//
// import (
// 	"fmt"
// 	"time"
//
// 	ui "github.com/gizak/termui"
// 	finance "github.com/piquette/go-finance"
// 	"github.com/spf13/cobra"
// )
//
// const (
// 	chartUsage     = "chart [symbol]"
// 	chartShortDesc = "Print stock chart to the current shell (still in beta)"
// 	chartLongDesc  = "Print stock chart to the current shell using a symbol, time frame, and interval (still in beta)"
// )
//
// var (
// 	// chart command.
// 	chartCmd = &cobra.Command{
// 		Use:     chartUsage,
// 		Short:   chartShortDesc,
// 		Long:    chartLongDesc,
// 		Aliases: []string{"c"},
// 		Example: "$ qtrn chart AAPL -s 2016-12-01 -e 2017-06-20 -i 1d",
// 		Run:     chartFunc,
// 	}
// 	// flagStartTime set flag to specify the start time of the chart frame.
// 	flagStartTime string
// 	// flagEndTime set flag to specify the end time of the chart frame.
// 	flagEndTime string
// 	// flagInterval set flag to specify time interval of each chart point.
// 	flagInterval string
// )
//
// func init() {
// 	// time frame, interval.
// 	chartCmd.Flags().StringVarP(&flagStartTime, "start", "s", "", "Set a date (formatted YYYY-MM-DD) using `--start` or `-s` to specify the start of the chart's time frame")
// 	chartCmd.Flags().StringVarP(&flagEndTime, "end", "e", "", "Set a date (formatted YYYY-MM-DD) using `--start` or `-s` to specify the start of the chart's time frame")
// 	chartCmd.Flags().StringVarP(&flagInterval, "interval", "i", finance.Day, "Set an interval ( 1d | 1wk | 1mo ) using `--interval` or `-i` to specify the time interval of each chart point")
// }
//
// // chartFunc implements the chart command
// func chartFunc(cmd *cobra.Command, args []string) {
//
// 	if len(args) == 0 {
// 		fmt.Printf("\nPlease provide a symbol\n\n")
// 		return
// 	}
//
// 	if len(args) > 1 {
// 		fmt.Printf("\nToo many symbols, only 1 symbol is allowed for charting.\n\n")
// 		return
// 	}
// 	sym := args[0]
// 	points, dates, err := fetchChartPoints(sym, flagInterval)
// 	if err != nil || len(points) == 0 {
// 		fmt.Printf("\nError fetching chart data, please try again\n\n")
// 		fmt.Println(err)
// 		return
// 	}
//
// 	err = ui.Init()
// 	if err != nil {
// 		fmt.Printf("\nCannot render chart\n\n")
// 		return
// 	}
// 	defer ui.Close()
//
// 	draw(sym, points, dates)
// }
//
// func fetchChartPoints(symbol string, interval string) (points []float64, dates []string, err error) {
//
// 	var start finance.Datetime
// 	var end finance.Datetime
// 	if flagStartTime == "" {
// 		start = finance.ParseDatetime(fmt.Sprintf("%v-01-01", time.Now().Year()))
// 	} else {
// 		start = finance.ParseDatetime(flagStartTime)
// 	}
// 	if flagEndTime == "" {
// 		t := time.Now()
// 		end = finance.ParseDatetime(fmt.Sprintf("%d-%02d-%02d", t.Year(), int(t.Month()), t.Day()))
// 	} else {
// 		end = finance.ParseDatetime(flagEndTime)
// 	}
//
// 	bars, err := finance.GetHistory(symbol, start, end, finance.Interval(interval))
// 	if err != nil {
// 		return
// 	}
//
// 	for _, b := range bars {
// 		close, _ := b.AdjClose.Round(2).Float64()
// 		datetime := fmt.Sprintf("%v/%v/%v", b.Date.Month, b.Date.Day, b.Date.Year)
// 		points = append(points, close)
// 		dates = append(dates, datetime)
// 	}
// 	return
// }
//
// func draw(symbol string, points []float64, dates []string) {
//
// 	chartPane := ui.NewLineChart()
// 	chartPane.Mode = "dot"
// 	chartPane.DotStyle = '+'
// 	chartPane.BorderLabel = fmt.Sprintf("  %+v Daily Chart (%+v - %+v)  ", symbol, dates[0], dates[len(dates)-1])
// 	chartPane.Data = points
// 	chartPane.DataLabels = dates
//
// 	chartPane.Width = len(points) + (len(points) / 10)
// 	chartPane.Height = 20
// 	chartPane.X = 0
// 	chartPane.Y = 0
// 	chartPane.AxesColor = ui.ColorWhite
// 	chartPane.LineColor = ui.ColorGreen | ui.AttrBold
//
// 	ui.Render(chartPane)
// 	ui.Handle("/sys/kbd", func(ui.Event) {
// 		ui.StopLoop()
// 	})
// 	ui.Loop()
//
// }
