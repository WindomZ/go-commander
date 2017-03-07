package commander

type ExecFunc func(args map[string]interface{}, exit ...bool)
