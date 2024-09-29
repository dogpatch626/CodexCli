package main

import (
	"CodexCli/banner"
	"CodexCli/commands"
	"bufio"
	"fmt"
	"os"
)

type value struct {
	Value func()
}

func DoSomething() {
	fmt.Println("does Something")

}
func main() {

	scanner := bufio.NewScanner(os.Stdin)
	banner.Banner()

	m := make(map[string]func())
	m[":do"] = commands.DoSomething
	m[":info"] = commands.Info
	var input []string
	for scanner.Scan() {

		in := scanner.Text()
		if in == string(":q") {
			return
		}

		input = append(input, in)

		switch in {
		case string(":info"):
			m[":info"]()
		case string(":do"):
			m[":do"]()
		}

		fmt.Println(input)

	}
}
