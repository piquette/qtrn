[![Go Report Card](https://goreportcard.com/badge/github.com/FlashBoys/qtrn)](https://goreportcard.com/badge/github.com/FlashBoys/qtrn)
[![Build Status](https://travis-ci.org/FlashBoys/qtrn.svg?branch=master)](https://travis-ci.org/FlashBoys/qtrn)
[![GoDoc](https://godoc.org/github.com/FlashBoys/qtrn?status.svg)](https://godoc.org/github.com/FlashBoys/qtrn)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

# :zap: qtrn :zap:
The official cli tool for making financial markets analysis as fast as you are.

## Commands
The current available commands are:
* `quote`  Prints tables of stock quotes to the current shell :ledger:
* `write`  Writes a csv of stock market data :open_file_folder:
* `chart`  Prints stock charts to the current shell (still in beta) :chart_with_upwards_trend:

## Installation
In order to use this awesome tool, you'll need to get it on your machine!

### Download Distribution
1. Head over to the official [releases page](https://github.com/FlashBoys/qtrn/releases)
2. Determine the appropriate distribution for your operating system (mac | windows | linux)
3. Download and untar the distribution. Shortcut for macs:
```
curl -sL https://github.com/FlashBoys/qtrn/releases/download/v0.5.1/qtrn-0.5.1-darwin-amd64.tgz | tar zx
```
4. Move the binary into your local `$PATH`.
5. Run `qtrn help`.

### From Source
qtrn is built in Go. To get started with Go, head to the official instructions [here](https://golang.org/doc/install). A common way to install go is through pacakge manager [homebrew](https://brew.sh/) using the command:

```
brew install go
```

You will also need Glide, a Go dependency management tool. It can be installed simply:

```
brew install glide
```

Next, you'll want to `go get` the source repo, like this:

```
go get github.com/FlashBoys/qtrn
```

If your `$GOPATH` is configured, and git is setup to know your credentials, in a few moments the command should complete with no output. The repository will exist under `$GOPATH/src/github.com/FlashBoys`.


## Usage
The preferred way to build qtrn for development is using `make`. Run `make dev && qtrn help` which builds the project, moves the binary into `$GOPATH/bin/` and lists possible commands.

### Contributing
1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request :)
