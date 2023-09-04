package main

import (
	"fmt"
	"os"
)

var (
	red    = color("\033[31m%s\033[0m")
	green  = color("\033[32m%s\033[0m")
	cyan   = color("\033[36m%s\033[0m")
	yellow = color("\033[33m%s\033[0m")
)

func color(s string) func(...interface{}) string {
	return func(args ...interface{}) string {
		return fmt.Sprintf(s, fmt.Sprint(args...))
	}
}

var JOVIAL_SYMBOL = map[string]string{
	"corner.top":      "╭─",
	"corner.bottom":   "╰─",
	"git.dirty":       "✘✘✘",
	"git.clean":       "✔",
	"arrow":           "─▶",
	"arrow.git-clean": "(๑˃̵ᴗ˂̵)و",
	"arrow.git-dirty": "(ﾉ˚Д˚)ﾉ",
}

func makePrompt() string {
	cwd, _ := os.Getwd()
	home := os.Getenv("HOME")
	return fmt.Sprintf(
		"\n%s",
		JOVIAL_SYMBOL["corner.top"]+yellow(trimPath(cwd, home))+"\n"+JOVIAL_SYMBOL["corner.bottom"]+JOVIAL_SYMBOL["arrow"],
	)
}

func main() {
	fmt.Println(makePrompt())
}
