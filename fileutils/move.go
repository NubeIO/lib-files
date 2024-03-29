package fileutils

import (
	"fmt"
	"io"
	"os"
)

func MoveFile(sourcePath, destPath string) error {
	err := CheckDelete(sourcePath)
	if err != nil {
		return err
	}
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("writing to output file failed: %s", err)
	}
	info, err := os.Stat(sourcePath)
	if err != nil {
		return err
	}
	if err := os.Chmod(destPath, info.Mode()); err != nil {
		return err
	}
	// The copy was successful, so now delete the original file
	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("failed removing original file: %s", err)
	}
	return nil
}
