[![Go Report Card](https://goreportcard.com/badge/github.com/piquette/qtrn)](https://goreportcard.com/badge/github.com/piquette/qtrn)
[![Build Status](https://travis-ci.org/piquette/qtrn.svg?branch=master)](https://travis-ci.org/piquette/qtrn)
[![GoDoc](https://godoc.org/github.com/piquette/qtrn?status.svg)](https://godoc.org/github.com/piquette/qtrn)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

# :zap: qtrn :zap:
The official cli tool for making financial markets analysis as fast as you are.

## Commands
The current available commands are:
* `equity` - prints tables of equity quotes to the current shell

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
curl -sL https://github.com/piquette/qtrn/releases/download/v0.0.1/qtrn_0.0.2_darwin_amd64.tar.gz | tar zx
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

### Example
The

The preferred way to build qtrn for development is using `make`. Run `make dev && qtrn help` which builds the project, moves the binary into `$GOPATH/bin/` and lists possible commands.

### Contributing
1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request :)
