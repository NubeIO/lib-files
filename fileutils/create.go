package fileutils

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// WriteFileByte implements ioutil.WriteFile
func WriteFileByte(filePath string, body []byte, perm os.FileMode) error {
	err := ioutil.WriteFile(filePath, body, perm)
	if err != nil {
		return err
	}
	return err
}

// WriteFile implements ioutil.WriteFile as body as a string
func WriteFile(filePath string, body string, perm os.FileMode) error {
	err := ioutil.WriteFile(filePath, []byte(body), perm)
	if err != nil {
		return err
	}
	return err
}

func CreateFile(filePath string, perm os.FileMode) (*os.File, error) {
	if path, err := MakeFilePath(filepath.Dir(filePath), filepath.Base(filePath), perm); err != nil {
		return nil, err
	} else {
		return os.Create(path)
	}
}

// MkdirAll implements os.Mkdir in this directory context.
func MkdirAll(name string, perm os.FileMode) error {
	if name = resolve(name); name == "" {
		return os.ErrNotExist
	}
	return os.MkdirAll(name, perm)
}

// MakeFilePath make file and path
func MakeFilePath(dirName, fileName string, perm os.FileMode) (string, error) {
	if err := EnsureDir(dirName, perm); err != nil {
		return "", err
	}
	return filepath.Join(dirName, fileName), nil
}

func EnsureDir(dirName string, perm os.FileMode) error {
	if err := os.MkdirAll(dirName, perm); err == nil || os.IsExist(err) {
		return nil
	} else {
		return err
	}
}
