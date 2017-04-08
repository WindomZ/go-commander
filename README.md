# go-commander
[![Build Status](https://travis-ci.org/WindomZ/go-commander.svg?branch=master)](https://travis-ci.org/WindomZ/go-commander)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/WindomZ/go-commander)](https://goreportcard.com/report/github.com/WindomZ/go-commander)
[![Coverage Status](https://coveralls.io/repos/github/WindomZ/go-commander/badge.svg?branch=master)](https://coveralls.io/github/WindomZ/go-commander?branch=master)

The solution for building command shell programs, 
drive by <[docopt](https://github.com/docopt/docopt.go)>, 
inspired by <[commander.js](https://github.com/tj/commander.js)>

![v0.17.2](https://img.shields.io/badge/version-v0.17.2-yellow.svg)
![status](https://img.shields.io/badge/status-beta-yellow.svg)

## Features

- [x] Has all the features of [docopt](https://github.com/docopt/docopt.go).
- [x] Usage like [commander.js](https://github.com/tj/commander.js) as simple and readable.
- [x] Automatic generated a help message, easy to use, or advanced usage see [documents](http://docopt.org/) of docopt.
- [x] Automatically execute the correct action function, don't worry about conflict.
- [x] Can customize the action function, better with context.
- [x] Code colloquial use, from top to bottom.

## Installation

To install commander according to your `$GOPATH`:

```bash
go get github.com/WindomZ/go-commander
```

To use commander in your Go code:

```go
import "github.com/WindomZ/go-commander"
```

## Examples

### [Quick example](https://github.com/WindomZ/go-commander/blob/master/examples/quick_example/quick_example.go)

Such as the following help message

```markdown
Usage:
  quick_example tcp <host> <port> [--timeout=<seconds>]
  quick_example serial <port> [--baud=9600] [--timeout=<seconds>]
  quick_example -h|--help
  quick_example -v|--version

Options:
  -h --help     show help message
  -v --version  show version
```

To coding with `go-commander` just like this:

```go
import "github.com/WindomZ/go-commander"
...

// quick_example
commander.Program.
	Command("quick_example").
	Version("0.1.1rc")

// quick_example tcp <host> <port> [--timeout=<seconds>]
commander.Program.
	Command("tcp <host> <port>").
	Option("--timeout=<seconds>").
	Action(func() {
		fmt.Printf("tcp %s %s %s\n",
			commander.Program.MustString("<host>"),
			commander.Program.MustString("<port>"),
			commander.Program.MustString("--timeout"),
		)
	})

// quick_example serial <port> [--baud=9600] [--timeout=<seconds>]
commander.Program.
	Command("serial <port>").
	Option("--baud=9600").
	Option("--timeout=<seconds>").
	Action(func() {
		fmt.Printf("serial %s %s %s\n",
			commander.Program.MustString("<port>"),
			commander.Program.MustString("--baud"),
			commander.Program.MustString("--timeout"),
		)
	})

commander.Program.Parse()
```

Get the terminal output:

```bash
$ quick_example --version
# output: 0.1.1rc

$ quick_example tcp 127.0.0.1 1080 --timeout=110
# output: tcp 127.0.0.1 1080 110

$ quick_example serial 80 --baud=5800 --timeout=120
# output: serial 80 5800 120
```

### [Counted example](https://github.com/WindomZ/go-commander/blob/master/examples/counted_example/counted_example.go)

Such as the following help message

```markdown
Usage:
  counted_example -v...
  counted_example go [go]
  counted_example (--path=<path>)...
  counted_example <file> <file>
  counted_example -h|--help
  counted_example -v|--version

Options:
  -h --help     show help message
  -v --version  show version
```

To coding with `go-commander` just like this:

```go
import "github.com/WindomZ/go-commander"
...

// counted_example -v...
commander.Program.
	Command("counted_example").
	Option("-v...", "", func(c commander.Context) {
		fmt.Println("-v =", c.Get("-v"))
	})

// counted_example go [go]
commander.Program.
	Command("go [go]").
	Action(func(c commander.Context) {
		fmt.Println("go =", c.Get("go"))
	})

// counted_example (--path=<path>)...
commander.Program.
	Command("(--path=<path>)...", "", func(c commander.Context) {
		fmt.Printf("--path = %q\n",
			c.MustStrings("--path"))
	})

// counted_example <file> <file>
commander.Program.
	Command("<file> <file>", "", func(c commander.Context) {
		fmt.Printf("<file> = %q\n",
			c.MustStrings("<file>"))
	})

commander.Program.Parse()
```

Get the terminal output:

```bash
$ counted_example -vvvvvvvvvv
# output: -v = 10

$ counted_example go go
# output: go = 2

$ counted_example --path ./here --path ./there
# output: --path = ["./here" "./there"]

$ counted_example this.txt that.txt
# output: <file> = ["this.txt" "that.txt"]
```

### [Calculator example](https://github.com/WindomZ/go-commander/blob/master/examples/calculator_example/calculator_example.go)

Such as the following help message

```markdown
simple calculator example

Usage:
  calculator_example <value> ( ( + | - | * | / ) <value> )...
  calculator_example <function> <value> [( , <value> )]...
  calculator_example -h|--help
  calculator_example -v|--version

Options:
  -h --help     show help message
  -v --version  show version

Examples:
  calculator_example 1 + 2 + 3 + 4 + 5
  calculator_example 1 + 2 '*' 3 / 4 - 5    # note quotes around '*'
  calculator_example sum 10 , 20 , 30 , 40
```

To coding with `go-commander` just like this:

```go
import . "github.com/WindomZ/go-commander"
...

// calculator_example
Program.Command("calculator_example").
	Version("0.0.1").
	Description("Simple calculator example")

// calculator_example <value> ( ( + | - | * | / ) <value> )...
Program.Command("<value> ( ( + | - | * | / ) <value> )...", "", func() {
	var result int
	values := Program.MustStrings("<value>")
	for index, value := range values {
		if i, err := strconv.Atoi(value); err != nil {
		} else if index == 0 {
			result = i
		} else {
			switch Program.GetArg(index*2 - 1) {
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
	fmt.Println(Program.ArgsString(), "=", result)
})

// calculator_example <function> <value> [( , <value> )]...
Program.Command("<function> <value> [( , <value> )]...", "", func() {
	var result int
	switch Program.MustString("<function>") {
	case "sum":
		values := Program.MustStrings("<value>")
		for _, value := range values {
			if i, err := strconv.Atoi(value); err == nil {
				result += i
			}
		}
	}
	fmt.Println(Program.ArgsString(), "=", result)
})

// Examples: ...
Program.Annotation("Examples",
	[]string{
		"calculator_example 1 + 2 + 3 + 4 + 5",
		"calculator_example 1 + 2 '*' 3 / 4 - 5    # note quotes around '*'",
		"calculator_example sum 10 , 20 , 30 , 40",
	},
)

commander.Program.Parse()
```

Get the terminal output:

```bash
$ calculator_example 1 + 2 + 3 + 4 + 5
# output: 15

$ calculator_example 1 + 2 '*' 3 / 4 - 5
# output: -3

$ calculator_example sum 10 , 20 , 30 , 40
# output: 100
```

## License

The [MIT License](https://github.com/WindomZ/gitclone/blob/master/LICENSE)
