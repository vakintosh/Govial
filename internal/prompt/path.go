package prompt

import (
	"path/filepath"
	"strconv"
	"strings"
)

// Truncates the current working directory
func trimPath(cwd, home string) string {
	var path string
	if cwd == home {
		return "~"
	}
	if strings.HasPrefix(cwd, home) {
		path = "~" + strings.TrimPrefix(cwd, home)
	} else {
		// If path doesn't contain $HOME, return the
		// entire path as is.
		path = cwd
		return path
	}
	// Splits the path string into individual directory names using "/" as separator
	items := strings.Split(path, "/")
	// Init an empty slice of strings to store the truncated directory names
	truncItems := []string{}
	for i, item := range items {
		if i == (len(items) - 1) {
			truncItems = append(truncItems, item)
			break
		}
		truncItems = append(truncItems, "")
	}
	// Print number of directories between ~ and CWD. If numberDirs is 0, the output only contains $HOME/CWD
	numberDirs := len(items) - 2 // Subtract 2 to exclude home and the last directory
	if numberDirs == 0 {
		return "~/" + filepath.Join(truncItems...)
	}
	return "~/" + strconv.Itoa(numberDirs) + "/" + filepath.Join(truncItems...)
}
