package main

import (
	"fmt"
	"github.com/WindomZ/go-commander"
)

func main() {
	// ----------- go-commander -----------
	// quick_example
	commander.Program.
		Version("0.1.1rc")

	// quick_example tcp <host> <port> [--timeout=<seconds>]
	commander.Program.
		Command("tcp <host> <port>").
		Option("--timeout=<seconds>").
		Action(func() {
			fmt.Printf("tcp %s %s %s\n",
				commander.Program.GetString("<host>"),
				commander.Program.GetString("<port>"),
				commander.Program.GetString("--timeout"),
			)
		})

	// quick_example serial <port> [--baud=9600] [--timeout=<seconds>]
	commander.Program.
		Command("serial <port>").
		Option("--baud=9600").
		Option("--timeout=<seconds>").
		Action(func() {
			fmt.Printf("serial %s %s %s\n",
				commander.Program.GetString("<port>"),
				commander.Program.GetString("--baud"),
				commander.Program.GetString("--timeout"),
			)
		})

	context, _ := commander.Program.Parse()

	//fmt.Println(commander.Program.HelpMessage()) // print help messages
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
