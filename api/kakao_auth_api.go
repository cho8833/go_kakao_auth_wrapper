package api

import (
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"log"
)

func RedirectAuthorizeCode(event events.APIGatewayProxyRequest) (string, error) {
	errorDescription, err := event.QueryStringParameters["error_description"]
	if err != true && errorDescription != "" {
		log.Printf(errorDescription)
		return "", errors.New(errorDescription)
	}
	code, _ := event.QueryStringParameters["code"]
	return code, nil
}
