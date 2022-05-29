package fileutils

import (
	"os"
	"path/filepath"
)

// RemoveAll implements os.RemoveAll in this directory context.
func (inst *Dirs) RemoveAll(name string, path string) error {
	if name = inst.resolve(name); name == "" {
		return os.ErrNotExist
	}

	if name == filepath.Clean(path) {
		// Prohibit removing the virtual root directory.
		return os.ErrInvalid
	}
	return os.RemoveAll(name)
}

// Mv is `mv` / os.Rename
func (inst *Dirs) Mv(oldName, newName string) error {
	return os.Rename(oldName, newName)
}

// Rm is `rm` / os.Remove
func (inst *Dirs) Rm(name string) error {
	return os.Remove(name)
}

// RmRF is `rm -rf` / os.RemoveAll
func (inst *Dirs) RmRF(path string) error {
	return os.RemoveAll(path)
}
