package quote

import (
	"fmt"
	"strings"

	finance "github.com/piquette/finance-go"
	"github.com/piquette/qtrn/utils"
)

const (
	// Quote fields.
	Symbol      = "Symbol"
	Market      = "Market"
	Time        = "Time"
	Last        = "Last"
	Change      = "Change"
	Vol         = "Volume"
	Bid         = "Bid"
	BidSize     = "Bid Size"
	Ask         = "Ask"
	AskSize     = "Ask Size"
	Open        = "Open"
	PrevClose   = "Prev Close"
	Exchange    = "Exchange"
	DayHigh     = "Day High"
	DayLow      = "Day Low"
	YearHigh    = "52wk High"
	YearLow     = "52wk Low"
	FiftyMA     = "50D MA"
	THundMA     = "200D MA"
	AvgDailyVol = "Avg. Vol"
	Security    = "Security"
	Name        = "Name"
	// Equity fields.
	EpsTrailing = "EPS Trailing"
	EpsForward  = "EPS Forward"
	EpsDate     = "Earnings Date"
	DivTrailing = "Div. Trailing"
	DivYield    = "Div. Yield"
	DivDate     = "Ex. Div. Date"
	PETrailing  = "P/E Trailing"
	PEForward   = "P/E Forward"
	BookValue   = "Book Value"
	PB          = "P/B"
	MarketCap   = "Market Cap"
	// ETF/MutualFund fields.
	YTDReturn    = "YTD Return"
	QtrReturn    = "Qtr Return"
	QtrNavReturn = "Qtr NAV Return"
	// Option/Future fields.
	Underlier    = "Underlying Symbol"
	OpenInterest = "Open Interest"
	ExpireDate   = "Expiration"
	Strike       = "Strike"
	// Crypto fields.
	Algorithm   = "Algorithm"
	StartDate   = "Start"
	MaxSupply   = "Max Supply"
	Circulating = "Circulation"
)

// FieldsQuote returns the fields for a plain quote.
func FieldsQuote() (fields []string) {
	return []string{Symbol, Market, Time, Last, Change, Vol,
		Bid, BidSize, Ask, AskSize, Open, PrevClose}
}

// FieldsInfoQuote returns the fields for an informative quote.
func FieldsInfoQuote() (fields []string) {
	return append(FieldsQuote(), Exchange,
		DayHigh, DayLow, YearHigh, YearLow, FiftyMA, THundMA, AvgDailyVol, Security, Name)
}

// FieldsEquity returns the fields for an informative quote.
func FieldsEquity() (fields []string) {
	return append(FieldsInfoQuote(), EpsTrailing, EpsForward, EpsDate, DivTrailing,
		DivYield, DivDate, PETrailing, PEForward, BookValue, PB, MarketCap)
}

// FieldsETF returns the fields for an etf.
func FieldsETF() (fields []string) {
	return append(FieldsInfoQuote(), YTDReturn, QtrReturn, QtrNavReturn)
}

// FieldsFuture returns the fields for a future.
func FieldsFuture() (fields []string) {
	return append([]string{Underlier, OpenInterest, ExpireDate, Strike}, FieldsInfoQuote()...)
}

// FieldsForex returns the fields for an informative quote.
func FieldsForex() (fields []string) {
	return append(FieldsInfoQuote(), Exchange,
		DayHigh, DayLow, YearHigh, YearLow, FiftyMA, THundMA, AvgDailyVol)
}

// FieldsCrypto returns the fields for an informative quote.
func FieldsCrypto() (fields []string) {
	return append(FieldsInfoQuote(), Algorithm, StartDate, MaxSupply, Circulating)
}

// FieldIndex returns the fields for an informative quote.
func FieldIndex() (fields []string) {
	return FieldsInfoQuote()
}

// FieldsOption returns the fields for an informative quote.
func FieldsOption() (fields []string) {
	return append([]string{Underlier, OpenInterest, ExpireDate, Strike}, FieldsInfoQuote()...)
}

// FieldsMutualFund returns the fields for an informative quote.
func FieldsMutualFund() (fields []string) {
	return append(FieldsInfoQuote(), YTDReturn, QtrReturn, QtrNavReturn)
}

