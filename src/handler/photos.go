package handler

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/BrayanAriasH/bp_microservice_exif_info/src/model"
	"github.com/BrayanAriasH/bp_microservice_exif_info/src/services"
	"github.com/BrayanAriasH/bp_microservice_exif_info/src/util"
)

func createResponseError(err error, writer http.ResponseWriter) {
	log.Fatalf("Error: %v", err)
	writer.WriteHeader(http.StatusBadRequest)
	writer.Write([]byte(fmt.Sprintf("500 - %s", err.Error())))
}

func CreatePhoto(writer http.ResponseWriter, request *http.Request) {
	log.Println("Handling request...")
	request.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := request.FormFile("the-file")
	if err != nil {
		log.Fatalf("Error retrieving the file: %v", err)
		return
	}
	if err != nil {
		createResponseError(err, writer)
	}
	defer file.Close()
	log.Printf("Uploaded File: %+v\n", handler.Filename)
	log.Printf("File Size: %+v\n", handler.Size)
	log.Printf("MIME Header: %+v\n", handler.Header)
	fileBytes := bytes.NewBuffer(nil)
	if _, err := io.Copy(fileBytes, file); err != nil {
		createResponseError(err, writer)
	}
	if fileBytes == nil {
		log.Fatal("The filebytes is nil.")
		createResponseError(err, writer)
	}
	photo, err := model.CreatePhotoFromFile(fileBytes.Bytes())
	if err != nil {
		createResponseError(err, writer)
	}
	compressedFile, err := util.CreateCompressedImage(fileBytes.Bytes())
	if err != nil {
		createResponseError(err, writer)
	}
	err = services.UploadImage(fileBytes.Bytes(), compressedFile, photo.Id)
	if err != nil {
		createResponseError(err, writer)
	}
	err = services.WritePhoto(photo)
	if err != nil {
		createResponseError(err, writer)
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte(photo.String()))
}
