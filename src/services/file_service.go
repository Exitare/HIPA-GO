package services

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"entities/cell"
	"entities/file"
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
func ParseFiles(w http.ResponseWriter, r *http.Request) error {

	// TODO create Folder
	folderName, err := createDirectory()

	if err != nil {
		return err
	}

	err = r.ParseMultipartForm(32 << 10)

	if err != nil {
		return err
	}

	m := r.MultipartForm

	files := m.File["files"]
	fmt.Println(r.Form["new_data"])

	for i := range files {
		file, err := files[i].Open()

		var Cells []*Cell
		inputFile := InputFile{"", "", folderName, 0.0, Cells, 0, 0, 0, 0, 0}
		inputFiles = append(inputFiles, &inputFile)

		inputFile.resolveName(files[i].Filename)

		fmt.Println("Received file: ", inputFile.Name)
		defer file.Close()

		if err != nil {
			return err
		}

		dst, err := os.Create(workingDir + folderName + "/" + inputFile.Name + ".txt")

		defer dst.Close()

		if err != nil {
			return err
		}

		if _, err := io.Copy(dst, file); err != nil {
			return err
		}

		fmt.Printf("Uploaded %s succesfully\n", inputFile.Name)

	}
	return nil
}
