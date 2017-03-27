package finance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetUSEquitySymbols(t *testing.T) {

	s := startTestServer("symbols_fixture.csv")
	defer s.Close()
	SymbolsURL = s.URL

	symbols, err := GetUSEquitySymbols()
	assert.Nil(t, err)

	// second symbol should be Alcoa.
	assert.Equal(t, "AA", symbols[1])
	// slice should contain AMD.
	assert.Contains(t, symbols, "AMD")
}
