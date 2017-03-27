package finance

import (
	"time"

	"github.com/shopspring/decimal"
)

type (
	// OptionsCycle contains the list of expirations for a symbol.
	OptionsCycle struct {
		Symbol          string
		UnderlyingPrice decimal.Decimal
		Expirations     []Datetime
	}

	// Contract represents an instance of an option contract.
	Contract struct {
		ID            string
		Security      string
		Strike        decimal.Decimal
		Price         decimal.Decimal
		Change        decimal.Decimal
		ChangePercent decimal.Decimal
		Bid           decimal.Decimal
		Ask           decimal.Decimal
		Volume        int
		OpenInterest  int
	}
)

// newContract creates a new instance of an option contract.
func newContract(option map[string]string) Contract {

	c := Contract{
		ID:           option["cid"],
		Security:     option["s"],
		Strike:       toDecimal(option["strike"]),
		Price:        toDecimal(option["p"]),
		Change:       toDecimal(option["c"]),
		Bid:          toDecimal(option["b"]),
		Ask:          toDecimal(option["a"]),
		Volume:       toInt(option["vol"]),
		OpenInterest: toInt(option["oi"]),
	}

	if c.Price.IntPart() != 0 {
		hundred, _ := decimal.NewFromString("100")
		c.ChangePercent = ((c.Change).Div(c.Price.Sub(c.Change))).Mul(hundred).Truncate(2)
	}

	return c
}

// newCycle fetches the expiration dates for an option cycle.
func newCycle(symbol string) (expirations []Datetime, price decimal.Decimal, err error) {

	url := buildOptionsURL(OptionsURL, symbol, NewDatetime(time.Now()))
	result, err := fetchOptions(url)
	if err != nil {
		return expirations, price, err
	}
	return result.Expirations, toDecimal(result.Price), err
}

// newChain creates a new chain of contracts.
func newChain(options []map[string]string) (contracts []Contract) {

	for _, opt := range options {
		contracts = append(contracts, newContract(opt))
	}
	return contracts
}

func (c *OptionsCycle) expirationExists(e Datetime) bool {

	for _, exp := range c.Expirations {
		if exp.Day == e.Day && exp.Month == e.Month && exp.Year == e.Year {
			return true
		}
	}
	return false
}
