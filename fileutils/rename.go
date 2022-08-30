package fileutils

import (
	"os"
)

// Rename implements os.Rename in this directory context.
func Rename(oldName, newName string) error {
	if oldName = resolve(oldName); oldName == "" {
		return os.ErrNotExist
	}
	if newName = resolve(newName); newName == "" {
		return os.ErrNotExist
	}
	return os.Rename(oldName, newName)
}
