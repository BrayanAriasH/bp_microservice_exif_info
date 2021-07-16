package handler

import "github.com/aws/aws-lambda-go/events"

func Route(request events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {
	return CreatePhoto(request)
}
