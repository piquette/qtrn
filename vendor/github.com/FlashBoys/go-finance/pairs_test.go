package finance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetCurrencyPairQuote(t *testing.T) {

	s := startTestServer("pair_fixture.csv")
	defer s.Close()
	QuoteURL = s.URL

	p, err := GetCurrencyPairQuote(USDEUR)
	assert.Nil(t, err)

	// result should be a USDEUR pair.
	assert.Equal(t, USDEUR, p.Symbol)
}
