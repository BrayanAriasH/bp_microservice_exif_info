package handler

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/BrayanAriasH/bp_microservice_exif_info/src/model"
	"github.com/BrayanAriasH/bp_microservice_exif_info/src/services"
	"github.com/BrayanAriasH/bp_microservice_exif_info/src/util"
)

func createResponseError(err error, tag string, writer http.ResponseWriter) {
	log.Printf("Error on CreatePhoto in section %s: %v", tag, err)
	writer.WriteHeader(http.StatusInternalServerError)
}

func CreatePhoto(writer http.ResponseWriter, request *http.Request) {
	log.Println("Handling request...")
	request.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := request.FormFile("the-file")
	if err != nil {
		log.Printf("Error retrieving the file: %v", err)
		createResponseError(err, "retrieving the file", writer)
	}
	defer file.Close()
	log.Printf("Uploaded File: %+v\n", handler.Filename)
	log.Printf("File Size: %+v\n", handler.Size)
	log.Printf("MIME Header: %+v\n", handler.Header)
	fileBytes := bytes.NewBuffer(nil)
	if _, err := io.Copy(fileBytes, file); err != nil {
		createResponseError(err, "copying files", writer)
		return
	}
	photo, err := model.CreatePhotoFromFile(fileBytes.Bytes())
	if err != nil {
		createResponseError(err, "CreatePhotoFromFile", writer)
		return
	}
	compressedFile, err := util.CreateCompressedImage(fileBytes.Bytes())
	if err != nil {
		createResponseError(err, "CreateCompressedImage", writer)
		return
	}
	err = services.UploadImage(fileBytes.Bytes(), compressedFile, photo.Id)
	if err != nil {
		createResponseError(err, "UploadImage", writer)
		return
	}
	err = services.WritePhoto(photo)
	if err != nil {
		createResponseError(err, "WritePhoto", writer)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte(photo.String()))
	log.Println("Image created with id", photo.Id)
}
