package fileutils

import (
	"fmt"
	"testing"
)

func TestDirs_ReadFile(t *testing.T) {
	a := New()
	file, err := a.ReadAllOk("/data/flow-framework/config/config.yml")
	if err != nil {
		return
	}
	fmt.Println(file)
}
