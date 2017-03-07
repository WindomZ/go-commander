package main

import (
	"fmt"
	"github.com/WindomZ/go-commander"
)

func main() {
	usage := `Usage:
  quick_example tcp <host> <port> [--timeout=<seconds>]
  quick_example serial <port> [--baud=9600] [--timeout=<seconds>]
  quick_example -h | --help | --version`

	arguments, _ := commander.Parse(usage, nil, true, "0.1.1rc", false)
	fmt.Println(arguments)

	cmd := commander.NewCommander("quick_example")

	cmd.Command("tcp <host> <port>").
		Option("--timeout <seconds>", "")

	cmd.Command("serial <port>").
		Option("--baud=9600", "").
		Option("--timeout <seconds>", "")

	fmt.Println(cmd.GetHelpMessage())

	arguments2, _ := cmd.Parse()
	fmt.Println(arguments2)

}
