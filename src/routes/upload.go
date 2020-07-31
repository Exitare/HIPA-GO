package routes

import (
	"fmt"
	"log"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	err := ParseFiles(w, r)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = prepareFiles()
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Error while processing files", http.StatusInternalServerError)
		return
	}

	startCalculations()
}