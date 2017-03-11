package commander

import "regexp"

// regexpCommand Regular screening out docopt commands
func regexpCommand(str string) []string {
	return regexp.MustCompile(`[A-Za-z0-9_-]+`).FindAllString(
		regexp.MustCompile(`^[A-Za-z0-9_|\(\)\s-]+`).FindString(str), -1)
}

// regexpArgument Regular screening out docopt arguments
func regexpArgument(str string) []string {
	return regexp.MustCompile(`(?i:<|\[)[A-Za-z0-9_\[\]<>-]+\b(?i:>|])`).
		FindAllString(str, -1)
}

// regexpOption Regular screening out docopt options
func regexpOption(str string) []string {
	return regexp.MustCompile(`-{1,2}[A-Za-z0-9_-]+\b`).
		FindAllString(regexp.MustCompile(`(<|\[)[A-Za-z0-9_\[\]<>-]+\b(>|])`).
			ReplaceAllString(str, ""), -1)
}

// containCommand str contain docopt command format
func containCommand(str string) (ok bool) {
	ok, _ = regexp.MatchString(`[A-Za-z0-9_-]+`,
		regexp.MustCompile(`^[A-Za-z0-9_|\(\)\s-]+`).FindString(str))
	return
}

// containArgument str contain docopt argument format
func containArgument(str string) (ok bool) {
	ok, _ = regexp.MatchString(`(?i:<|\[)[A-Za-z0-9_\[\]<>-]+\b(?i:>|])`, str)
	return
}

// containOption str contain docopt option format
func containOption(str string) (ok bool) {
	ok, _ = regexp.MatchString(`-{1,2}[A-Za-z0-9_-]+\b`,
		regexp.MustCompile(`(<|\[)[A-Za-z0-9_\[\]<>-]+\b(>|])`).
			ReplaceAllString(str, ""))
	return
}

//// isCommand str contain docopt command format
//func isCommand(str string) bool {
//	return !isArgument(str) && !isOption(str)
//}
//
//// isArgument str contain docopt command format
//func isArgument(str string) (ok bool) {
//	ok, _ = regexp.MatchString(`^<[A-Za-z0-9_-]+>$`, str)
//	return
//}
//
//// isOption str contain docopt command format
//func isOption(str string) (ok bool) {
//	ok, _ = regexp.MatchString(`^-{1,2}[A-Za-z0-9_-]+$`, str)
//	return
//}
