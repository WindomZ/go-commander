package main

import (
	"fmt"
	"github.com/WindomZ/go-commander"
)

func main() {
	// ----------- go-commander -----------
	// counted_example -v...
	commander.Program.
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

	//fmt.Println(commander.Program.HelpMessage()) // print help messages

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
