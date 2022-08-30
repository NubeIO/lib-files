package fileutils

import (
	"fmt"
	"testing"
)

func TestDirs_ReadFile(t *testing.T) {
	file, err := ReadAllOk("/data/flow-framework/config/config.yml")
	if err != nil {
		return
	}
	fmt.Println(file)
}
