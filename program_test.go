package commander

import (
	"github.com/WindomZ/testify/assert"
	"strconv"
	"testing"
)

func TestProgram_Ping(t *testing.T) {
	var sum int
	var host string

	Program.Command("test").
		Version("0.0.1").
		Description("This is a test cli.")

	Program.Command("add <x> <y>").
		Description("addition operation").
		Action(func() error {
			x, _ := Program.GetInt("<x>")
			y, _ := Program.GetInt("<y>")
			sum = x + y
			return nil
		})

	Program.Command("ping <host>").
		Option("--timeout=<seconds>",
			"",
			func() error {
				seconds := Program.GetString("<seconds>")

				t.Log("seconds =", seconds)

				assert.Equal(t, seconds, 0)
				return nil
			},
		).Action(func() error {
		host = Program.GetString("<host>")
		return nil
	})

	if _, err := Program.Parse([]string{"test", "add", "10", "20"}); err != nil {
		t.Fatal(err)
	}
	if _, err := Program.Parse([]string{"test", "ping", "127.0.0.1"}); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, sum, 30)
	assert.Equal(t, host, "127.0.0.1")
}

func TestProgram_Calculator(t *testing.T) {
	var result int

	Program.Command("calculator_example").
		Version("0.0.1").
		Description("Simple calculator example")

	Program.Command("<value> ( ( + | - | * | / ) <value> )...", "", func() error {
		values := Program.GetStrings("<value>")
		for index, value := range values {
			if i, err := strconv.Atoi(value); err != nil {
			} else if index == 0 {
				result = i
			} else {
				switch Program.GetArg(index*2 - 1) {
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

	Program.Command("<function> <value> [( , <value> )]...", "", func() error {
		result = 0
		switch Program.GetString("<function>") {
		case "sum":
			values := Program.GetStrings("<value>")
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
