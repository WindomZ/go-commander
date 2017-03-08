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
	cmd := NewCommander("test [data]").
		Version("0.0.1").
		Description("This is a test cli.")
		//Option("--info=<info>",
		//	"info operation",
		//	func(args DocoptMap) Result {
		//		info := args.GetString("<info>")
		//
		//		t.Log("info =", info)
		//
		//		assert.Equal(t, info, 30)
		//		return nil
		//	},
		//	"defalut",
		//)
	cmd.Command("add <x> <y>").
		Description("addition operation").
		Action(func(args DocoptMap) Result {
			x, _ := args.GetInt("<x>")
			y, _ := args.GetInt("<y>")

			t.Log("x =", x)
			t.Log("y =", y)
			t.Log("add x + y =", x+y)

			assert.Equal(t, x+y, 30)
			return nil
		})
	cmd.Command("ping <host>").
		Action(func(args DocoptMap) Result {
			host := args.GetString("<host>")

			t.Log("host =", host)

			assert.Equal(t, host, "127.0.0.1")
			return nil
		}).
		Option("--timeout=<seconds>",
			"",
			func(args DocoptMap) Result {
				seconds := args.GetString("<seconds>")

				t.Log("seconds =", seconds)

				assert.Equal(t, seconds, 0)
				return nil
			},
		)
	cmd.ShowHelpMessage() // only print help message

	//if d, err := cmd.Parse([]string{"test", "--info", "hello", "world"}); err != nil {
	//	t.Logf("1: %s", d.String())
	//	t.Fatal(err)
	//}
	if _, err := cmd.Parse([]string{"test", "add", "10", "20"}); err != nil {
		t.Fatal(err)
	}
	if _, err := cmd.Parse([]string{"test", "ping", "127.0.0.1"}); err != nil {
		t.Fatal(err)
	}
}

func TestCommander_2(t *testing.T) {
}
