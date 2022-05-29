package fileutils

import (
	"fmt"
	"os"
	"testing"
)

func TestDir(t *testing.T) {

	var err error
	perm := 0777
	dirs := New(&Dirs{})
	err = dirs.MkdirAll("test/1/2/3", os.FileMode(perm))
	if err != nil {
		t.Error(err)
	}
	zipFolder := "bin/testzip.zip"
	//zip a folder
	perm = 0777
	files := []string{"testdata/file.txt"}
	err = dirs.ZipFiles(zipFolder, files)
	if err != nil {
		t.Error(err)
	}

	//unzip a folder
	dest := fmt.Sprintf("bin/test/1/2/3")
	extract, err := dirs.UnZip(zipFolder, dest, os.FileMode(perm))
	fmt.Println(extract, err)
	fmt.Println(dirs)

	size, err := GetFileSize("testdata/file.txt")
	if err != nil {
		//return
	}

	fmt.Println(size, err)
	//err = testdata.RemoveAll("test")
	//fmt.Println(err)
	//if err != nil {
	//	t.Error(err)
	//}

}
