package main

import (
	"fmt"
	"github.com/WindomZ/go-commander"
)

func main() {
	// ----------- go-commander -----------
	cmd := commander.NewCommander("naval_fate").
		Description("Naval Fate.")

	ship := cmd.Command("ship")
	ship.Command("new <name>...")
	ship.LineArgument("<name>").
		Command("move <x> <y>").
		Option("--speed=<kn>", "Speed in knots", nil, 10)
	ship.Command("shoot <x> <y>")

	mine := cmd.Command("mine")
	mine.Command("(set|remove) <x> <y>").
		Option("--moored|--drifting")

	arguments2, _ := cmd.Parse()

	//fmt.Println(cmd.GetHelpMessage())
	fmt.Println(arguments2.Doc)

	fmt.Println("-------------")

	// ----------- docopt-go -----------
	usage := `Naval Fate.

Usage:
  naval_fate ship new <name>...
  naval_fate ship <name> move <x> <y> [--speed=<kn>]
  naval_fate ship shoot <x> <y>
  naval_fate mine (set|remove) <x> <y> [--moored|--drifting]
  naval_fate -h | --help
  naval_fate --version

Options:
  -h --help     Show this screen.
  --version     Show version.
  --speed=<kn>  Speed in knots [default: 10].
  --moored      Moored (anchored) mine.
  --drifting    Drifting mine.`

	arguments, _ := commander.Parse(usage, nil, true, "Naval Fate 2.0", false)
	//fmt.Println(usage)
	fmt.Println(arguments)

	fmt.Println("===============================")
}
