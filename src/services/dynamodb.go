package services

import (
	"fmt"
	"time"

	"github.com/BrayanAriasH/bp_microservice_exif_info/src/constant"
	"github.com/BrayanAriasH/bp_microservice_exif_info/src/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var svc = dynamodb.New(sess)

func createItem(photo *model.Photo) (item map[string]*dynamodb.AttributeValue) {
	item = map[string]*dynamodb.AttributeValue{
		"id": {
			S: aws.String(photo.Id),
		},

		"make": {
			S: aws.String(photo.MakeWith),
		},
		"date_time_original": {
			S: aws.String(photo.DateTimeOriginal.String()),
		},
		"creation_date": {
			S: aws.String(time.Now().String()),
		},
		"date_time_unix": {
			N: aws.String(fmt.Sprintf("%v", time.Now().Unix())),
		},
		"exposure_time": {
			S: aws.String(photo.ExposureTime),
		},
		"orientation": {
			S: aws.String(photo.Orientation),
		},
		"model": {
			S: aws.String(photo.Model),
		},
		"exposure_mode": {
			S: aws.String(photo.Model),
		},
		"f_number": {
			N: aws.String(fmt.Sprintf("%v", photo.FNumber)),
		},
		"latitude": {
			N: aws.String(fmt.Sprintf("%v", photo.Latitude)),
		},
		"longitude": {
			N: aws.String(fmt.Sprintf("%v", photo.Longitude)),
		},
		"x_resolution": {
			N: aws.String(fmt.Sprintf("%v", photo.PixelXDimension)),
		},
		"y_resolution": {
			N: aws.String(fmt.Sprintf("%v", photo.PixelYDimension)),
		},
		"iso": {
			N: aws.String(fmt.Sprintf("%v", photo.ISO)),
		},
	}
	return item
}

func WritePhoto(photo *model.Photo) (err error) {
	input := &dynamodb.PutItemInput{
		Item:                   createItem(photo),
		ReturnConsumedCapacity: aws.String("TOTAL"),
		TableName:              aws.String(constant.BPImagesTableName),
	}
	_, err = svc.PutItem(input)
	if err != nil {
		return err
	}
	return nil
}
