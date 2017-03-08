package tests

import "github.com/WindomZ/go-commander"

func RegexpCommand(str string) []string {
	return commander.RegexpCommand(str)
}
func RegexpArgument(str string) []string {
	return commander.RegexpArgument(str)
}
func RegexpOption(str string) []string {
	return commander.RegexpOption(str)
}
