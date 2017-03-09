package commander

import "regexp"

func RegexpCommand(str string) []string {
	return regexp.MustCompile(`[A-Za-z0-9_-]+`).FindAllString(
		regexp.MustCompile(`^[A-Za-z0-9_|\(\)\s-]+`).FindString(str), -1)
}

func RegexpArgument(str string) []string {
	return regexp.MustCompile(`(?i:<|\[)[A-Za-z0-9_\[\]<>-]+\b(?i:>|])`).
		FindAllString(str, -1)
}

func RegexpOption(str string) []string {
	return regexp.MustCompile(`-{1,2}[A-Za-z0-9_-]+\b`).
		FindAllString(regexp.MustCompile(`(<|\[)[A-Za-z0-9_\[\]<>-]+\b(>|])`).
			ReplaceAllString(str, ""), -1)
}

func ContainCommand(str string) (ok bool) {
	ok, _ = regexp.MatchString(`[A-Za-z0-9_-]+`,
		regexp.MustCompile(`^[A-Za-z0-9_|\(\)\s-]+`).FindString(str))
	return
}

func ContainArgument(str string) (ok bool) {
	ok, _ = regexp.MatchString(`(?i:<|\[)[A-Za-z0-9_\[\]<>-]+\b(?i:>|])`, str)
	return
}

func ContainOption(str string) (ok bool) {
	ok, _ = regexp.MatchString(`-{1,2}[A-Za-z0-9_-]+\b`,
		regexp.MustCompile(`(<|\[)[A-Za-z0-9_\[\]<>-]+\b(>|])`).
			ReplaceAllString(str, ""))
	return
}
