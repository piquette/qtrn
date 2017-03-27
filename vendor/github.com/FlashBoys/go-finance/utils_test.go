package finance

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func Test_ToInt(t *testing.T) {

	// result should be 34.
	assert.Equal(t, 34, toInt("34"))
	// result should be 0.
	assert.Equal(t, 0, toInt("-"))
}

func Test_ToDecimal(t *testing.T) {

	// result should be a decimal of 34.4.
	assert.Equal(t, decimal.NewFromFloat(34.4), toDecimal("34.4"))
	// result should be the Decimal zero-value.
	assert.Equal(t, decimal.Decimal{}, toDecimal("-"))
	// result should be a decimal of 0.34.
	assert.Equal(t, decimal.NewFromFloat(0.34), toDecimal("0.34%"))
}

func Test_ToEventValue(t *testing.T) {

	// event from split Ratio.
	split := toEventValue("1:5")
	// split should have a ratio of 1 to 5.
	assert.Equal(t, "1:5", split.Ratio)
	// split should not have a dividend amt.
	assert.Equal(t, decimal.Decimal{}, split.Dividend)

	// event from dividend value.
	div := toEventValue("3.02")
	// dividend should have a dividend amt of 3.02.
	assert.Equal(t, decimal.NewFromFloat(3.02), div.Dividend)
	// dividend should not have a split ratio.
	assert.Equal(t, "", div.Ratio)
}

func Test_ParseDashedDate(t *testing.T) {

	pd, err := parseDashedDate("2016-04-01")
	assert.Nil(t, err)

	loc, _ := time.LoadLocation("America/New_York")
	d := time.Date(2016, 4, 1, 0, 0, 0, 0, loc)

	// result should be April 1, 2016.
	assert.Equal(t, d.Year(), pd.Year())
	assert.Equal(t, d.Month(), pd.Month())
	assert.Equal(t, d.Day(), pd.Day())

	bd, err := parseDashedDate("N/A")
	assert.Nil(t, err)

	// result should be the time zero-value.
	assert.Equal(t, time.Time{}, bd)

	bd, err = parseDashedDate("5434")

	// result should be the time zero-value.
	assert.NotNil(t, err)
	assert.Equal(t, time.Time{}, bd)
}

func Test_ParseMalformedDate(t *testing.T) {

	// result should be a valid date string.
	assert.Equal(t, "2011-05-06", parseMalformedDate("020110506"))
}
