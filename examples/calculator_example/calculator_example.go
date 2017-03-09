package main

import (
	"fmt"
	"github.com/WindomZ/go-commander"
	"strconv"
)

func main() {
	// ----------- go-commander -----------
	// new calculator_example
	cmd := commander.NewCommander("calculator_example").
		Version("0.0.1").
		Description("simple calculator example")

	// calculator_example <value> ( ( + | - | * | / ) <value> )...
	cmd.LineArgument("<value> ( ( + | - | * | / ) <value> )...").
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
			fmt.Println(result)
			return nil
		})

	// calculator_example <function> <value> [( , <value> )]...
	cmd.LineArgument("<function> <value> [( , <value> )]...").
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
			fmt.Println(result)
			return nil
		})

	// Examples: ...
	cmd.Annotation("Examples", []string{
		"calculator_example 1 + 2 + 3 + 4 + 5",
		"calculator_example 1 + 2 '*' 3 / 4 - 5    # note quotes around '*'",
		"calculator_example sum 10 , 20 , 30 , 40",
	})

	arguments2, _ := cmd.Parse()

	//fmt.Println(cmd.GetHelpMessage())
	fmt.Println(arguments2.Doc)

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

	//fmt.Println(usage)
	fmt.Println(arguments)

	fmt.Println("===============================")
}
