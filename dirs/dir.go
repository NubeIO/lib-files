package fileutils

import (
	"os"
	"path/filepath"
	"strings"
)

// A Dir uses the native file system restricted to a specific directory tree.
// An empty Dir is treated as ".".

type Dirs struct {
	Path string
}

func New(dirs *Dirs) *Dirs {
	return dirs

}

func (inst *Dirs) resolve(name string) string {
	// This implementation is based on Dir.Open's code in the standard net/http package.
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) ||
		strings.Contains(name, "\x00") {
		return ""
	}
	dir := inst.Path
	if dir == "" {
		dir = "."
	}
	return filepath.Join(dir, filepath.FromSlash(SlashClean(name)))
}

func (inst *Dirs) String() string {
	return inst.Path
}

// Mkdir implements os.Mkdir in this directory context.
func (inst *Dirs) Mkdir(name string, perm os.FileMode) error {
	if name = inst.resolve(name); name == "" {
		return os.ErrNotExist
	}
	return os.MkdirAll(name, perm)
}

// OpenFile implements os.OpenFile in this directory context.
func (inst *Dirs) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	if name = inst.resolve(name); name == "" {
		return nil, os.ErrNotExist
	}
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// RemoveAll implements os.RemoveAll in this directory context.
func (inst *Dirs) RemoveAll(name string) error {
	if name = inst.resolve(name); name == "" {
		return os.ErrNotExist
	}

	if name == filepath.Clean(inst.Path) {
		// Prohibit removing the virtual root directory.
		return os.ErrInvalid
	}
	return os.RemoveAll(name)
}

// Rename implements os.Rename in this directory context.
func (inst *Dirs) Rename(oldName, newName string) error {
	if oldName = inst.resolve(oldName); oldName == "" {
		return os.ErrNotExist
	}
	if newName = inst.resolve(newName); newName == "" {
		return os.ErrNotExist
	}
	if root := filepath.Clean(inst.Path); root == oldName || root == newName {
		// Prohibit renaming from or to the virtual root directory.
		return os.ErrInvalid
	}
	return os.Rename(oldName, newName)
}

// Stat implements os.Stat in this directory context.
func (inst *Dirs) Stat(name string) (os.FileInfo, error) {
	if name = inst.resolve(name); name == "" {
		return nil, os.ErrNotExist
	}
	return os.Stat(name)
}

// Copy copies a file or directory from src to dst. If it is
// a directory, all of the files and sub-directories will be copied.
func (inst *Dirs) Copy(src, dst string) error {
	if src = inst.resolve(src); src == "" {
		return os.ErrNotExist
	}
	if dst = inst.resolve(dst); dst == "" {
		return os.ErrNotExist
	}
	if root := filepath.Clean(inst.Path); root == src || root == dst {
		// Prohibit copying from or to the virtual root directory.
		return os.ErrInvalid
	}
	if dst == src {
		return os.ErrInvalid
	}
	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return CopyDir(src, dst)
	}
	return CopyFile(src, dst)
}
