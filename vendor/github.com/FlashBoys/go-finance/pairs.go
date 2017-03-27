package finance

import "github.com/shopspring/decimal"

const (
	// USDGBP pair.
	USDGBP = "USDGBP=X"
	// USDEUR pair.
	USDEUR = "USDEUR=X"
	// USDAUD pair.
	USDAUD = "USDAUD=X"
	// USDCHF pair.
	USDCHF = "USDCHF=X"
	// USDJPY pair.
	USDJPY = "USDJPY=X"
	// USDCAD pair.
	USDCAD = "USDCAD=X"
	// USDSGD pair.
	USDSGD = "USDSGD=X"
	// USDNZD pair.
	USDNZD = "USDNZD=X"
	// USDHKD pair.
	USDHKD = "USDHKD=X"
	// GBPUSD pair.
	GBPUSD = "GBPUSD=X"
	// GBPEUR pair.
	GBPEUR = "GBPEUR=X"
	// GBPAUD pair.
	GBPAUD = "GBPAUD=X"
	// GBPCHF pair.
	GBPCHF = "GBPCHF=X"
	// GBPJPY pair.
	GBPJPY = "GBPJPY=X"
	// GBPCAD pair.
	GBPCAD = "GBPCAD=X"
	// GBPSGD pair.
	GBPSGD = "GBPSGD=X"
	// GBPNZD pair.
	GBPNZD = "GBPNZD=X"
	// GBPHKD pair.
	GBPHKD = "GBPHKD=X"
	// EURUSD pair.
	EURUSD = "EURUSD=X"
	// EURGBP pair.
	EURGBP = "EURGBP=X"
	// EURAUD pair.
	EURAUD = "EURAUD=X"
	// EURCHF pair.
	EURCHF = "EURCHF=X"
	// EURJPY pair.
	EURJPY = "EURJPY=X"
	// EURCAD pair.
	EURCAD = "EURCAD=X"
	// EURSGD pair.
	EURSGD = "EURSGD=X"
	// EURNZD pair.
	EURNZD = "EURNZD=X"
	// EURHKD pair.
	EURHKD = "EURHKD=X"
	// AUDUSD pair.
	AUDUSD = "AUDUSD=X"
	// AUDGBP pair.
	AUDGBP = "AUDGBP=X"
	// AUDEUR pair.
	AUDEUR = "AUDEUR=X"
	// AUDCHF pair.
	AUDCHF = "AUDCHF=X"
	// AUDJPY pair.
	AUDJPY = "AUDJPY=X"
	// AUDCAD pair.
	AUDCAD = "AUDCAD=X"
	// AUDSGD pair.
	AUDSGD = "AUDSGD=X"
	// AUDNZD pair.
	AUDNZD = "AUDNZD=X"
	// AUDHKD pair.
	AUDHKD = "AUDHKD=X"
	// CHFGBP pair.
	CHFGBP = "CHFGBP=X"
	// CHFEUR pair.
	CHFEUR = "CHFEUR=X"
	// CHFAUD pair.
	CHFAUD = "CHFAUD=X"
	// CHFJPY pair.
	CHFJPY = "CHFJPY=X"
	// CHFCAD pair.
	CHFCAD = "CHFCAD=X"
	// CHFSGD pair.
	CHFSGD = "CHFSGD=X"
	// CHFNZD pair.
	CHFNZD = "CHFNZD=X"
	// CHFHKD pair.
	CHFHKD = "CHFHKD=X"
	// JPYUSD pair.
	JPYUSD = "JPYUSD=X"
	// JPYGBP pair.
	JPYGBP = "JPYGBP=X"
	// JPYEUR pair.
	JPYEUR = "JPYEUR=X"
	// JPYAUD pair.
	JPYAUD = "JPYAUD=X"
	// JPYCHF pair.
	JPYCHF = "JPYCHF=X"
	// JPYCAD pair.
	JPYCAD = "JPYCAD=X"
	// JPYSGD pair.
	JPYSGD = "JPYSGD=X"
	// JPYNZD pair.
	JPYNZD = "JPYNZD=X"
	// JPYHKD pair.
	JPYHKD = "JPYHKD=X"
	// CADUSD pair.
	CADUSD = "CADUSD=X"
	// CADGBP pair.
	CADGBP = "CADGBP=X"
	// CADEUR pair.
	CADEUR = "CADEUR=X"
	// CADAUD pair.
	CADAUD = "CADAUD=X"
	// CADCHF pair.
	CADCHF = "CADCHF=X"
	// CADJPY pair.
	CADJPY = "CADJPY=X"
	// CADSGD pair.
	CADSGD = "CADSGD=X"
	// CADNZD pair.
	CADNZD = "CADNZD=X"
	// CADHKD pair.
	CADHKD = "CADHKD=X"
	// SGDUSD pair.
	SGDUSD = "SGDUSD=X"
	// SGDGBP pair.
	SGDGBP = "SGDGBP=X"
	// SGDEUR pair.
	SGDEUR = "SGDEUR=X"
	// SGDAUD pair.
	SGDAUD = "SGDAUD=X"
	// SGDCHF pair.
	SGDCHF = "SGDCHF=X"
	// SGDJPY pair.
	SGDJPY = "SGDJPY=X"
	// SGDCAD pair.
	SGDCAD = "SGDCAD=X"
	// SGDNZD pair.
	SGDNZD = "SGDNZD=X"
	// SGDHKD pair.
	SGDHKD = "SGDHKD=X"
	// NZDUSD pair.
	NZDUSD = "NZDUSD=X"
	// NZDGBP pair.
	NZDGBP = "NZDGBP=X"
	// NZDEUR pair.
	NZDEUR = "NZDEUR=X"
	// NZDAUD pair.
	NZDAUD = "NZDAUD=X"
	// NZDCHF pair.
	NZDCHF = "NZDCHF=X"
	// NZDJPY pair.
	NZDJPY = "NZDJPY=X"
	// NZDCAD pair.
	NZDCAD = "NZDCAD=X"
	// NZDSGD pair.
	NZDSGD = "NZDSGD=X"
	// NZDHKD pair.
	NZDHKD = "NZDHKD=X"
	// HKDUSD pair.
	HKDUSD = "HKDUSD=X"
	// HKDGBP pair.
	HKDGBP = "HKDGBP=X"
	// HKDEUR pair.
	HKDEUR = "HKDEUR=X"
	// HKDAUD pair.
	HKDAUD = "HKDAUD=X"
	// HKDCHF pair.
	HKDCHF = "HKDCHF=X"
	// HKDJPY pair.
	HKDJPY = "HKDJPY=X"
	// HKDCAD pair.
	HKDCAD = "HKDCAD=X"
	// HKDSGD pair.
	HKDSGD = "HKDSGD=X"
	// HKDNZD pair.
	HKDNZD = "HKDNZD=X"
)

// FXPairQuote represents the quote of a currency pair.
type FXPairQuote struct {
	Symbol           string          `yfin:"s"`
	PairName         string          `yfin:"n"`
	LastTradeTime    Datetime        `yfin:"t1"`
	LastTradeDate    Datetime        `yfin:"d1"`
	LastRate         decimal.Decimal `yfin:"l1"`
	ChangeNominal    decimal.Decimal `yfin:"c1"`
	ChangePercent    decimal.Decimal `yfin:"p2"`
	DayLow           decimal.Decimal `yfin:"g"`
	DayHigh          decimal.Decimal `yfin:"h"`
	FiftyTwoWeekLow  decimal.Decimal `yfin:"j"`
	FiftyTwoWeekHigh decimal.Decimal `yfin:"k"`
}

// GetCurrencyPairQuote fetches a single currency pair's quote from Yahoo Finance.
func GetCurrencyPairQuote(symbol string) (fq FXPairQuote, err error) {

	f, c := structFields(fq)
	params := map[string]string{
		"s": symbol,
		"f": f,
		"e": ".csv",
	}

	t, err := fetchCSV(buildURL(QuoteURL, params))
	if err != nil {
		return
	}

	mapFields(t[0], c, &fq)

	return
}
