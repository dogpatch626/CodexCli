package TextProcessor

import (
	"CodexCli/commands"
)

func Processor(command []string) {

	switch command[0] {
	case string(":info"):
		commands.Info()
	case string(":do"):
		commands.DoSomething()
	case string(":ab"):
		commands.AboutMe()
	case string(":q"):
		commands.Quit()
	case string(":setKey"):
		commands.SetKey(command[1])
	}

}
