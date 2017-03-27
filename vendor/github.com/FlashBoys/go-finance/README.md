# go-finance

[![GoDoc](https://godoc.org/github.com/FlashBoys/go-finance?status.svg)](https://godoc.org/github.com/FlashBoys/go-finance)
[![Build Status](https://travis-ci.org/FlashBoys/go-finance.svg?branch=master)](https://travis-ci.org/FlashBoys/go-finance) [![codecov.io](https://codecov.io/github/FlashBoys/go-finance/coverage.svg?branch=master)](https://codecov.io/github/FlashBoys/go-finance?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/Flashboys/go-finance)](https://goreportcard.com/report/github.com/Flashboys/go-finance)
[![License MIT](https://img.shields.io/npm/l/express.svg)](http://opensource.org/licenses/MIT)

![codecov.io](https://codecov.io/github/FlashBoys/go-finance/branch.svg?branch=master)

`go-finance` is a Go library for retrieving financial data for quantitative analysis.

To install go-finance, use the following command:

```
go get github.com/FlashBoys/go-finance
```


## Features

### Single security quotes

```go
package main

import (
	"fmt"

	"github.com/FlashBoys/go-finance"
)

func main() {
	// 15-min delayed full quote for Apple.
	q, err := finance.GetQuote("AAPL")
	if err == nil {
		fmt.Println(q)
	}
}
```

### Multiple securities quotes

```go
package main

import (
	"fmt"

	"github.com/FlashBoys/go-finance"
)

func main() {
	// 15-min delayed full quotes for Apple, Twitter, and Facebook.
	symbols := []string{"AAPL", "TWTR", "FB"}
	quotes, err := finance.GetQuotes(symbols)
	if err == nil {
		fmt.Println(quotes)
	}
}
```

### Currency pair quote

```go
package main

import (
	"fmt"

	"github.com/FlashBoys/go-finance"
)

func main() {
	// Predefined pair constants
	// e.g
	//
	// USDJPY
	// EURUSD
	// NZDUSD
	//
	pairquote, err := finance.GetCurrencyPairQuote(finance.USDJPY)
	if err == nil {
		fmt.Println(pairquote)
	}
}
```

### Quote history

```go
package main

import (
	"fmt"
	"time"

	"github.com/FlashBoys/go-finance"
)

func main() {
	// Set time frame to 1 month starting Jan. 1.
	start := finance.ParseDatetime("1/1/2017")
	end := finance.ParseDatetime("2/1/2017")

	// Request daily history for TWTR.
	// IntervalDaily OR IntervalWeekly OR IntervalMonthly are supported.
	bars, err := finance.GetHistory("TWTR", start, end, finance.IntervalDaily)
	if err == nil {
		fmt.Println(bars)
	}
}
```

### Dividend/Split event history

```go
package main

import (
	"fmt"
	"time"

	"github.com/FlashBoys/go-finance"
)

func main() {
	// Set time range from Jan 2010 up to the current date.
	// This example will return a slice of both dividends and splits.
	start, _ := finance.ParseDatetime("1/1/2010")
	end := finance.NewDatetime(time.Now())

	// Request event history for AAPL.
	events, err := finance.GetEventHistory("AAPL", start, end)
	if err == nil {
		fmt.Println(events)
	}
}
```

### Symbols download

```go
package main

import (
	"fmt"

	"github.com/FlashBoys/go-finance"
)

func main() {
	// Request all BATS symbols.
	symbols, err := finance.GetUSEquitySymbols()
	if err == nil {
		fmt.Println(symbols)
	}
}

```

### Options chains

```go
package main

import (
	"fmt"

	"github.com/FlashBoys/go-finance"
)

func main() {
	// Fetches the available expiration dates.
	c, err := finance.NewCycle("AAPL")
	if err != nil {
		panic(err)
	}

	// Some examples - see docs for full details.

	// Fetches the chain for the front month.
	calls, puts, err := c.GetFrontMonth()
	if err == nil {
		panic(err)
	}
	fmt.Println(calls)
	fmt.Println(puts)

	// Fetches the chain for the specified expiration date.
	calls, puts, err := c.GetChainForExpiration(chain.Expirations[1])
	if err == nil {
		panic(err)
	}
	fmt.Println(calls)
	fmt.Println(puts)

	// Fetches calls for the specified expiration date.
	calls, err := c.GetCallsForExpiration(chain.Expirations[1])
	if err == nil {
		panic(err)
	}
	fmt.Println(calls)
}

```


## Intentions

The primary technical tenants of this project are:

  * Make financial data easy and fun to work with in Go-lang.
  * Abstract the burden of non-sexy model serialization away from the end-user.
  * Provide a mature framework where the end-user needs only be concerned with analysis instead of data sourcing.

There are several applications for this library. It's intentions are to be conducive to the following activities:

  * Quantitative financial analysis in Golang.
  * Academic study/comparison in a clean, easy language.
  * Algorithmic/Statistical-based strategy implementation.

## To-do

- [ ] Add greeks calculations to options data
- [ ] Key stats (full profile) for securities

## Contributing

If you find this repo helpful, please give it a star! If you wish to discuss changes to it, please open an issue. This project is not as mature as it could be, and financial projects in Golang are in drastic need of some basic helpful dependencies.

## Similar Projects

  * [pandas datareader](https://github.com/pydata/pandas-datareader) (Python) wide-spread use in academia.
  * [yahoofinance-api](https://github.com/sstrickx/yahoofinance-api) (Java) most popular java library for this purpose.
  * [quantmod](http://www.quantmod.com/) (R) a package for development/testing/deployment of quant strategy.
