package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type MainPage struct {
	FileName string
	LOC      uint32
}

func main() {
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveMainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/main.html"))
	data := MainPage{
		FileName: "",
		LOC:      0,
	}

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
	var count uint32
	if contentType[0] == "text/plain" {
		fmt.Printf("Content type is txt file.")
		io.Copy(&Buf, file)
		// do something with the contents...
		// I normally have a struct defined and unmarshal into a struct, but this will
		// work as an example
		contents := Buf.String()
		tests, err := Buf.ReadString('\n')
		if err == nil {
			fmt.Println(tests)
		}

		lines := strings.Split(contents, "\n")

		for _, element := range lines {
			count++

			fmt.Println(element)
			// index is the index where we are
			// element is the element from someSlice for where we are
		}

		fmt.Println("Count is", count)
	}

	tmpl := template.Must(template.ParseFiles("./static/main.html"))
	data := MainPage{
		FileName: handler.Filename,
		LOC:      count,
	}

	tmpl.Execute(w, data)
}
