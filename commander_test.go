package commander

import (
	"fmt"
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestNewCommander(t *testing.T) {
	c := NewCommander("hello")
	assert.NotEmpty(t, c)
}

func TestCommander_1(t *testing.T) {
	cmd := NewCommander("test [text]").
		Version("0.0.1").
		Description("This is a test cli.").
		Option(
			"[(-a|--add) <x> <y>]",
			"addition operation",
			func(args DocoptMap) Result {
				x, _ := args.GetInt("<x>")
				y, _ := args.GetInt("<y>")

				//fmt.Println("x =", x)
				//fmt.Println("y =", y)
				//fmt.Println("x + y =", x+y)

				assert.Equal(t, x+y, 30)
				return nil
			},
			1, 2,
		)
	cmd.Command("ping <host>").
		Action(func(args DocoptMap) Result {
			host := args.GetString("<host>")

			fmt.Println("host =", host)

			assert.Equal(t, host, "127.0.0.1")
			return nil
		}).
		Option(
			"--timeout=<seconds>",
			"",
			func(args DocoptMap) Result {
				seconds := args.GetString("<seconds>")

				fmt.Println("seconds =", seconds)

				assert.Equal(t, seconds, 0)
				return nil
			},
		)
	cmd.ShowHelpMessage() // only print help message

	if _, err := cmd.Parse([]string{"test", "-a", "10", "20"}); err != nil {
		t.Fatal(err)
	}

	if _, err := cmd.Parse([]string{"test", "ping", "127.0.0.1"}); err != nil {
		t.Fatal(err)
	}
}

func TestCommander_2(t *testing.T) {
}
