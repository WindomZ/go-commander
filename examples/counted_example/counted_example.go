package main

import (
	"fmt"
	"github.com/WindomZ/go-commander"
)

func main() {
	// ----------- go-commander -----------
	cmd := commander.NewCommander("counted_example")
	cmd.Option("-v...")
	cmd.Command("go [go]")
	cmd.LineOption("(--path=<path>)...")
	cmd.LineArgument("<file> <file>")
	arguments2, _ := cmd.Parse()

	//fmt.Println(cmd.GetHelpMessage())
	fmt.Println(arguments2.Doc)

	fmt.Println("------------------------------")

	// ----------- docopt-go -----------
	usage := `Usage: counted_example --help
       counted_example -v...
       counted_example go [go]
       counted_example (--path=<path>)...
       counted_example <file> <file>

Try: counted_example -vvvvvvvvvv
     counted_example go go
     counted_example --path ./here --path ./there
     counted_example this.txt that.txt`

	arguments, _ := commander.Parse(usage, nil, true, "", false)

	//fmt.Println(usage)
	fmt.Println(arguments)

	fmt.Println("===============================")
}
