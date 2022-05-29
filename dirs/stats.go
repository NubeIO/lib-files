package fileutils

import (
	"github.com/NubeIO/lib-dirs/dirs/size"
	"os"
)

// Stat implements os.Stat in this directory context.
func (inst *Dirs) Stat(name string) (os.FileInfo, error) {
	if name = inst.resolve(name); name == "" {
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