// MapQuote maps quote fields.
func MapQuote(field string, q *finance.Quote) string {
	switch field {
	case Symbol:
		return q.Symbol
	case Market:
		return utils.MktStateF(q.MarketState)
	case Time:
		return utils.DateF(q.RegularMarketTime)
	case Last:
		return utils.ToStringF(q.RegularMarketPrice)
	case Change:
		{
			direction := utils.PriceDirection(q)
			change := fmt.Sprintf("%s [%s%%]", utils.ToStringF(q.RegularMarketChange), utils.ToStringF(q.RegularMarketChangePercent))
			return utils.Color(change, direction)
		}
	case Vol:
		return utils.NumberF(q.RegularMarketVolume)
	case Bid:
		return utils.ToStringF(q.Bid)
	case BidSize:
		return utils.ToString(q.BidSize * 100)
	case Ask:
		return utils.ToStringF(q.Ask)
	case AskSize:
		return utils.ToString(q.AskSize * 100)
	case Open:
		return utils.ToStringF(q.RegularMarketOpen)
	case PrevClose:
		return utils.ToStringF(q.RegularMarketPreviousClose)
	case Exchange:
		return q.ExchangeID
	case DayHigh:
		return utils.ToStringF(q.RegularMarketDayHigh)
	case DayLow:
		return utils.ToStringF(q.RegularMarketDayLow)
	case YearHigh:
		return utils.ToStringF(q.FiftyTwoWeekHigh)
	case YearLow:
		return utils.ToStringF(q.FiftyTwoWeekLow)
	case FiftyMA:
		return utils.ToStringF(q.FiftyDayAverage)
	case THundMA:
		return utils.ToStringF(q.TwoHundredDayAverage)
	case AvgDailyVol:
		return utils.NumberF(q.AverageDailyVolume10Day)
	case Name:
		return q.ShortName
	case Security:
		return strings.ToTitle(string(q.QuoteType))
	default:
		return ""
	}
}

// MapEquity maps equity fields.
func MapEquity(field string, q *finance.Equity) string {
	v := MapQuote(field, &q.Quote)
	if v != "" {
		return v
	}
	switch field {
	case EpsTrailing:
		return utils.ToStringF(q.EpsTrailingTwelveMonths)
	case EpsForward:
		return utils.ToStringF(q.EpsForward)
	case EpsDate:
		return utils.DateF(q.EarningsTimestamp)
	case DivTrailing:
		return utils.ToStringF(q.TrailingAnnualDividendRate)
	case DivYield:
		return utils.ToStringF(q.TrailingAnnualDividendYield)
	case DivDate:
		return utils.DateF(q.DividendDate)
	case PETrailing:
		return utils.ToStringF(q.TrailingPE)
	case PEForward:
		return utils.ToStringF(q.ForwardPE)
	case BookValue:
		return utils.ToStringF(q.BookValue)
	case PB:
		return utils.ToStringF(q.PriceToBook)
	case MarketCap:
		return utils.NumberFancyF(q.MarketCap)
	default:
		return ""
	}
}

// MapETF maps ETF fields.
func MapETF(field string, q *finance.ETF) string {
	v := MapQuote(field, &q.Quote)
	if v != "" {
		return v
	}
	switch field {
	case YTDReturn:
		return utils.ToStringF(q.YTDReturn)
	case QtrReturn:
		return utils.ToStringF(q.TrailingThreeMonthReturns)
	case QtrNavReturn:
		return utils.ToStringF(q.TrailingThreeMonthNavReturns)
	default:
		return ""
	}
}

// MapMutualFund maps mutual fund fields.
func MapMutualFund(field string, q *finance.MutualFund) string {
	v := MapQuote(field, &q.Quote)
	if v != "" {
		return v
	}
	switch field {
	case YTDReturn:
		return utils.ToStringF(q.YTDReturn)
	case QtrReturn:
		return utils.ToStringF(q.TrailingThreeMonthReturns)
	case QtrNavReturn:
		return utils.ToStringF(q.TrailingThreeMonthNavReturns)
	default:
		return ""
	}
}

// MapIndex maps index fields.
func MapIndex(field string, q *finance.Index) string {
	return MapQuote(field, &q.Quote)
}

// MapOption maps option fields.
func MapOption(field string, q *finance.Option) string {
	v := MapQuote(field, &q.Quote)
	if v != "" {
		return v
	}
	switch field {
	case Underlier:
		return q.UnderlyingSymbol
	case OpenInterest:
		return utils.ToString(q.OpenInterest)
	case ExpireDate:
		return utils.DateF(q.ExpireDate)
	case Strike:
		return utils.ToStringF(q.Strike)
	default:
		return ""
	}
}

// MapFuture maps future fields.
func MapFuture(field string, q *finance.Future) string {
	v := MapQuote(field, &q.Quote)
	if v != "" {
		return v
	}
	switch field {
	case Underlier:
		return q.UnderlyingSymbol
	case OpenInterest:
		return utils.ToString(q.OpenInterest)
	case ExpireDate:
		return utils.DateF(q.ExpireDate)
	case Strike:
		return utils.ToStringF(q.Strike)
	default:
		return ""
	}
}

// MapForex maps forex pair fields.
func MapForex(field string, q *finance.ForexPair) string {
	return MapQuote(field, &q.Quote)
}

// MapCrypto maps crypto pair fields.
func MapCrypto(field string, q *finance.CryptoPair) string {
	v := MapQuote(field, &q.Quote)
	if v != "" {
		return v
	}
	switch field {
	case Algorithm:
		return q.Algorithm
	case StartDate:
		return utils.DateF(q.StartDate)
	case MaxSupply:
		return utils.NumberF(q.MaxSupply)
	case Circulating:
		return utils.NumberF(q.CirculatingSupply)
	default:
		return ""
	}
}
