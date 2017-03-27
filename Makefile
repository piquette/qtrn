
SHELL = /bin/sh
.SUFFIXES:
.SUFFIXES: .go

get:
	go fmt && go clean && go build && ./qtrn get AAPL TWTR
	go fmt && go clean && go build && ./qtrn get AAPL TWTR -f


test:
	curl https://glide.sh/get | sh
	go test $(glide novendor)
