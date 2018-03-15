# banksaurus

Personal finance manager driven mainly by fun.

## Current status

This is a work in progress and is mainly a way to learn software stuff.

For now, it offers only a command line interface.

Currently it works only with [Caixa Geral de Depositos](https://www.cgd.pt) exported csv file.

## How to use

Login into your [Caixa Geral de Depositos](https://www.cgd.pt) and export the csv file with your account movements.

### Available commands

```bash
$ banksaurus --help

    Your command line finance manager.

Usage:
	banksaurus -h | --help
	banksaurus report --input <file> [ --grouped ]
	banksaurus load --input <file>
	banksaurus seller change <id> --pretty <name>
	banksaurus seller new <name>
	banksaurus seller show

Options:
	--grouped     The report is present grouped by seller
	--input       The path to the records list.
	--name        Specifies the name.
	-h --help     Show this screen.
```

## Setup

If you're interested in hacking or trying `banksaurus`, you can install via `go get`:

```bash
go get -u github.com/luistm/banksaurus/cmd/banksaurus
```

On macOS install the latest [dep](https://github.com/golang/dep) version with [Homebrew](https://brew.sh):

```bash
brew install dep
brew upgrade dep
```

Change directory to your go path:

```bash
cd $GOPATH/src/github.com/luistm/banksaurus
```

To setup, run the following in the project root directory:

```bash
make deps
```

To run tests, execute the following:

```bash
make tests
````

For help about make commands just type:

```bash
make
```

## Feedback

Feedback is greatly appreciated.

## Contributing

Although this is mainly a way to learn Go, contributions are greatly appreciated.
