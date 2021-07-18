package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BrayanAriasH/bp_microservice_exif_info/src/handler"
)

func main() {

	http.HandleFunc("/create-image", handler.CreatePhoto)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"response": "pong ;)"}`)
	})
	log.Println(">>>>Starting server in port 1503")
	log.Fatal(http.ListenAndServe(":1503", nil))

}
