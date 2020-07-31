package main

import (
	"log"
	"net/http"
	"routes"
)



type ResultPage struct {
	Name     string
	fileName string
	LOC      uint32
}

func main() {

	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}


func setupRoutes() {
	http.HandleFunc("/upload/", routes.UploadFile)
	http.HandleFunc("/", routes.ServeMainPage)
	http.ListenAndServe(":8080", nil)
}

