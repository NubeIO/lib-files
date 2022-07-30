package main

import (
	"fmt"
	fileutils "github.com/NubeIO/lib-dirs/dirs"
)

func main() {
	err := fileutils.New().RecursiveZip("/data", "./test.zip")
	fmt.Println(err)
	if err != nil {
		return
	}
}
