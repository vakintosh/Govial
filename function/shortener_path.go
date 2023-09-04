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
			fmt.Println(items)
			fmt.Println(item)
			fmt.Println(truncItems)
			break
		}
		truncItems = append(truncItems, item[:1])
	}
	return filepath.Join(truncItems...)
}
