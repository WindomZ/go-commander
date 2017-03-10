package commander

import (
	"github.com/WindomZ/testify/assert"
	"strconv"
	"testing"
)

func TestCommander_Ping(t *testing.T) {
	var sum int
	var host string

	Program.Command("test").
		Version("0.0.1").
		Description("This is a test cli.")

	Program.Command("add <x> <y>").
		Description("addition operation").
		Action(func(c *Context) error {
			x, _ := c.Doc.GetInt("<x>")
			y, _ := c.Doc.GetInt("<y>")
			sum = x + y
			return nil
		})

	Program.Command("ping <host>").
		Action(func(c *Context) error {
			host = c.Doc.GetString("<host>")
			return nil
		}).
		Option("--timeout=<seconds>",
			"",
			func(c *Context) error {
				seconds := c.Doc.GetString("<seconds>")

				t.Log("seconds =", seconds)

				assert.Equal(t, seconds, 0)
				return nil
			},
		)

	if _, err := Program.Parse([]string{"test", "add", "10", "20"}); err != nil {
		t.Fatal(err)
	}
	if _, err := Program.Parse([]string{"test", "ping", "127.0.0.1"}); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, sum, 30)
	assert.Equal(t, host, "127.0.0.1")
}

func TestCommander_Calculator(t *testing.T) {
	var result int

	Program.Command("calculator_example").
		Version("0.0.1").
		Description("simple calculator example")

	Program.LineArgument("<value> ( ( + | - | * | / ) <value> )...").
		Action(func(c *Context) error {
			if c.Contain("<function>") {
				return nil
			}
			values := c.Doc.GetStrings("<value>")
			for index, value := range values {
				if i, err := strconv.Atoi(value); err != nil {
				} else if index == 0 {
					result = i
				} else {
					switch c.Args.Get(index*2 - 1) {
					case "+":
						result += i
					case "-":
						result -= i
					case "*":
						result *= i
					case "/":
						result /= i
					}
				}
			}
			return nil
		})

	Program.LineArgument("<function> <value> [( , <value> )]...").
		Action(func(c *Context) error {
			result = 0
			switch c.Doc.GetString("<function>") {
			case "sum":
				values := c.Doc.GetStrings("<value>")
				for _, value := range values {
					if i, err := strconv.Atoi(value); err == nil {
						result += i
					}
				}
			}
			return nil
		})

	if _, err := Program.Parse([]string{"calculator_example",
		"1", "+", "2", "+", "3", "+", "4", "+", "5"}); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, result, 15)

	if _, err := Program.Parse([]string{"calculator_example",
		"1", "+", "2", "*", "3", "/", "4", "-", "5"}); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, result, -3)

	if _, err := Program.Parse([]string{"calculator_example",
		"sum", "10", ",", "20", ",", "30", ",", "40"}); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, result, 100)
}
