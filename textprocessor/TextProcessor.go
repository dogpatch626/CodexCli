package TextProcessor

import (
	"CodexCli/commands"
)

func Processor(command []string) {

	m := make(map[string]func())

	m[":do"] = commands.DoSomething
	m[":info"] = commands.Info
	m[":ab"] = commands.AboutMe
	m[":q"] = commands.Quit

	switch command[0] {
	case string(":info"):
		m[":info"]()
	case string(":do"):
		m[":do"]()
	case string(":ab"):
		m[":ab"]()
	case string(":q"):
		m[":q"]()
	}

}
