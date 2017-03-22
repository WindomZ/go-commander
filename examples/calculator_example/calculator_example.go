package main

import (
	"fmt"
	. "github.com/WindomZ/go-commander"
	"strconv"
)

func main() {
	// ----------- go-commander -----------
	// calculator_example
	Program.Command("calculator_example").
		Version("0.0.1").
		Description("Simple calculator example")

	// calculator_example <value> ( ( + | - | * | / ) <value> )...
	Program.Command("<value> ( ( + | - | * | / ) <value> )...", "", func() {
		var result int
		values := Program.GetStrings("<value>")
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
		switch Program.GetString("<function>") {
		case "sum":
			values := Program.GetStrings("<value>")
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

	context, _ := Program.Parse()

	//fmt.Println(Program.HelpMessage()) // print help messages
	fmt.Println(context.String())

	fmt.Println("-------------")

	// ----------- docopt-go -----------
	usage := `Not a serious example.

Usage:
  calculator_example <value> ( ( + | - | * | / ) <value> )...
  calculator_example <function> <value> [( , <value> )]...
  calculator_example (-h | --help)

Examples:
  calculator_example 1 + 2 + 3 + 4 + 5
  calculator_example 1 + 2 '*' 3 / 4 - 5    # note quotes around '*'
  calculator_example sum 10 , 20 , 30 , 40

Options:
  -h, --help
`
	arguments, _ := Parse(usage, nil, true, "", false)

	//fmt.Println(usage) // print help messages
	fmt.Println(arguments)

	fmt.Println("===============================")
	fmt.Println()
}
