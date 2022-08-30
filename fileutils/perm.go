package fileutils

import (
	"os"
	"path/filepath"
)

// ChmodR is like `chmod -R`
func (inst *FileUtils) ChmodR(name string, mode os.FileMode) error {
	return filepath.Walk(name, func(path string, info os.FileInfo, err error) error {
		if err == nil {
			err = os.Chmod(path, mode)
		}
		return err
	})
}

// ChownR is like `chown -R`
func (inst *FileUtils) ChownR(path string, uid, gid int) error {
	return filepath.Walk(path, func(name string, info os.FileInfo, err error) error {
		if err == nil {
			err = os.Chown(name, uid, gid)
		}
		return err
	})
}
