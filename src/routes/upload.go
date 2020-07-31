package routes

import (
	"fmt"
	"net/http"
	"services/FileManagement"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	err := FileManagement.ParseFiles(w, r)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = FileManagement.Prepare()
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Error while processing files", http.StatusInternalServerError)
		return
	}

	// startCalculations()
}