package commander

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestNewCommander(t *testing.T) {
	c := NewCommander("hello")
	assert.NotEmpty(t, c)
}

func TestCommander_1(t *testing.T) {
	var sum int
	var host string

	cmd := NewCommander("test").
		Version("0.0.1").
		Description("This is a test cli.")
	cmd.Command("add <x> <y>").
		Description("addition operation").
		Action(func(c *Context) Result {
			x, _ := c.Doc.GetInt("<x>")
			y, _ := c.Doc.GetInt("<y>")
			sum = x + y
			return nil
		})
	cmd.Command("ping <host>").
		Action(func(c *Context) Result {
			host = c.Doc.GetString("<host>")
			return nil
		}).
		Option("--timeout=<seconds>",
			"",
			func(c *Context) Result {
				seconds := c.Doc.GetString("<seconds>")

				t.Log("seconds =", seconds)

				assert.Equal(t, seconds, 0)
				return nil
			},
		)

	if _, err := cmd.Parse([]string{"test", "add", "10", "20"}); err != nil {
		t.Fatal(err)
	}
	if _, err := cmd.Parse([]string{"test", "ping", "127.0.0.1"}); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, sum, 30)
	assert.Equal(t, host, "127.0.0.1")
}
