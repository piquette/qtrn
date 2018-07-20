package quote

import (
	"fmt"

	finance "github.com/piquette/finance-go"
	"github.com/piquette/qtrn/utils"
)

const (
	// Quote fields.
	Symbol    = "Symbol"
	Market    = "Market"
	Time      = "Time"
	Last      = "Last"
	Change    = "Change"
	Vol       = "Volume"
	Bid       = "Bid"
	BidSize   = "Bid Size"
	Ask       = "Ask"
	AskSize   = "Ask Size"
	Open      = "Open"
	PrevClose = "Prev Close"
	Exchange  = "Exchange"
	DayHigh   = "Day High"
	DayLow    = "Day Low"
	YearHigh  = "52wk High"
	YearLow   = "52wk Low"
	FiftyMA   = "50D MA"
	THundMA   = "200D MA"
	ADV       = "Avg Daily Vol"
)

// FieldsQ returns the fields for a short plain quote.
func FieldsQ() (fields []string) {
	return []string{Symbol, Market, Time, Last, Change, Vol,
		Bid, BidSize, Ask, AskSize, Open, PrevClose}
}

// FieldsFullQ returns the fields for a full plain quote.
func FieldsFullQ() (fields []string) {
	return []string{Symbol, Market, Time, Last, Change, Vol,
		Bid, BidSize, Ask, AskSize, Open, PrevClose, Exchange,
		DayHigh, DayLow, YearHigh, YearLow, FiftyMA, THundMA, ADV}
}

// MapQ maps quote fields.
func MapQ(field string, q *finance.Quote) string {
	switch field {
	case Symbol:
		return q.Symbol
	case Market:
		return utils.MktStateF(q.MarketState)
	case Time:
		return utils.DateF(q)
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
	case ADV:
		return utils.NumberF(q.AverageDailyVolume10Day)
	default:
		return ""
	}
}
