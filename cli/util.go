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
	"strconv"
	"strings"
	"time"

	finance "github.com/piquette/finance-go"
)

func getPrefix(e *finance.Equity) string {

	pricelast := e.RegularMarketPrice
	priceclose := e.RegularMarketPreviousClose
	if pricelast > priceclose {
		return "+"
	}

	if pricelast < priceclose {
		return ""
	}
	return " "
}

func getFormattedDate(e *finance.Equity) string {
	stamp := e.Quote.RegularMarketTime
	dt := time.Unix(int64(stamp), 0)
	y, m, d := dt.Date()
	hr, min, sec := dt.Clock()
	return fmt.Sprintf("%02d:%02d:%02d %02d/%02d/%d", hr, min, sec, int(m), d, y)
}

// toInt converts a string to an int.
func toInt(value string) int {
	i, _ := strconv.Atoi(value)
	return i
}

// toString converts an int to a string.
func toString(v int) string {
	return strconv.Itoa(v)
}

// toStringF converts an int to a string.
func toStringF(v float64) string {
	return fmt.Sprintf("%.2f", v)
}

// capitalizes a string.
func capitalize(str string) string {
	return strings.ToUpper(str)
}

// combine a string.
func combine(strs []string, sep string) string {
	return strings.Join(strs, sep)
}
