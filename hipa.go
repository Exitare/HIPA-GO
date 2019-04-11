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

	success := ParseFiles(w, r)

	if success {
		prepareFiles()
		startCalculations()
	}
}

func prepareFiles() {
	for _, inputFile := range inputFiles {
		inputFile.getFileName()
		fmt.Printf("Filename is %s stored in folder %s\n", inputFile.Name, inputFile.Foldername)
	}
}

func startCalculations() {

}
