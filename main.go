// This file is part of Govial Prompt.
// It is the main file of the project
// The main function is used to call the other functions and display responsive-design  prompt
//

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	red    = color("\033[31m%s\033[0m")
	green  = color("\033[32m%s\033[0m")
	cyan   = color("\033[36m%s\033[0m")
	yellow = color("\033[33m%s\033[0m")
)

func color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

var JOVIAL_SYMBOL = map[string]string{
	"corner.top":      "╭─ ",
	"corner.bottom":   "╰─",
	"git.dirty":       "✘✘✘",
	"git.clean":       "✔",
	"arrow":           "─▶",
	"arrow.git-clean": "(๑˃̵ᴗ˂̵)و",
	"arrow.git-dirty": "(ﾉ˚Д˚)ﾉ",
}

func getTerminalWidth() int {
	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	width := strings.TrimSpace(string(out))
	fmt.Println("Terminal Width:", width)
	return len(width)
}

// Function to make the prompt
// The prompt is made up of the following components:
// 1. The current working directory on the left side of the screen
// 3. The current time from the function getCurrentTime on the right side of the screen
// 6. The arrow
// 8. The Terraform workspace

func makePrompt() string {
	cwd, _ := os.Getwd()
	home := os.Getenv("HOME")

	hour, minute := getCurrentTime()

	leftPart := JOVIAL_SYMBOL["corner.top"] + yellow(trimPath(cwd, home))
	bottomPart := JOVIAL_SYMBOL["corner.bottom"] + JOVIAL_SYMBOL["arrow"]

	terminalWidth := getTerminalWidth()
	leftWidth := len(leftPart)
	bottomWidth := len(bottomPart)

	timeString := cyan(fmt.Sprintf("%02d:%02d", hour, minute))
	timeWidth := len(timeString)

	// Calculate available space for the right side
	availableSpace := terminalWidth - leftWidth - timeWidth - bottomWidth - 1 // Subtract 1 for the space before the time

	// Ensure available space is not negative
	if availableSpace < 0 {
		availableSpace = 0
	}

	// Pad spaces on the right to align time to the end
	timeAligned := strings.Repeat(" ", availableSpace) + timeString

	return fmt.Sprintf("%s%s\n%s", leftPart, timeAligned, bottomPart)
}

func main() {
	fmt.Println(makePrompt())
}
