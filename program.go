package commander

import (
	"fmt"
	"os"
)

type GoCommander interface {
	Commander
	Context
}

var Program GoCommander = newProgram()

type _Program struct {
	_Command
	_Context
	Format  Formatter
	Exec    Execute
	initial bool
	version string // version if root command
	errFunc func(error)
}

func newProgram() *_Program {
	return &_Program{
		_Command: *newCommand(true),
		Format:   Format,
		Exec:     Exec,
		errFunc: func(err error) {
			fmt.Println(err.Error())
		},
	}
}

func (p *_Program) init() *_Program {
	if p.initial {
		return p
	}
	p.initial = true
	p._Command.init()
	p.Command("-h --help", "output usage information", func() _Result {
		p.ShowHelpMessage()
		return resultBreak()
	})
	p.Command("-v --version", "output the version number", func() _Result {
		p.ShowVersion()
		return resultBreak()
	})
	return p
}

// Version defines the version of this command.
// Only valid under the root command.
func (p *_Program) Version(ver string) Commander {
	p._Command.Version(ver)
	p.version = ver
	return p
}

// ShowVersion display the version.
func (p _Program) ShowVersion() string {
	fmt.Println(p.version)
	return p.version
}

// HelpMessage returns, as a string, all help messages that generated according to the docopt format.
func (p *_Program) HelpMessage() string {
	return p.init()._Command.HelpMessage()
}

// ShowHelpMessage display all help messages.
func (p _Program) ShowHelpMessage() string {
	s := p.HelpMessage()
	fmt.Println(s)
	return s
}

// Parse parse `argv` based on the command-line interface.
func (p *_Program) Parse(args ...[]string) (Context, error) {
	var argv []string = nil
	if len(args) != 0 {
		argv = args[0]
	}
	if argv == nil && len(os.Args) > 1 {
		argv = os.Args
	}
	if argv == nil || len(argv) == 0 {
		argv = []string{"", "-h"}
	}
	d, err := Parse(p.HelpMessage(), argv, false, "", false, false)
	if err != nil {
		return nil, err
	}
	p._Context = *newContext(argv, d)
	if r := p.run(&p._Context); r != nil {
		if r.Break() {
			if p.errFunc != nil && r.Error() != nil {
				p.errFunc(r.Error())
			}
			return &p._Context, r.Error()
		}
	}
	return &p._Context, nil
}

// ErrorHandling defines the function f to handle the throwing error.
func (p *_Program) ErrorHandling(f func(error)) Commander {
	p.errFunc = f
	return p
}
