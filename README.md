[![Go Report Card](https://goreportcard.com/badge/github.com/luistm/banksaurus)](https://goreportcard.com/report/github.com/luistm/banksaurus)

## Banksaurus

When managing your finances you must ask yourself, where does your money go? 

This is a personal finance manager driven mainly by fun.

### Current status

This is a work in progress and is mainly a way to learn software stuff.

For now, it offers only a command line interface.

Currently it works only with [Caixa Geral de Depositos](https://www.cgd.pt) exported csv file.

### How to use

Login into your [Caixa Geral de Depositos](https://www.cgd.pt) and export the csv file with your account movements.

#### Available commands

```bash
$ bscli --help

    Your command line finance manager.

Usage:
	bscli -h | --help
	bscli report
	bscli report --input <file> [ --grouped ]
	bscli load --input <file>
	bscli seller change <id> --pretty <name>
	bscli seller show

Options:
	--grouped     The report is present grouped by seller
	--input       The path to the records list.
	--name        Specifies the name.
	-h --help     Show this screen.
```

### Setup

If you're interested in hacking or trying `banksaurus`, you can install via `go get`:

```bash
go get -u github.com/luistm/banksaurus/cmd/bscli
```

To run tests, execute the following:

```bash
make tests
````

For help about make commands just type:

```bash
make
```

### Feedback

Feedback is greatly appreciated.

### Contributing

Although this is mainly a way to learn Go, contributions are greatly appreciated.
