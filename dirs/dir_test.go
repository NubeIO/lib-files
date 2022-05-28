package fileutils

import (
	"fmt"
	"os"
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

	path := "bin"
	//err = CopyDir("./testdata", folder)
	//if err != nil {
	//	t.Error(err)
	//}
	perm := 0700

	dirs := New(&Dirs{path})

	fmt.Println(111111, dirs)
	err = dirs.Mkdir("test/1/2/3", os.FileMode(perm))
	if err != nil {
		t.Error(err)
	}

	zip := fmt.Sprintf("%s/%s", "testdata", "rubix-bios-1.5.2-d5764dc0.amd64.zip")
	fmt.Println(zip)
	dest := fmt.Sprintf("%s/%s", dirs, "test/1/2/3")
	extract, err := dirs.UnZip(zip, dest, os.FileMode(perm))

	fmt.Println(extract, err)
	fmt.Println(dirs)
	//err = testdata.RemoveAll("test")
	//fmt.Println(err)
	//if err != nil {
	//	t.Error(err)
	//}

}
