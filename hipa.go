package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
	fmt.Println(r.Header.Get("X-Forwarded-For"))
	fmt.Println("Client IP is ", r.Header.Get("X-Forwarded-For"))

	// TODO create Folder
	folderName, err := CreateDirectory()

	if err != nil {
		fmt.Println("Could not create directory")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = r.ParseMultipartForm(32 << 10)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	m := r.MultipartForm

	files := m.File["files"]

	for i := range files {
		file, err := files[i].Open()
		fmt.Println("Received file: ", files[i].Filename)
		defer file.Close()

		if err != nil {
			fmt.Println("Error while open file")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dst, err := os.Create("./data/" + folderName + "/" + files[i].Filename)

		defer dst.Close()

		if err != nil {
			fmt.Println("Could not create file")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := io.Copy(dst, file); err != nil {
			fmt.Println("Could not copy file")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("Upload succesfull")

	}
}

func startCalculations() {
	for _, inputFile := range inputFiles {
		inputFile.countLines()
	}
}
