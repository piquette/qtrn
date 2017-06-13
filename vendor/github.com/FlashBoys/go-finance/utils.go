package finance

import (
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

// toInt converts a string to an int.
func toInt(value string) int {
	i, _ := strconv.Atoi(value)
	return i
}

// toDecimal converts a string to a decimal value.
func toDecimal(value string) (d decimal.Decimal) {

	value = strings.Replace(value, "%", "", -1)
	d, _ = decimal.NewFromString(value)
	return
}

func toEventValue(value string) Value {

	if strings.Contains(value, "/") {
		return Value{
			Ratio: value,
		}
	}
	return Value{
		Dividend: toDecimal(value),
	}
}

// parseDashedDate converts a string to a proper date and sets time to market close.
func parseDashedDate(s string) (d time.Time, err error) {

	if !strings.ContainsAny(s, "0123456789") {
		return
	}

	d, err = time.Parse("2006-01-02", s)
	if err != nil {
		s = parseMalformedDate(s)
		d, err = time.Parse("2006-01-02", s)
		if err != nil {
			return time.Time{}, err
		}
	}
	return
}

func parseMalformedDate(s string) string {

	if len(s) < 7 {
		return s
	}

	chars := strings.Split(s, "")
	chars = chars[1:]
	chars = insert(chars, 4, "-")
	chars = insert(chars, 7, "-")
	return strings.Join(chars[:], "")
}

func insert(s []string, i int, x string) []string {

	s = append(s, "")
	copy(s[i+1:], s[i:])
	s[i] = x
	return s
}
