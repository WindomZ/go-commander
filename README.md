# go-commander
[![Build Status](https://travis-ci.org/WindomZ/go-commander.svg?branch=master)](https://travis-ci.org/WindomZ/go-commander)
![License](https://img.shields.io/badge/license-MIT-green.svg)

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

## Example

### Quick example

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

cmd := commander.NewCommander("quick_example").Version("0.1.1rc")
cmd.Command("tcp <host> <port>").
Option("--timeout=<seconds>")
cmd.Command("serial <port>").
Option("--baud=9600").
Option("--timeout=<seconds>")
```

### Counted example

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

cmd := commander.NewCommander("counted_example")
cmd.Option("-v...")
cmd.Command("go [go]")
cmd.LineOption("(--path=<path>)...")
cmd.LineArgument("<file> <file>")
```

## License

The [MIT License](https://github.com/WindomZ/gitclone/blob/master/LICENSE)
