package FileManagement

import (
	"fmt"
	"globals"
	"os"
	"services/Misc"
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
func createDirectory() (string, error) {
	folderName := Misc.GenerateRandomString(20)
	exists, err := dirExists(globals.WorkingDir + folderName)

	if err != nil {
		return "", err
	}

	fmt.Println("Folder exists ", exists)
	if !exists {
		os.MkdirAll(globals.WorkingDir+folderName, os.FileMode.Perm(0777))
	} else {
		createDirectory()
	}

	return folderName, err
}

