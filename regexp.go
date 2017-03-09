package commander

import "regexp"

// RegexpCommand Regular screening out docopt commands
func RegexpCommand(str string) []string {
	return regexp.MustCompile(`[A-Za-z0-9_-]+`).FindAllString(
		regexp.MustCompile(`^[A-Za-z0-9_|\(\)\s-]+`).FindString(str), -1)
}

// RegexpArgument Regular screening out docopt arguments
func RegexpArgument(str string) []string {
	return regexp.MustCompile(`(?i:<|\[)[A-Za-z0-9_\[\]<>-]+\b(?i:>|])`).
		FindAllString(str, -1)
}

// RegexpOption Regular screening out docopt options
func RegexpOption(str string) []string {
	return regexp.MustCompile(`-{1,2}[A-Za-z0-9_-]+\b`).
		FindAllString(regexp.MustCompile(`(<|\[)[A-Za-z0-9_\[\]<>-]+\b(>|])`).
			ReplaceAllString(str, ""), -1)
}

// ContainCommand str contain docopt command format
func ContainCommand(str string) (ok bool) {
	ok, _ = regexp.MatchString(`[A-Za-z0-9_-]+`,
		regexp.MustCompile(`^[A-Za-z0-9_|\(\)\s-]+`).FindString(str))
	return
}

// ContainArgument str contain docopt argument format
func ContainArgument(str string) (ok bool) {
	ok, _ = regexp.MatchString(`(?i:<|\[)[A-Za-z0-9_\[\]<>-]+\b(?i:>|])`, str)
	return
}

// ContainOption str contain docopt option format
func ContainOption(str string) (ok bool) {
	ok, _ = regexp.MatchString(`-{1,2}[A-Za-z0-9_-]+\b`,
		regexp.MustCompile(`(<|\[)[A-Za-z0-9_\[\]<>-]+\b(>|])`).
			ReplaceAllString(str, ""))
	return
}
