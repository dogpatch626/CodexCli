package commands

import (
	"os"

	"github.com/fatih/color"
)

func Info() {
	color.Red("CodexCli is a in development cli made in go as a project to learn what is even going on with this lang.")
}

func DoSomething() {
	color.Red("do something")
}

func AboutMe() {
	color.Green("Greetings, Im Danny! I'm a software Engineer from the Greater NYC area. I have 3 years of professional experience now previously working at American Express as an Engineer 3")
}

func Scan() {

}

func Quit() {
	os.Exit(0)
}
