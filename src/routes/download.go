package routes

import (
	"fmt"
	"log"
	"net/http"
)

func DownloadRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Download  Endpoint Hit")
	switch r.Method {
	case "GET":
		handleGet(w, r)
	default:
		break
	}


	defer r.Body.Close()
}


func handleGet(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["file"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'file' is missing")
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	key := keys[0]

	log.Println("Url Param 'key' is: " + string(key))
}
