package fileutils

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

// GetFileName returns file name
func (inst *FileUtils) GetFileName(filePath string) string {
	return path.Base(filePath)
}

// GetExt returns extension name
// Will return `` when giving a string `.txt` or `.foo.zip.` etc
func (inst *FileUtils) GetExt(filePath string) string {
	if len(filePath) == 0 {
		return empty
	}
	if -1 == strings.Index(filePath, `.`) {
		return empty
	}
	if ok, _ := regexp.MatchString(`^\.[^\.]*$`, filePath); ok {
		return empty
	}
	if string(filePath[len(filePath)-1]) == `.` {
		return empty
	}
	return path.Ext(filePath)
}

func (inst *FileUtils) ReadJSON(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("open: %s", err.Error())
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("read all: %s", err.Error())
	}
	return bytes, nil
}

func (inst *FileUtils) ListFiles(file string) ([]string, error) {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return nil, err
	}
	var dirContent []string
	if fileInfo.IsDir() {
		files, err := ioutil.ReadDir(file)
		if err != nil {
			return nil, err
		}
		for _, file := range files {
			dirContent = append(dirContent, file.Name())
		}
	}
	return dirContent, nil
}

// ReadFile returns file content with err
func (inst *FileUtils) ReadFile(filePath string) (content string, err error) {
	f, err := os.Stat(filePath)
	if err != nil {
		return
	}
	if f.IsDir() {
		return empty, errors.New("not a file")
	}
	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Errorf("unable to read file: %v", err)
	}
	return string(body), nil
}

// ReadAll returns file content,will return `` if err
func (inst *FileUtils) ReadAll(filePath string) string {
	f, err := os.Stat(filePath)
	if err != nil {
		return empty
	}
	if f.IsDir() {
		return empty
	}
	fo, err := os.Open(filePath)
	if err != nil {
		return empty
	}
	defer fo.Close()
	fd, err := ioutil.ReadAll(fo)
	if err != nil {
		return empty
	}
	return string(fd)
}

// ReadAllOk returns file content with err
func (inst *FileUtils) ReadAllOk(filePath string) (content string, err error) {
	f, err := os.Stat(filePath)
	if err != nil {
		return
	}
	if f.IsDir() {
		return empty, errors.New("not a file")
	}
	fo, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer fo.Close()
	fd, err := ioutil.ReadAll(fo)
	if err != nil {
		return
	}
	return string(fd), nil
}
