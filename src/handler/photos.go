package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/BrayanAriasH/bp_microservice_exif_info/src/model"
	"github.com/BrayanAriasH/bp_microservice_exif_info/src/services"
	"github.com/BrayanAriasH/bp_microservice_exif_info/src/util"
)

func createResponseError(err error, writer http.ResponseWriter) {
	log.Fatalf("Error: %v", err)
	writer.WriteHeader(http.StatusInternalServerError)
	writer.Write([]byte(fmt.Sprintf("500 - %s", err.Error())))
}

func CreatePhoto(writer http.ResponseWriter, request *http.Request) {
	log.Println("Handling request...")
	requestPhoto := model.NewRequestPhoto()
	err := json.NewDecoder(request.Body).Decode(requestPhoto)
	if err != nil {
		createResponseError(err, writer)
	}
	log.Println("file-name:", requestPhoto.FileName)
	fileBytes, err := base64.StdEncoding.DecodeString(requestPhoto.Base64Content)
	if err != nil {
		createResponseError(err, writer)
	}
	photo, err := model.CreatePhotoFromFile(fileBytes)
	if err != nil {
		createResponseError(err, writer)
	}
	compressedFile, err := util.CreateCompressedImage(fileBytes)
	if err != nil {
		createResponseError(err, writer)
	}
	err = services.UploadImage(fileBytes, compressedFile, photo.Id)
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
