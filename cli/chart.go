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
	"time"

	finance "github.com/FlashBoys/go-finance"
	ui "github.com/gizak/termui"
	"github.com/spf13/cobra"
)

const (
	chartUsage     = "chart"
	chartShortDesc = "Print stock chart to the current shell"
	chartLongDesc  = ""
)

var (
	chartCmd = &cobra.Command{
		Use:          chartUsage,
		Short:        chartShortDesc,
		Long:         chartLongDesc,
		Run:          executeChart,
		SilenceUsage: true,
	}
)

func init() {
	// symbol, frame, interval,
}

// executeChart implements the chart command
func executeChart(cmd *cobra.Command, args []string) {

	if len(args) > 1 {
		fmt.Println("incorrect number of parameters")
		return
	}
	sym := args[0]
	p, d, err := fetchChartPoints(sym, "", finance.Day)
	if err != nil {
		panic(err)
	}

	if len(p) == 0 {
		panic("no ")
	}

	err = ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	draw(sym, p, d)
}

func fetchChartPoints(symbol string, frame string, interval finance.Interval) (points []float64, dates []string, err error) {

	start := finance.ParseDatetime("1/1/2017")
	end := finance.NewDatetime(time.Now())

	bars, err := finance.GetHistory(symbol, start, end, interval)
	if err != nil {
		return
	}

	for _, b := range bars {
		close, _ := b.AdjClose.Round(2).Float64()
		datetime := fmt.Sprintf("%v/%v/%v", b.Date.Month, b.Date.Day, b.Date.Year)
		points = append(points, close)
		dates = append(dates, datetime)
	}
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
