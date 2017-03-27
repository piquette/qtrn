
SHELL = /bin/sh
.SUFFIXES:
.SUFFIXES: .go

get:
	go fmt && go clean && go build && ./qtrn get AAPL TWTR
	go fmt && go clean && go build && ./qtrn get AAPL TWTR -f


test:
	go test $(glide novendor)
