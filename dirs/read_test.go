package fileutils

import (
	"fmt"
	"testing"
)

func TestDirs_ReadFile(t *testing.T) {

	a := New()
	file, err := a.ReadFile("/data/auth/user.txt")
	if err != nil {
		return
	}
	fmt.Println(file)

}
