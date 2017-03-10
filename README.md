# go-commander
[![Build Status](https://travis-ci.org/WindomZ/go-commander.svg?branch=master)](https://travis-ci.org/WindomZ/go-commander)
![License](https://img.shields.io/badge/license-MIT-green.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/WindomZ/go-commander)](https://goreportcard.com/report/github.com/WindomZ/go-commander)

![v0.7.0](https://img.shields.io/badge/version-v0.7.0-orange.svg)
![status](https://img.shields.io/badge/status-beta-yellow.svg)

The solution for Go command-line interfaces, 
drive by <[docopt](https://github.com/docopt/docopt.go)>, 
inspired by <[commander](https://github.com/tj/commander.js)>

## Installation

To use commander in your Go code:

```go
import "github.com/WindomZ/go-commander"
```

To install commander according to your `$GOPATH`:

```bash
go get github.com/WindomZ/go-commander
```

## Examples

### [Quick example](https://github.com/WindomZ/go-commander/blob/master/examples/quick_example/quick_example.go)

Such as the following help message

```markdown
Usage:
  quick_example tcp <host> <port> [--timeout=<seconds>]
  quick_example serial <port> [--baud=9600] [--timeout=<seconds>]
  quick_example -h | --help
  quick_example --version
```

To coding with `go-commander` just like this:

```go
import "github.com/WindomZ/go-commander"

// quick_example
commander.Program.
	Command("quick_example").
	Version("0.1.1rc")

// quick_example tcp <host> <port> [--timeout=<seconds>]
commander.Program.
	Command("tcp <host> <port>").
	Option("--timeout=<seconds>")

// quick_example serial <port> [--baud=9600] [--timeout=<seconds>]
commander.Program.
	Command("serial <port>").
	Option("--baud=9600").
	Option("--timeout=<seconds>")
```

### [Counted example](https://github.com/WindomZ/go-commander/blob/master/examples/counted_example/counted_example.go)

Such as the following help message

```markdown
Usage:
  counted_example [-v...]
  counted_example go [go]
  counted_example (--path=<path>)...
  counted_example <file> <file>
  counted_example -h | --help
  counted_example --version
```

To coding with `go-commander` just like this:

```go
import "github.com/WindomZ/go-commander"

// counted_example -v...
commander.Program.
	Command("counted_example").
	Option("-v...")

// counted_example go [go]
commander.Program.
	Command("go [go]")

// counted_example (--path=<path>)...
commander.Program.
	LineOption("(--path=<path>)...")

// counted_example <file> <file>
commander.Program.
	LineArgument("<file> <file>")
```

### [Calculator example](https://github.com/WindomZ/go-commander/blob/master/examples/calculator_example/calculator_example.go)

Such as the following help message

```markdown
Simple calculator example

Usage:
  calculator_example <value> ( ( + | - | * | / ) <value> )...
  calculator_example <function> <value> [( , <value> )]...
  calculator_example -h | --help
  calculator_example --version

Examples:
  calculator_example 1 + 2 + 3 + 4 + 5
  calculator_example 1 + 2 '*' 3 / 4 - 5    # note quotes around '*'
  calculator_example sum 10 , 20 , 30 , 40
```

To coding with `go-commander` just like this:

```go
import "github.com/WindomZ/go-commander"

// calculator_example
commander.Program.
	Command("calculator_example").
	Version("0.0.1").
	Description("Simple calculator example")

// calculator_example <value> ( ( + | - | * | / ) <value> )...
commander.Program.
	LineArgument("<value> ( ( + | - | * | / ) <value> )...").
	Action(func(c *commander.Context) error {
		if c.Contain("<function>") {
			return nil
		}
		var result int
		values := c.Doc.GetStrings("<value>")
		for index, value := range values {
			if i, err := strconv.Atoi(value); err != nil {
			} else if index == 0 {
				result = i
			} else {
				switch c.Args.Get(index*2 - 1) {
				case "+":
					result += i
				case "-":
					result -= i
				case "*":
					result *= i
				case "/":
					result /= i
				}
			}
		}
		fmt.Println(c.Args.String(), "=", result)
		return nil
	})

// calculator_example <function> <value> [( , <value> )]...
commander.Program.
	LineArgument("<function> <value> [( , <value> )]...").
	Action(func(c *commander.Context) error {
		var result int
		switch c.Doc.GetString("<function>") {
		case "sum":
			values := c.Doc.GetStrings("<value>")
			for _, value := range values {
				if i, err := strconv.Atoi(value); err == nil {
					result += i
				}
			}
		}
		fmt.Println(c.Args.String(), "=", result)
		return nil
	})

// Examples: ...
commander.Program.
	Annotation(
		"Examples",
		[]string{
			"calculator_example 1 + 2 + 3 + 4 + 5",
			"calculator_example 1 + 2 '*' 3 / 4 - 5    # note quotes around '*'",
			"calculator_example sum 10 , 20 , 30 , 40",
		},
	)
```

Get the terminal output:
```bash
calculator_example 1 + 2 + 3 + 4 + 5
# output: 15

calculator_example 1 + 2 '*' 3 / 4 - 5
# output: -3

calculator_example sum 10 , 20 , 30 , 40
# output: 100
```

## License

The [MIT License](https://github.com/WindomZ/gitclone/blob/master/LICENSE)
