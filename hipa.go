package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type MainPage struct {
}

type ResultPage struct {
	Name     string
	fileName string
	LOC      uint32
}

func main() {

	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveMainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/main.html"))
	data := MainPage{}

	tmpl.Execute(w, data)
}

func setupRoutes() {
	http.HandleFunc("/upload/", uploadFile)
	http.HandleFunc("/", serveMainPage)
	http.ListenAndServe(":8080", nil)
}

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

func prepareFiles() error {
	for _, inputFile := range inputFiles {
		err := inputFile.readContent()
		if err != nil {
			return err
		}

		inputFile.countLines()
		fmt.Println(inputFile.Lines)
	}

	return nil
}

func startCalculations() {

}
