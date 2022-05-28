package fileutils

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestDir(t *testing.T) {
	//tempdir, err := ioutil.TempDir("", "aidan")
	//if err != nil {
	//	t.Error(err)
	//}
	//fmt.Println(tempdir)
	//defer os.RemoveAll(tempdir)

	var err error

	folder := filepath.Join("bin", "folder")
	//err = CopyDir("./testdata", folder)
	//if err != nil {
	//	t.Error(err)
	//}
	perm := 0700
	testdata := Dir(folder)
	fmt.Println(111111, testdata)
	err = testdata.Mkdir("test/1/2/3", os.FileMode(perm))
	if err != nil {
		t.Error(err)
	}

	zip := fmt.Sprintf("%s/%s", "testdata", "rubix-bios-1.5.2-d5764dc0.amd64.zip")
	fmt.Println(zip)
	dest := fmt.Sprintf("%s/%s", testdata, "test/1/2/3")
	extract, err := testdata.Extract(zip, dest, os.FileMode(perm))

	fmt.Println(extract, err)
	fmt.Println(testdata)
	//err = testdata.RemoveAll("test")
	//fmt.Println(err)
	//if err != nil {
	//	t.Error(err)
	//}

}
