package router

import (
	"fmt"
	"net/http"

	"github.com/BrayanAriasH/bp_microservice_exif_info/src/handler"
)

func write404NotFound(writer http.ResponseWriter, path string) {
	writer.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(writer, "The path %s doesn't exists. Please contact the administrator.", path)
}

func writeInfoMicroservice(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(writer, `{"app": "bp-create-images"}`)
}

func Route(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	method := request.Method

	if method == "GET" {
		switch path {
		case "/":
			writeInfoMicroservice(writer)
		case "/ping":
			handler.Ping(writer, request)
		default:
			write404NotFound(writer, path)
		}
	} else if method == "POST" {
		switch path {
		case "/create-image":
			handler.CreatePhoto(writer, request)
		default:
			write404NotFound(writer, path)
		}
	} else {
		write404NotFound(writer, path)
	}
}
