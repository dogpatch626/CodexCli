package main

import (
	"CodexCli/banner"
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	banner.Banner()
	var input []string
	for scanner.Scan() {

		// fmt.Println(" typed ", scanner.Text() == string('.'))

		// in := bufio.NewReader(os.Stdin)
		in := scanner.Text()
		if in == string(":q") {
			return
		}
		// for idk := 0; idk < 5; idk++ {
		// 	fmt.Println("Weeeeeee")
		// }
		input = append(input, in)

		switch in {
		case string(":info"):
			color.Red("CodexCli is a in development cli made in go as a project to learn what is even going on with this lang.")
		}
		// s = strings.TrimSpace(s)

		fmt.Println(input)

	}
}
