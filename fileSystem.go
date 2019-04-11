package main

import (
	"fmt"
	"os"
)

// dirExists returns whether the given file or directory exists
func dirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// CreateDirectory creates a new directory
func CreateDirectory() (string, error) {
	folderName := GenerateRandomString(20)
	exists, err := dirExists("./data/" + folderName)

	if err != nil {
		return "", err
	}

	fmt.Println("Folder exists ", exists)
	if !exists {
		os.MkdirAll("./data/"+folderName, os.FileMode.Perm(0777))
	} else {
		CreateDirectory()
	}

	return folderName, err
}
