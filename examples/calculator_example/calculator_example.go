package main

import (
	"fmt"
	"github.com/WindomZ/go-commander"
)

func main() {
	// ----------- go-commander -----------
	cmd := commander.NewCommander("calculator_example")
	cmd.LineArgument("<value> ( ( + | - | * | / ) <value> )...")
	cmd.LineArgument("<function> <value> [( , <value> )]...")
	cmd.Annotation("Examples", []string{
		"calculator_example 1 + 2 + 3 + 4 + 5",
		"calculator_example 1 + 2 '*' 3 / 4 - 5    # note quotes around '*'",
		"calculator_example sum 10 , 20 , 30 , 40",
	})
	arguments2, _ := cmd.Parse()

	fmt.Println(cmd.GetHelpMessage())
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

	fmt.Println(usage)
	fmt.Println(arguments)

	fmt.Println("===============================")
}
