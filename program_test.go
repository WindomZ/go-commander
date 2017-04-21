package commander

import (
	"github.com/WindomZ/testify/assert"
	"strconv"
	"testing"
)

func TestProgram_AutomaticHelp(t *testing.T) {
	Program = newProgram()

	Program.Command("go-commander").
		Version("0.0.1").
		Description("this is a test cli.")

	if _, err := Program.Parse([]string{}); err != nil {
		t.Fatal(err)
	}
	if _, err := Program.Parse([]string{"go-commander", "-h"}); err != nil {
		t.Fatal(err)
	}
}

func TestProgram_ShowVersion(t *testing.T) {
	Program = newProgram()

	Program.Command("go-commander").
		Version("0.0.1").
		Description("this is a test cli.")

	if _, err := Program.Parse([]string{"go-commander", "-v"}); err != nil {
		t.Fatal(err)
	}
}

func TestProgram_ErrorHandling(t *testing.T) {
	Program = newProgram()

	Program.Command("go-commander").
		Version("0.0.1").
		Description("this is a test cli.")

	Program.ErrorHandling(func(err error) {
		assert.Error(t, err)
	})

	Program.Command("test").Action(func() error {
		return newError("this is a test error")
	})

	if _, err := Program.Parse([]string{"go-commander", "test"}); err == nil {
		assert.Error(t, err)
	}
}

func TestProgram_LineOption(t *testing.T) {
	var result string
	var result2 string

	Program = newProgram()

	Program.Command("test").
		Version("0.0.1").
		Description("this is a test cli.")

	Program.Command("-a --aaa").
		Action(func() {
			result = "aaa"
		}).
		Option("-d --ddd").
		Action(func() {
			result2 = "ddd"
		})
	Program.Command("-b --bbb", "",
		func() {
			result = "bbb"
		})
	Program.Command("-c").
		Action(func() {
			result = "ccc"
		})

	if _, err := Program.Parse([]string{"test", "-a"}); err != nil {
		t.Fatal(err)
	} else {
		assert.Equal(t, result, "aaa")
	}
	if _, err := Program.Parse([]string{"test", "-a", "-d"}); err != nil {
		t.Fatal(err)
	} else {
		assert.Equal(t, result, "aaa")
		assert.Equal(t, result2, "ddd")
	}
	if _, err := Program.Parse([]string{"test", "--bbb"}); err != nil {
		t.Fatal(err)
	} else {
		assert.Equal(t, result, "bbb")
	}
	if _, err := Program.Parse([]string{"test", "-c"}); err != nil {
		t.Fatal(err)
	} else {
		assert.Equal(t, result, "ccc")
	}
}

func TestProgram_Aliases(t *testing.T) {
	Program = newProgram()

	Program.Command("go-commander")

	Program.Command("-i --init")

	Program.Command("-o").
		Aliases([]string{"--origin"})

	assert.Equal(t, Program.HelpMessage(), `  Usage:
    go-commander -i|--init
    go-commander -o|--origin
    go-commander -h|--help
    go-commander -v|--version

  Options:
    -h --help     output usage information
    -v --version  output the version number
`)
}

func TestProgram_CommandDescription(t *testing.T) {
	Program = newProgram()

	Program.Command("go-commander")

	Program.Command("-i --init").
		Description("this is init flag")

	Program.Command("-o").
		Aliases([]string{"--origin"}).
		Description("this is origin flag")

	assert.Equal(t, Program.HelpMessage(), `  Usage:
    go-commander -i|--init
    go-commander -o|--origin
    go-commander -h|--help
    go-commander -v|--version

  Options:
    -h --help     output usage information
    -i --init     this is init flag
    -o --origin   this is origin flag
    -v --version  output the version number
`)
}

func TestProgram_Ping(t *testing.T) {
	var sum int
	var host string
	Program = newProgram()

	Program.Command("test").
		Version("0.0.1").
		Description("this is a test cli.")

	Program.Command("add <x> <y>").
		Description("addition operation").
		Action(func() error {
			sum = Program.MustInt("<x>") + Program.MustInt("<y>")
			return nil
		})

	Program.Command("ping <host>").
		Option("--timeout=<seconds>",
			"",
			func() error {
				seconds := Program.MustString("<seconds>")

				t.Log("seconds =", seconds)

				assert.Equal(t, seconds, 0)
				return nil
			},
		).Action(func() error {
		host = Program.MustString("<host>")
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
	Program = newProgram()

	Program.Command("calculator_example").
		Version("0.0.1").
		Description("simple calculator example")

	Program.Command("<value> ( ( + | - | * | / ) <value> )...", "", func() error {
		values := Program.MustStrings("<value>")
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
		switch Program.MustString("<function>") {
		case "sum":
			values := Program.MustStrings("<value>")
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
