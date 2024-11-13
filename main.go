package main

import (
	"CodexCli/banner"
	TextProcessor "CodexCli/textprocessor"
	"bufio"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	banner.Banner()

	for scanner.Scan() {

		in := scanner.Text()
		input := strings.Split(in, " ")
		// pass args to the processor
		// arg format :<Command> ...Perams
		TextProcessor.Processor(input)

	}
}
