package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
	"os"
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
	fmt.Println("ich bin ein Test")
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	err := r.ParseMultipartForm(32 << 20)

	if err != nil {
		fmt.Println("error0")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}



	//get a ref to the parsed multipart form
	m := r.MultipartForm

	//get the *fileheaders
	files := m.File["files"]

	for i, _ := range files {
		//for each fileheader, get a handle to the actual file
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			fmt.Println("error1")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//create destination file making sure the path is writeable.
		fmt.Println(files[i].Filename)
		dst, err := os.Create("./data/" + files[i].Filename)
		defer dst.Close()
		if err != nil {
			fmt.Println("error2")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//copy the uploaded file to the destination file
		if _, err := io.Copy(dst, file); err != nil {
			fmt.Println("error3")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
	}
	fmt.Println("Upload successful.")
	



	//startCalculations()
	//for _, inputFile := range inputFiles {
	//	fmt.Println(inputFile.Name)
//		fmt.Printf("Addr: %p\n", &inputFile)
//	}

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
