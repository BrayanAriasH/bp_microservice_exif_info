package main

import (
	"github.com/BrayanAriasH/bp_microservice_exif_info/src/handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.Route)
}
