package main

import (
	"fmt"
	"github.com/WindomZ/go-commander"
)

func main() {
	// ----------- go-commander -----------
	// counted_example -v...
	commander.Program.
		Command("counted_example").
		Option("-v...", "", func() {
			fmt.Println("-v =", commander.Program.Get("-v"))
		})

	// counted_example go [go]
	commander.Program.
		Command("go [go]").
		Action(func(c commander.Context) {
			fmt.Println("go =", c.Get("go"))
		})

	// counted_example (--path=<path>)...
	commander.Program.
		LineOption("(--path=<path>)...", "", func() {
			fmt.Printf("--path = %q\n",
				commander.Program.GetStrings("--path"))
		})

	// counted_example <file> <file>
	commander.Program.
		LineArgument("<file> <file>").
		Action(func(c commander.Context) {
			fmt.Printf("<file> = %q\n",
				commander.Program.GetStrings("<file>"))
		})

	context, _ := commander.Program.Parse()

	//fmt.Println(commander.Program.HelpMessage()) // print help messages
	fmt.Println(context.String())

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

	//fmt.Println(usage) // print help messages
	fmt.Println(arguments)

	fmt.Println("===============================")
}
