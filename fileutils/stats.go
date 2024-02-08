package fileutils

import (
	"github.com/NubeIO/lib-files/fileutils/size"
	"os"
	"path/filepath"
)

// Stat implements os.Stat in this directory context.
func Stat(name string) (os.FileInfo, error) {
	if name = resolve(name); name == "" {
		return nil, os.ErrNotExist
	}
	return os.Stat(name)
}

// GetFileSize returns file size )
func GetFileSize(name string) (out size.ByteSize, err error) {
	fi, err := os.Stat(name)
	if err != nil {
		return 0, err
	}
	f := float64(fi.Size())
	return size.NewSize(f), nil
}

// GetFolderSize returns folder size
func GetFolderSize(path string) (out size.ByteSize, err error) {
	var f int64
	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			f += info.Size()
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return size.NewSize(float64(f)), nil
}
