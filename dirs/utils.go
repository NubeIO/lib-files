package fileutils

import (
	"os/exec"
	"path"
)

var (
	empty = ``
)

// SlashClean is equivalent to but slightly more efficient than
// path.Clean("/" + name).
func SlashClean(name string) string {
	if name == "" || name[0] != '/' {
		name = "/" + name
	}
	return path.Clean(name)
}

// Which is `which` / exec.LookPath
func Which(file string) (string, error) {
	return exec.LookPath(file)
}
