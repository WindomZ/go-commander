package commander

import (
	"fmt"
	"strconv"
)

var Format Formatter = newDefaultFormat()

type Formatter interface {
	Description(title, desc string) string
}

type _Format struct {
	MinSpace int
	MaxSpace int
	Line     bool
}

func newFormat(minSpace, maxSpace int, line bool) *_Format {
	return &_Format{
		MinSpace: minSpace,
		MaxSpace: maxSpace,
		Line:     line,
	}
}

func newDefaultFormat() *_Format {
	return newFormat(2, 14, true)
}

// Description format a line description to symmetric string
// title and desc are shown the content by default format.
func (f _Format) Description(title, desc string) string {
	return formatDescription(title, desc, 2, 14, true)
}

// formatDescription format a line description to symmetric string
// title and desc are shown the content,
// minSpace and maxSpace are indentation range numbers.
func formatDescription(title, desc string, minSpace, maxSpace int, line bool) string {
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
