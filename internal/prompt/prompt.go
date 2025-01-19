// prompt.go
package prompt

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
	return len(width)
}

// MakePrompt generates the terminal prompt
func MakePrompt() string {
	cwd, _ := os.Getwd()
	home := os.Getenv("HOME")

	hour, minute, second := getCurrentTime()

	leftPart := JOVIAL_SYMBOL["corner.top"] + yellow(trimPath(cwd, home))
	bottomPart := JOVIAL_SYMBOL["corner.bottom"] + JOVIAL_SYMBOL["arrow"]

	terminalWidth := getTerminalWidth()
	leftWidth := len(leftPart)
	bottomWidth := len(bottomPart)

	timeString := cyan(fmt.Sprintf("%02d:%02d:%02d", hour, minute, second))
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
