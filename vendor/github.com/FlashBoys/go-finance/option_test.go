package finance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewContractSlice(t *testing.T) {

	// Given we have a map of option contract data,
	testSlice := []map[string]string{
		{
			"p":      "51.87",
			"c":      "0.00",
			"vol":    "-",
			"cid":    "1057390848156240",
			"cs":     "chb",
			"a":      "57.40",
			"oi":     "41",
			"expiry": "Apr 15, 2016",
			"name":   "",
			"strike": "55.00",
			"cp":     "0.00",
			"e":      "OPRA",
			"b":      "57.10",
			"s":      "AAPL160415C00055000",
		},
		{
			"p":      "50.35",
			"c":      "0.00",
			"vol":    "-",
			"cid":    "315601104477798",
			"cs":     "chb",
			"a":      "52.35",
			"oi":     "322",
			"expiry": "Apr 15, 2016",
			"name":   "",
			"strike": "60.00",
			"cp":     "0.00",
			"e":      "OPRA",
			"b":      "52.10",
			"s":      "AAPL160415C00060000",
		},
	}

	// When we create a new option contract instance,
	contracts := newChain(testSlice)

	// Then slice of contracts length should be equal.
	assert.Len(t, contracts, len(testSlice))

}

func Test_NewOptionContract(t *testing.T) {

	// Given we have a map of option contract data,
	testMap := map[string]string{
		"p":      "51.87",
		"c":      "0.00",
		"vol":    "-",
		"cid":    "1057390848156240",
		"cs":     "chb",
		"a":      "57.40",
		"oi":     "41",
		"expiry": "Apr 15, 2016",
		"name":   "",
		"strike": "55.00",
		"cp":     "0.00",
		"e":      "OPRA",
		"b":      "57.10",
		"s":      "AAPL160415C00055000",
	}

	// When we create a new option contract instance,
	oc := newContract(testMap)

	// Then some fields should equal-
	assert.Equal(t, "1057390848156240", oc.ID)
	assert.Equal(t, "AAPL160415C00055000", oc.Security)
	assert.Equal(t, 0, oc.Volume)
	assert.Equal(t, 41, oc.OpenInterest)

}
