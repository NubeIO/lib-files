package fileutils

import (
	"fmt"
	"testing"
)

func Test_checkDelete(t *testing.T) {
	err := CheckDelete("/")
	fmt.Println(err)
	dir, err := HomeDir()
	if err != nil {
		return
	}
	err = CheckDelete(dir)
	fmt.Println(err)
	err = CheckDelete("/lib")

	err = New().Rm("/")
	fmt.Println("lets test it for real ERROR", err)
}
