package main

import (
	"fmt"
	"github.com/WindomZ/go-commander"
)

func main() {
	// ----------- go-commander -----------
	commander.Program.
		Description("Naval Fate.")

	ship := commander.Program.
		Command("ship")
	ship.Command("new <name>...")
	ship.Command("<name>").
		Command("move <x> <y>").
		Option("--speed=<kn>", "Speed in knots", nil, 10)
	ship.Command("shoot <x> <y>")

	mine := commander.Program.
		Command("mine")
	mine.Command("(set|remove) <x> <y>").
		Option("--moored|--drifting")

	commander.Program.Parse()

	//fmt.Println(commander.Program.HelpMessage()) // print help messages

	fmt.Println("-------------")

	// ----------- docopt-go -----------
	usage := `Naval Fate.

Usage:
  naval_fate_example ship new <name>...
  naval_fate_example ship <name> move <x> <y> [--speed=<kn>]
  naval_fate_example ship shoot <x> <y>
  naval_fate_example mine (set|remove) <x> <y> [--moored|--drifting]
  naval_fate_example -h | --help
  naval_fate_example --version

Options:
  -h --help     Show this screen.
  --version     Show version.
  --speed=<kn>  Speed in knots [default: 10].
  --moored      Moored (anchored) mine.
  --drifting    Drifting mine.`

	arguments, _ := commander.Parse(usage, nil, true, "Naval Fate 2.0", false)
	//fmt.Println(usage) // print help messages
	fmt.Println(arguments)

	fmt.Println("===============================")
}
