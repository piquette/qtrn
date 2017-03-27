package finance

import "fmt"

// GetUSEquitySymbols fetches the symbols available through BATS, ~8k symbols.
func GetUSEquitySymbols() (symbols []string, err error) {

	t, err := fetchCSV(SymbolsURL)
	if err != nil {
		return []string{}, fmt.Errorf("error fetching symbols:  (error was: %s)\n", err.Error())
	}

	for i, r := range t {
		if i != 0 {
			symbols = append(symbols, r[0])
		}
	}
	return
}
