# go-bank-cli

Personal finance manager driven mainly by fun.

## Current status

This is a work in progress and is mainly a way to learn software stuff.


A command line interface is available, but may be broken from time to time.

A stable 1.0.0 version will be available in the near future.


A web interface will also be available in the future.


Currently it works only with [Caixa Geral de Depositos](https://www.cgd.pt) exported csv file.

### Available commands

```bash
bankcli

Usage:
  bankcli -h | --help
  bankcli report --input <file> [ --grouped ]
  bankcli load --input <file>
  bankcli category new <name>
  bankcli category show
  bankcli seller change <id> --pretty <name>
  bankcli seller new <name>
  bankcli seller show
```

## Setup

On macOS install the latest [dep](https://github.com/golang/dep) version with [Homebrew](https://brew.sh):

```bash
brew install dep
brew upgrade dep
```

If you're interested in hacking or trying `go-bank-cli`, you can install via `go get`:

```bash
go get -u github.com/luistm/go-cli-bank
```

To setup, run the following in the project root directory:

```bash
make deps
```

To run tests, execute the following:

```bash
make test
````

For help about make commands just type:

```bash
make
```

## Feedback

Feedback is greatly appreciated.

## Contributing

Although this is mainly a way to learn Go, contributions are greatly appreciated.
