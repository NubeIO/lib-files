package fileutils

import (
	"errors"
	"io"
	"os"
	"path/filepath"
)

// Copy copies a file or directory from src to dst. If it is
// a directory, all the files and subdirectories will be copied.
func Copy(src, dst string) error {
	if src = resolve(src); src == "" {
		return os.ErrNotExist
	}
	if dst = resolve(dst); dst == "" {
		return os.ErrNotExist
	}
	if dst == src {
		return os.ErrInvalid
	}
	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return CopyDir(src, dst)
	}
	return CopyFile(src, dst)
}

// CopyFile copies a file from source to dest and returns
// an error if any.
func CopyFile(source string, dest string) error {
	// Open the source file.
	src, err := os.Open(source)
	if err != nil {
		return err
	}
	defer src.Close()

	srcinfo, err := os.Stat(filepath.Dir(source))
	if err != nil {
		return err
	}

	// Makes the directory needed to create the dst file.
	err = os.MkdirAll(filepath.Dir(dest), srcinfo.Mode())
	err = os.Chmod(filepath.Dir(dest), srcinfo.Mode())
	if err != nil {
		return err
	}

	// Create the destination file.
	dst, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy the contents of the file.
	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	// Copy the mode if the user can't
	// open the file.
	info, err := os.Stat(source)
	if err != nil {
		return err
	} else {
		err = os.Chmod(dest, info.Mode())
		if err != nil {
			return err
		}
	}

	return nil
}

// CopyDir copies a directory from source to dest and all
// of its subdirectories. It doesn't stop if it finds an error
// during the copy. Returns an error if any.
func CopyDir(source string, dest string) error {
	// Get properties of source.
	srcinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// Create the destination directory.
	err = os.MkdirAll(dest, srcinfo.Mode())
	if err != nil {
		return err
	}

	dir, _ := os.Open(source)
	obs, err := dir.Readdir(-1)
	if err != nil {
		return err
	}

	var errs []error

	for _, obj := range obs {
		fsource := source + "/" + obj.Name()
		fdest := dest + "/" + obj.Name()

		if obj.IsDir() {
			// Create sub-directories, recursively.
			err = CopyDir(fsource, fdest)
			if err != nil {
				errs = append(errs, err)
			}
		} else {
			// Perform the file copy.
			err = CopyFile(fsource, fdest)
			if err != nil {
				errs = append(errs, err)
			}
		}
	}

	var errString string
	for _, err := range errs {
		errString += err.Error() + "\n"
	}

	if errString != "" {
		return errors.New(errString)
	}

	return nil
}
