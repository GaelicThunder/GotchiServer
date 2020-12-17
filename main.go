package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var database Database

func main() {
	lambda.Start(GotchiStatus)
}

// GotchiStatus is the gotchi core logic which permit to do the following action
// GET /myGotchiID: return a list of the gotchi seen
// PUT /myGotchiID/gotchiID: permit to save a new gotchi
func GotchiStatus(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if database == nil {
		database = NewDynamoDB()
	}
	if request.HTTPMethod == "GET" {
		return getListOfKnowDevice(&request)
	}
	if request.HTTPMethod == "POST" {
		return storeNewGotchi(&request)
	}
	var buf bytes.Buffer
	body, err := json.Marshal(map[string]interface{}{
		"message": fmt.Sprintf("Oh hello there, I just receive your request with method %s and with this body: '%s' path: '%s'", request.HTTPMethod, request.Body, request.Path),
	})
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
			"X-Handler":    "GotchiStatus-handler",
		},
	}

	return resp, nil
}

// getListOfKnowDevice return the know device for the specific gotchiID in the path
func getListOfKnowDevice(request *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var buf bytes.Buffer
	splitted := strings.Split(request.Path, "/")
	if len(splitted) != 2 {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("no valid path")
	}
	knowGotchi, err := database.GetGotchi(splitted[1])
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}
	body, err := json.Marshal(map[string]interface{}{
		"know_gotchi": knowGotchi,
	})
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}
	json.HTMLEscape(&buf, body)
	return events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
			"X-Handler":    "GotchiStatus-handler",
		},
	}, nil
}

// storeNewGotchi store a gotchiID to the give gothi
func storeNewGotchi(request *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	splitted := strings.Split(request.Path, "/")
	if len(splitted) != 3 {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("no valid path")
	}
	err := database.SaveGotchi(splitted[1], splitted[2])
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type": "application/json",
			"X-Handler":    "GotchiStatus-handler",
		},
	}, nil
}
