package main

import (
	"fmt"
	"github.com/WindomZ/go-commander"
	"strconv"
)

func main() {
	// ----------- go-commander -----------
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

	context, _ := commander.Program.Parse()

	//fmt.Println(commander.Program.HelpMessage()) // print help messages
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
	arguments, _ := commander.Parse(usage, nil, true, "", false)

	//fmt.Println(usage) // print help messages
	fmt.Println(arguments)

	fmt.Println("===============================")
	fmt.Println()
}
