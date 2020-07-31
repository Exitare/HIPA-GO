package hipafile

import (
	Cell "entities/cell"
	"services/FileManagement"
	"fmt"
	"io"
	"net/http"
	"os"
)

//ParseFiles parse the uploaded files per request
func ParseFiles(w http.ResponseWriter, r *http.Request) error {

	// TODO create Folder
	folderName, err := FileManagement.CreateDirectory()

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

		var Cells []*Cell.Cell
		inputFile := InputFile{"", "", folderName, 0.0, Cells, 0, 0, 0, 0, 0}
		InputFiles = append(InputFiles, &inputFile)

		inputFile.ResolveName(files[i].Filename)

		fmt.Println("Received file: ", inputFile.Name)
		defer file.Close()

		if err != nil {
			return err
		}

		dst, err := os.Create(FileManagement.WorkingDir + folderName + "/" + inputFile.Name + ".txt")

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

