package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func copyConfig(src, dest string) error {
	source := readPath(src)
	destination := readPath(dest)

	isSrcDir, err := isDirectory(source)
	if err != nil {
		return err
	}
	if isSrcDir {
		return copyDir(source, destination)
	}

	// Create dir if destination folders if they dont exist
	if !exists(destination) {
		createDir(filepath.Dir(destination), os.ModePerm)
	}

	return copyFile(source, destination)
}

// Recursively copy the contents of the source directory and
// create it in the destination directory
func copyDir(src, dest string) error {
	dirs, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, dir := range dirs {
		srcPath := filepath.Join(src, dir.Name())
		destPath := filepath.Join(dest, dir.Name())

		srcInfo, err := os.Stat(srcPath)
		if err != nil {
			return err
		}

		switch srcInfo.Mode() & os.ModeType {
		case os.ModeDir:
			if err := createDir(destPath, 0755); err != nil {
				return err
			}
			if err := copyDir(srcPath, destPath); err != nil {
				return err
			}
		case os.ModeSymlink:
			fmt.Errorf("Symlink is not currently supported")
		default:
			// Create dir if destination folders if they dont exist
			if !exists(destPath) {
				createDir(filepath.Dir(destPath), os.ModePerm)
			}

			if err := copyFile(srcPath, destPath); err != nil {
				return err
			}
		}
	}
	return nil
}

func copyFile(src, dest string) error {
	srcStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !srcStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	// Create a file with permission os.ModePerm
	// If create the file if it doesn't exist
	// Overwrites if file is exists
	destination, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

func exists(p string) bool {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return false
	}
	return true
}

func createDir(p string, perm os.FileMode) error {
	if exists(p) {
		return nil
	}

	if err := os.MkdirAll(p, perm); err != nil {
		return err
	}

	return nil
}

func isDirectory(path string) (bool, error) {
	pathInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return pathInfo.IsDir(), err
}
