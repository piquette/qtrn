package finance

import "fmt"

// NewCycle creates a new OptionsCycle instance.
func NewCycle(symbol string) (oc OptionsCycle, err error) {

	oc = OptionsCycle{Symbol: symbol}
	oc.Expirations, oc.UnderlyingPrice, err = newCycle(symbol)
	return
}

// GetChainForExpiration fetches the option chain for the given expiration datetime.
func (c *OptionsCycle) GetChainForExpiration(e Datetime) (calls []Contract, puts []Contract, err error) {

	if !c.expirationExists(e) {
		return calls, puts, fmt.Errorf("expiration does not exist")
	}

	url := buildOptionsURL(OptionsURL, c.Symbol, e)
	result, err := fetchOptions(url)
	if err != nil {
		return
	}

	return newChain(result.Calls), newChain(result.Puts), nil
}

// GetFrontMonth fetches the option chain for the front month.
func (c *OptionsCycle) GetFrontMonth() (calls []Contract, puts []Contract, err error) {
	return c.GetChainForExpiration(c.Expirations[0])
}

// GetCallsForExpiration fetches calls for the given expiration date.
func (c *OptionsCycle) GetCallsForExpiration(e Datetime) (calls []Contract, err error) {
	calls, _, err = c.GetChainForExpiration(e)
	return
}

// GetPutsForExpiration fetches puts for the given expiration date.
func (c *OptionsCycle) GetPutsForExpiration(e Datetime) (puts []Contract, err error) {
	_, puts, err = c.GetChainForExpiration(e)
	return
}
