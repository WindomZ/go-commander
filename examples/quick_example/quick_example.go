package main

import (
	"fmt"
	"github.com/WindomZ/go-commander"
)

func main() {
	// ----------- go-commander -----------
	// quick_example
	commander.Program.
		Command("quick_example").
		Version("0.1.1rc")

	// quick_example tcp <host> <port> [--timeout=<seconds>]
	commander.Program.
		Command("tcp <host> <port>").
		Option("--timeout=<seconds>")

	// quick_example serial <port> [--baud=9600] [--timeout=<seconds>]
	commander.Program.
		Command("serial <port>").
		Option("--baud=9600").
		Option("--timeout=<seconds>")

	context, _ := commander.Program.Parse()

	//fmt.Println(commander.Program.GetHelpMessage()) // print help messages
	fmt.Println(context.String())

	fmt.Println("-------------")

	// ----------- docopt-go -----------
	usage := `Usage:
  quick_example tcp <host> <port> [--timeout=<seconds>]
  quick_example serial <port> [--baud=9600] [--timeout=<seconds>]
  quick_example -h | --help | --version`

	arguments, _ := commander.Parse(usage, nil, true, "0.1.1rc", false)

	//fmt.Println(usage) // print help messages
	fmt.Println(arguments)

	fmt.Println("===============================")
}
