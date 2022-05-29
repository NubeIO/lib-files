package fileutils

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// WriteFile implements ioutil.WriteFile
func (inst *Dirs) WriteFile(filePath string, file string, perm os.FileMode) error {
	err := ioutil.WriteFile(filePath, []byte(file), perm)
	if err != nil {
		log.Println("write file error", err)
		return err
	}
	return err
}

func (inst *Dirs) CreateFile(filePath string, perm os.FileMode) (*os.File, error) {
	if path, err := inst.MakeFilePath(filepath.Dir(filePath), filepath.Base(filePath), perm); err != nil {
		return nil, err
	} else {
		return os.Create(path)
	}
}

// MkdirAll implements os.Mkdir in this directory context.
func (inst *Dirs) MkdirAll(name string, perm os.FileMode) error {
	if name = inst.resolve(name); name == "" {
		return os.ErrNotExist
	}
	return os.MkdirAll(name, perm)
}

// MakeFilePath make file and path
func (inst *Dirs) MakeFilePath(dirName, fileName string, perm os.FileMode) (string, error) {
	if err := inst.EnsureDir(dirName, perm); err != nil {
		return "", err
	}
	return filepath.Join(dirName, fileName), nil
}

func (inst *Dirs) EnsureDir(dirName string, perm os.FileMode) error {
	if err := os.MkdirAll(dirName, perm); err == nil || os.IsExist(err) {
		return nil
	} else {
		return err
	}
}
