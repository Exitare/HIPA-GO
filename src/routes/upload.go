package routes

import (
	File "entities/file"
	"fmt"
	"net/http"
)

func UploadRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	err := File.ParseFiles(w, r)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = File.Prepare()
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Error while processing files", http.StatusInternalServerError)
		return
	}


	defer r.Body.Close()
	// startCalculations()
}

