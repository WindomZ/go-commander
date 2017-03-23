package commander

import "os"

type GoCommander interface {
	Commander
	Context
}

var Program GoCommander = newProgram()

type _Program struct {
	_Command
	_Context
	version string // version if root command
}

func newProgram() *_Program {
	return &_Program{
		_Command: *newCommand(true),
	}
}

func (p *_Program) Version(ver string) Commander {
	p.version = ver
	return p
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
	d, err := Parse(p.HelpMessage(), argv, true, p.version, false, false)
	if err != nil {
		return nil, err
	}
	p._Context = *newContext(argv, d)
	if r := p.run(&p._Context); r != nil {
		if r.Break() {
			// TODO: Handle error
			return &p._Context, r.Error()
		}
	} else {
		// TODO: Should be print help message, but docopt auto to do this.
	}
	return &p._Context, nil
}
