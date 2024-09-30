package main

import (
	"CodexCli/banner"
	"CodexCli/commands"
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	banner.Banner()

	m := make(map[string]func())
	m[":do"] = commands.DoSomething
	m[":info"] = commands.Info
	m[":ab"] = commands.AboutMe
	m[":q"] = commands.Quit
	var input []string
	for scanner.Scan() {

		in := scanner.Text()

		input = append(input, in)

		switch in {
		case string(":info"):
			m[":info"]()
		case string(":do"):
			m[":do"]()
		case string(":ab"):
			m[":ab"]()
		case string(":q"):
			m[":q"]()
		}

		fmt.Println(input)

	}
}
