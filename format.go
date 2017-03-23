package commander

import (
	"fmt"
	"strconv"
)

func formatDescriptionLine(title, desc string, minSpace, maxSpace int, line bool) string {
	if minSpace < 0 {
		minSpace = 0
	}
	if maxSpace < minSpace {
		maxSpace = minSpace
	}
	var pattern string
	if len(title) > (maxSpace - minSpace) {
		if line {
			pattern = "%s\n%" + strconv.Itoa(len(desc)+maxSpace) + "s"
		} else {
			pattern = "%s%" + strconv.Itoa(len(desc)+minSpace) + "s"
		}
	} else {
		pattern = "%-" + strconv.Itoa(maxSpace) + "s%s"
	}
	return fmt.Sprintf(pattern, title, desc)
}
