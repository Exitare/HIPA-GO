package main

import (
	"bytes"
	"fmt"
	"io"
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
	var Buf bytes.Buffer
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	fmt.Printf("Content Type: %s", handler.Header["Content-Type"])

	contentType := handler.Header["Content-Type"]

	if contentType[0] == "text/plain" {
		fmt.Printf("Content type is txt file.")
		io.Copy(&Buf, file)
		// do something with the contents...
		// I normally have a struct defined and unmarshal into a struct, but this will
		// work as an example
		content := Buf.String()

		inputFile := InputFile{
			Name:    handler.Filename,
			Content: content,
		}

		inputFiles = append(inputFiles, inputFile)
	}

	startCalculations()
	for _, inputFile := range inputFiles {
		fmt.Println(inputFile.Name)
		fmt.Printf("Addr: %p\n", &inputFile)
	}

}

func showResultPage(w http.ResponseWriter) {
	tmpl := template.Must(template.ParseFiles("./static/result.html"))
	data := ResultPage{
		Name:     "Raphael",
		fileName: "Testpage",
		LOC:      1500,
	}

	tmpl.Execute(w, data)
}

func startCalculations() {
	for _, inputFile := range inputFiles {
		inputFile.countLines()
	}
}
