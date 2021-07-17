package handler

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/BrayanAriasH/bp_microservice_exif_info/src/model"
	"github.com/BrayanAriasH/bp_microservice_exif_info/src/services"
	"github.com/BrayanAriasH/bp_microservice_exif_info/src/util"
	"github.com/aws/aws-lambda-go/events"
)

func createResponseError(err error) events.APIGatewayProxyResponse {
	errorResponse := model.NewErrorResponse(http.StatusBadGateway, err)
	jsonString, err := json.Marshal(errorResponse)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Body:       string(jsonString),
	}
}

func CreatePhoto(request events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {
	requestPhoto := model.NewRequestPhoto()
	err = json.Unmarshal([]byte(request.Body), requestPhoto)
	if err != nil {
		return createResponseError(err), err
	}
	fileBytes, err := base64.StdEncoding.DecodeString(requestPhoto.Base64Content)
	if err != nil {
		return createResponseError(err), err
	}
	photo, err := model.CreatePhotoFromFile(fileBytes)
	if err != nil {
		return createResponseError(err), err
	}
	compressedFile, err := util.CreateCompressedImage(fileBytes)
	if err != nil {
		return createResponseError(err), err
	}
	err = services.UploadImage(fileBytes, compressedFile, photo.Id)
	if err != nil {
		return createResponseError(err), err
	}
	err = services.WritePhoto(photo)
	if err != nil {
		return createResponseError(err), err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       photo.String(),
	}, nil
}
