package main

import (
	"log"
	"net/http"

	"github.com/BrayanAriasH/bp_microservice_exif_info/src/router"
)

func main() {

	http.HandleFunc("/", router.Route)
	log.Println("Starting server in port 1503...")
	log.Fatal(http.ListenAndServe(":1503", nil))
}
