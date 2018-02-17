package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"

	"github.com/markwilson/markwilson.me-get-in-touch/api"
	"github.com/markwilson/markwilson.me-get-in-touch/mail"
)

type Event map[string]interface{}

func HandleRequest(rawRequest events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing Lambda request %s\n", rawRequest.RequestContext.RequestID)

	if len(rawRequest.Body) < 1 {
		return events.APIGatewayProxyResponse{}, api.ErrorNoBody
	}

	request, err := api.RequestFromBody(rawRequest.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	err = mail.SendEmail(request)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return api.Response(), nil
}

func main() {
	lambda.Start(HandleRequest)
}
