package fileutils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

// FileExists returns whether a file exists
func (inst *Dirs) FileExists(filePath string) bool {
	f, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return !f.IsDir()
}

// DirExists returns whether a directory exists
func (inst *Dirs) DirExists(dirPath string) bool {
	f, err := os.Stat(dirPath)
	if err != nil {
		return false
	}
	return f.IsDir()
}

// DirExistsErr returns whether a directory exists, returns an error if not exist
func (inst *Dirs) DirExistsErr(dirPath string) error {
	_, err := os.Stat(dirPath)
	if err != nil {
		err = fmt.Errorf("no dir exist name:%s \n", dirPath)
		log.Error(err)
		return err
	}
	return nil
}
