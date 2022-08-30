package main

import (
	"fmt"
	"github.com/NubeIO/lib-dirs/fileutils"
)

func main() {
	err := fileutils.New().RecursiveZip("/data", "./test.zip")
	fmt.Println(err)
	if err != nil {
		return
	}
}
