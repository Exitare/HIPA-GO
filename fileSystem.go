package main

import (
	"fmt"
	"io"
	"net/http"
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
func createDirectory() (string, error) {
	folderName := GenerateRandomString(20)
	exists, err := dirExists(workingDir + folderName)

	if err != nil {
		return "", err
	}

	fmt.Println("Folder exists ", exists)
	if !exists {
		os.MkdirAll(workingDir+folderName, os.FileMode.Perm(0777))
	} else {
		createDirectory()
	}

	return folderName, err
}

//ParseFiles parse the uploaded files per request
func ParseFiles(w http.ResponseWriter, r *http.Request) bool {

	// TODO create Folder
	folderName, err := createDirectory()

	if err != nil {
		fmt.Println("Could not create directory")
		return false
	}

	err = r.ParseMultipartForm(32 << 10)

	if err != nil {
		fmt.Println("Could not parse file")
		return false
	}

	m := r.MultipartForm

	files := m.File["files"]

	for i := range files {
		file, err := files[i].Open()
		fmt.Println("Received file: ", files[i].Filename)
		defer file.Close()

		if err != nil {
			fmt.Println("Error while open file")
			return false
		}

		dst, err := os.Create(workingDir + folderName + "/" + files[i].Filename)

		defer dst.Close()

		if err != nil {
			fmt.Println("Could not create file")
			return false
		}

		if _, err := io.Copy(dst, file); err != nil {
			fmt.Println("Could not copy file")
			return false
		}

		var Cells []Cell
		inputFile := InputFile{files[i].Filename, "", folderName, 0, 0.0, Cells, 0, 0, 0, 0, 0}

		inputFiles = append(inputFiles, inputFile)
		fmt.Printf("Uploaded %s succesfully\n", files[i].Filename)

	}
	return true
}
