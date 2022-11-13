package fileutils

import (
	"fmt"
	"os"
)

// FileExists returns whether a file exists
func FileExists(filePath string) bool {
	f, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return !f.IsDir()
}

// DirExists returns whether a directory exists
func DirExists(dirPath string) bool {
	f, err := os.Stat(dirPath)
	if err != nil {
		return false
	}
	return f.IsDir()
}

// DirExistsErr returns whether a directory exists, returns an error if not exist
func DirExistsErr(dirPath string) error {
	_, err := os.Stat(dirPath)
	if err != nil {
		err = fmt.Errorf("no dir exist name: %s", dirPath)
		return err
	}
	return nil
}

func FileOrDirExists(dirPath string) bool {
	_, err := os.Stat(dirPath)
	return err == nil
}
