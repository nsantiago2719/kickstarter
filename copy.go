package main

import (
	"fmt"
	"io"
	"os"
)

func copyConfig(src, d string) error {
	source := readPath(src)
	destination := readPath(d)

	isSrcDir, err := isDirectory(source)
	if err != nil {
		return err
	}
	if isSrcDir {
		copyDir(source, destination)
	} else {
		fmt.Println("copying: ", source)
		copyFile(source, destination)
	}
	return nil
}

func copyDir(_, _ string) error {
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

	// Create a file with permission 0644
	// If create the file if it doesn't exist
	// Overwrites if file is exists
	destination, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
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
