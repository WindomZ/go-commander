package main

import (
	"fmt"
	"github.com/WindomZ/go-commander"
)

func main() {
	// ----------- go-commander -----------
	// new counted_example
	cmd := commander.NewCommander("counted_example")
	// counted_example -v...
	cmd.Option("-v...")
	// counted_example go [go]
	cmd.Command("go [go]")
	// counted_example (--path=<path>)...
	cmd.LineOption("(--path=<path>)...")
	// counted_example <file> <file>
	cmd.LineArgument("<file> <file>")

	arguments2, _ := cmd.Parse()

	//fmt.Println(cmd.GetHelpMessage())
	fmt.Println(arguments2.Doc)

	fmt.Println("-------------")

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
