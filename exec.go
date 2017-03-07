package commander

type ExecFunc func(args DocoptMap, exit ...bool)
