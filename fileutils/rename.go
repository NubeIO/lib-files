package fileutils

import (
	"os"
)

// Rename implements os.Rename in this directory context.
func (inst *FileUtils) Rename(oldName, newName string) error {
	if oldName = inst.resolve(oldName); oldName == "" {
		return os.ErrNotExist
	}
	if newName = inst.resolve(newName); newName == "" {
		return os.ErrNotExist
	}
	return os.Rename(oldName, newName)
}
