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
	"strconv"
	"strings"

	"github.com/FlashBoys/go-finance"
)

func getPrefix(q finance.Quote) string {

	pricelast, _ := q.LastTradePrice.Float64()
	priceclose, _ := q.PreviousClose.Float64()
	if pricelast > priceclose {
		return "+"
	}

	if pricelast < priceclose {
		return ""
	}
	return " "
}

func getFormattedDate(q finance.Quote) string {
	return fmt.Sprintf("%02d:%02d:%02d %02d/%02d/%d", q.LastTradeTime.Hour, q.LastTradeTime.Minute, q.LastTradeTime.Second, q.LastTradeDate.Month, q.LastTradeDate.Day, q.LastTradeDate.Year)
}

// toInt converts a string to an int.
func toInt(value string) int {
	i, _ := strconv.Atoi(value)
	return i
}

// toString converts an int to a string.
func toString(value int) string {
	return strconv.Itoa(value)
}

// capitalizes a string.
func capitalize(str string) string {
	return strings.ToUpper(str)
}

// combine a string.
func combine(strs []string, sep string) string {
	return strings.Join(strs, sep)
}
