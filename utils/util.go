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

package utils

import (
	"fmt"
	"html"
	"strconv"
	"strings"
	"time"

	humanize "github.com/dustin/go-humanize"
	finance "github.com/piquette/finance-go"
)

// Direction is a price direction.
type Direction int

const (
	Flat Direction = iota
	Down
	Up
)

// ESC is the escape string.
const ESC = "\033"

// PriceDirection returns a plus/minus indicating price direction.
func PriceDirection(q *finance.Quote) Direction {

	last := q.RegularMarketPrice
	close := q.RegularMarketPreviousClose
	if last > close {
		return Up
	}
	if last < close {
		return Down
	}
	return Flat
}

// Bold makes a string bold.
func Bold(s string) string {
	return fmt.Sprintf("%s[%dm%s%s[%dm", ESC, 1, s, ESC, 0)
}

// MktStateF formats market state.
func MktStateF(m finance.MarketState) string {
	switch m {
	case finance.MarketStateRegular:
		return "Open"
	case finance.MarketStatePre,
		finance.MarketStatePrePre:
		return "Pre-Market"
	case finance.MarketStatePost,
		finance.MarketStatePostPost:
		return "After-Hours"
	default:
		return "Closed"
	}
}

// Color formats a string according to price direction.
func Color(s string, d Direction) string {
	if d == Flat {
		return s
	}

	code := "31"
	in := s
	if d == Up {
		in = "+" + s
		code = "32"
	}

	parts := strings.Split(in, " ")
	var out string

	// Down.
	pre := fmt.Sprintf("%s[%sm", ESC, code)
	post := fmt.Sprintf("%s[%dm", ESC, 0)

	for i, str := range parts {
		out = out + pre + str + post
		if i == 0 {
			out = out + " "
		}
	}

	return out
}

// NumberF formats a big number with commas.
func NumberF(i int) string {
	return humanize.Comma(int64(i))
}

// Strip strips weird html strings.
func Strip(s string) string {
	s = strings.Replace(s, "&nbsp;", "", -1)
	return html.UnescapeString(s)
}

// DateF returns a formatted date string from a quote.
func DateF(q *finance.Quote) string {
	stamp := q.RegularMarketTime
	dt := time.Unix(int64(stamp), 0)
	y, m, d := dt.Date()
	hr, min, sec := dt.Clock()
	return fmt.Sprintf("%02d:%02d:%02d %02d/%02d/%d", hr, min, sec, int(m), d, y)
}

// ToInt converts a string to an int.
func ToInt(value string) int {
	i, _ := strconv.Atoi(value)
	return i
}

// ToString converts an int to a string.
func ToString(v int) string {
	return strconv.Itoa(v)
}

// ToStringF converts an int to a string.
func ToStringF(v float64) string {
	return fmt.Sprintf("%.2f", v)
}

// Capitalize a string.
func Capitalize(str string) string {
	return strings.ToUpper(str)
}

// combine a string.
func combine(strs []string, sep string) string {
	return strings.Join(strs, sep)
}
