package fileutils

import (
	"fmt"
	"testing"
)

func TestDirs_RecursiveZip(t *testing.T) {
	err := RecursiveZip("/data", "./test.zip")
	fmt.Println(err)
	if err != nil {
		return
	}
}
