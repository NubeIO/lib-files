package fileutils

import (
	"fmt"
	"os"
	"testing"
)

func TestDir(t *testing.T) {
	var err error
	perm := 0600
	err = MkdirAll("test/1/2/3", os.FileMode(perm))
	if err != nil {
		t.Error(err)
	}
	zipFolder := "bin/testzip.zip"
	// zip a folder
	perm = 0777
	files := []string{"testdata/file.txt"}
	err = ZipFiles(zipFolder, files)
	if err != nil {
		t.Error(err)
	}

	// unzip a folder
	dest := fmt.Sprintf("bin/test/1/2/3")
	extract, err := UnZip(zipFolder, dest, os.FileMode(perm))
	fmt.Println(extract, err)
}
