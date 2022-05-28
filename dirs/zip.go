package fileutils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func (d Dir) Extract(source, destination string, perm os.FileMode) ([]string, error) {
	r, err := zip.OpenReader(source)
	if err != nil {
		return nil, err
	}
	if perm == 0 {
		perm = 700
	}

	defer func() {
		if err = r.Close(); err != nil {
			panic(err)
		}
	}()

	err = os.MkdirAll(destination, perm)
	if err != nil {
		return nil, err
	}

	var extractedFiles []string
	for _, f := range r.File {
		err := extractAndWriteFile(destination, f)
		if err != nil {
			return nil, err
		}

		extractedFiles = append(extractedFiles, f.Name)
	}

	return extractedFiles, nil
}

func extractAndWriteFile(destination string, f *zip.File) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer func() {
		if err = rc.Close(); err != nil {
			panic(err)
		}
	}()

	path := filepath.Join(destination, f.Name)
	if !strings.HasPrefix(path, filepath.Clean(destination)+string(os.PathSeparator)) {
		return fmt.Errorf("%s: illegal file path", path)
	}
	if f.FileInfo().IsDir() {
		err = os.MkdirAll(path, f.Mode())
		if err != nil {
			return err
		}
	} else {
		err = os.MkdirAll(filepath.Dir(path), f.Mode())
		if err != nil {
			return err
		}

		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer func() {
			if err = f.Close(); err != nil {
				panic(err)
			}
		}()

		_, err = io.Copy(f, rc)
		if err != nil {
			return err
		}
	}

	return nil
}
