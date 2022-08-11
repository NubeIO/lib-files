package fileutils

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

// filePath make the file path work for unix or windows
func filePath(path string, debug ...bool) string {
	updated := filepath.FromSlash(path)
	if len(debug) > 0 {
		if debug[0] {
			log.Infof("existing-path: %s", path)
			log.Infof("updated-path: %s", updated)
		}
	}
	return filepath.FromSlash(updated)
}

// CheckDelete stop the user from delete the root ot home dir
func CheckDelete(path string) error {
	path = filePath(path)
	home, err := HomeDir()
	home = filePath(home)
	root := filePath("/")
	if err != nil {
		return err
	}
	if path == home {
		return errors.New(fmt.Sprintf("user tried to delete home dir %s", home))
	}
	if path == root {
		return errors.New(fmt.Sprintf("user tried to delete root dir /"))
	}
	return nil
}

// RemoveAll implements os.RemoveAll in this directory context.
func (inst *Dirs) RemoveAll(name string, path string) error {
	err := CheckDelete(name)
	if err != nil {
		return err
	}
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
	err := CheckDelete(oldName)
	if err != nil {
		return err
	}
	return os.Rename(oldName, newName)
}

// Rm is `rm` / os.Remove
func (inst *Dirs) Rm(name string) error {
	err := CheckDelete(name)
	if err != nil {
		return err
	}
	return os.Remove(name)
}

// RmRF is `rm -rf` / os.RemoveAll
func (inst *Dirs) RmRF(path string) error {
	err := CheckDelete(path)
	if err != nil {
		return err
	}
	return os.RemoveAll(path)
}
