// This file is part of the GOVIAL PROMPT project
// It contains the trimPath function which is used to truncate the current working directory
// The function is used in the main.go file to truncate the current working directory and display it on the left side of the screen

package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// Truncates the current working directory:
//
//	/home/icy/foo/bar -> ~/f/bar
func trimPath(cwd, home string) string {
	var path string
	if strings.HasPrefix(cwd, home) {
		path = "~" + strings.TrimPrefix(cwd, home)
	} else {
		// If path doesn't contain $HOME, return the
		// entire path as is.
		path = cwd
		return path
	}
	items := strings.Split(path, "/")
	truncItems := []string{}
	for i, item := range items {
		if i == (len(items) - 1) {
			truncItems = append(truncItems, item)
			fmt.Println(len(items) - 2)
			fmt.Println(items)
			fmt.Println(item)
			fmt.Println(truncItems)
			break
		}
		// ~/Documents/DevOps/Git/govial
		// truncItems = append(truncItems, item)

		// ~/D/D/G/govial
		// truncItems = append(truncItems, item[:1])

		// govial
		truncItems = append(truncItems, "")
	}
	return filepath.Join(truncItems...)
}
