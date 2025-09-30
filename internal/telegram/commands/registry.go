package commands

var AllCommands []Command

func Register(cmd Command) {
	AllCommands = append(AllCommands, cmd)
}

func Init() {
	Register(&StartCommand{})
	Register(&SomethingCommand{})
}
