package commander

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestFormat_DescriptionLine(t *testing.T) {
	assert.Equal(t,
		formatDescriptionLine("a", "desc...", 2, 5, false),
		"a    desc...")
	assert.Equal(t,
		formatDescriptionLine("abcdef", "desc...", 2, 5, false),
		"abcdef  desc...")

	assert.Equal(t,
		formatDescriptionLine("a", "desc...", 2, 5, true),
		"a    desc...")
	assert.Equal(t,
		formatDescriptionLine("abcde", "desc...", 2, 5, true),
		"abcde\n     desc...")
	assert.Equal(t,
		formatDescriptionLine("abcdef", "desc...", 2, 5, true),
		"abcdef\n     desc...")
}
