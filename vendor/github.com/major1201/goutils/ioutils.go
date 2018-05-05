package goutils

import (
	"os"
)

// IsExist is used to determine whether a path exists or not
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// IsFile tells a path is a file or not
func IsFile(path string) bool {
	fi, err := os.Stat(path)
	if err == nil && fi.Mode().IsRegular() == true {
		return true
	}
	return false
}

// IsDir tells a path is a directory or not
func IsDir(path string) bool {
	fi, err := os.Stat(path)
	if err == nil && fi.Mode().IsDir() == true {
		return true
	}
	return false
}
