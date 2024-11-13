package main

import (
	"CodexCli/banner"
	TextProcessor "CodexCli/textprocessor"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	banner.Banner()

	err := godotenv.Load()

	if err != nil {
		f, CreationError := os.Create(".env")
		if CreationError != nil {
			log.Fatal("Error loading and writing .env file")
		}
		defer f.Close()
	}

	var myEnv map[string]string
	myEnv, ReadError := godotenv.Read()

	if ReadError != nil {
		log.Fatal("Error reading .env")
	}

	if len(myEnv) == 0 {
		fmt.Println("your .env is empty")
	}

	for scanner.Scan() {

		in := scanner.Text()
		input := strings.Split(in, " ")
		// pass args to the processor
		// arg format :<Command> ...Perams
		TextProcessor.Processor(input)

	}
}
