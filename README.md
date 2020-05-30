[![Go Report Card](https://goreportcard.com/badge/github.com/piquette/qtrn)](https://goreportcard.com/badge/github.com/piquette/qtrn)
[![Build Status](https://travis-ci.org/piquette/qtrn.svg?branch=master)](https://travis-ci.org/piquette/qtrn)
[![GoDoc](https://godoc.org/github.com/piquette/qtrn?status.svg)](https://godoc.org/github.com/piquette/qtrn)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

# :zap: qtrn :zap:
The official cli tool for making financial markets analysis as fast as you are.

Pronounced "quote-tron" as a throwback to those awesome financial terminals of the 80's. This project is intended as a living example of the capabilities of the [finance-go] library.

## Commands
The current available commands are:
* `quote` - prints tables of quotes to the current shell
* `options` - prints tables of options contract quotes to the current shell
* `write` - writes tables of quotes/history to csv files
* `chart` - (beta) prints a nifty sparkline chart of security for various time frames

## Installation
In order to use this awesome tool, you'll need to get it on your machine!

### From Homebrew
If you're on macOS, the easiest way to get qtrn is through the homebrew tap.
```
brew tap piquette/qtrn
brew install qtrn
```

### From Release
1. Head over to the official [releases page](https://github.com/piquette/qtrn/releases)
2. Determine the appropriate distribution for your operating system (mac | windows | linux)
3. Download and untar the distribution. Shortcut for macs:
```
curl -sL https://github.com/piquette/qtrn/releases/download/v0.0.9/qtrn_0.0.9_darwin_amd64.tar.gz | tar zx
```
4. Move the binary into your local `$PATH`.
5. Run `qtrn help`.

### From Source
qtrn is built in Go. `go get` the source repo, like this:

```
go get github.com/piquette/qtrn
```

If your `$GOPATH` is configured, and git is setup to know your credentials, in a few moments the command should complete with no output. The repository will exist under `$GOPATH/src/github.com/piquette`.

## Usage
Run the command `qtrn` in your shell for the list of possible commands.

### Contributing
The preferred way to build qtrn for development is using `make`. Run `make build && qtrn help`.

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request :)


## Release

Release builds are generated with [goreleaser]. Make sure you have the software
and a `GITHUB_TOKEN`: set in your env.

``` sh
go get -u github.com/goreleaser/goreleaser
export GITHUB_TOKEN=...
```

Commit changes and tag `HEAD`:

``` sh
git tag v[NEW_VERSION_NUMBER]
git push origin --tags
```

Then run goreleaser and you're done! Check [releases] (it also pushes to the
Homebrew tap).

``` sh
goreleaser --rm-dist
```

[goreleaser]: https://github.com/goreleaser/goreleaser
[releases]: https://github.com/piquette/qtrn/releases
[finance-go]: https://github.com/piquette/finance-go
