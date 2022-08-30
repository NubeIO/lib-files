package fileutils

import (
	"os"
	"path/filepath"
	"strings"
)

// A Dir uses the native file system restricted to a specific directory tree.
// An empty Dir is treated as ".".

type FileUtils struct{}

func New() *FileUtils {
	return &FileUtils{}
}

func resolve(name string) string {
	// This implementation is based on Dir.Open's code in the standard net/http package.
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) ||
		strings.Contains(name, "\x00") {
		return ""
	}
	return filepath.FromSlash(SlashClean(name))
}

// OpenFile implements os.OpenFile in this directory context.
func OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	if name = resolve(name); name == "" {
		return nil, os.ErrNotExist
	}
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}
