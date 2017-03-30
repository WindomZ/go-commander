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
	initial bool
	version string // version if root command
}

func newProgram() *_Program {
	return &_Program{
		_Command: *newCommand(true),
	}
}

func (p *_Program) init() *_Program {
	if p.initial {
		return p
	}
	p.initial = true
	p._Command.init()
	p.Command("-h --help", "show help message", func() _Result {
		p.ShowHelpMessage()
		return resultBreak()
	})
	p.Command("-v --version", "show version", func() _Result {
		p.ShowVersion()
		return resultBreak()
	})
	return p
}

func (p *_Program) Version(ver string) Commander {
	p.version = ver
	return p._Command.init()
}

func (p _Program) ShowVersion() string {
	fmt.Println(p.version)
	return p.version
}

func (p _Program) HelpMessage() string {
	return p.init()._Command.HelpMessage()
}

func (p _Program) ShowHelpMessage() string {
	s := p.HelpMessage()
	fmt.Println(s)
	return s
}

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
			// TODO: Handle error
			return &p._Context, r.Error()
		}
	}
	return &p._Context, nil
}
