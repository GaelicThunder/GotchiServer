package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(GotchiStatus)
}

// GotchiStatus is the gotchi core logic which permit to do the following action
// GET /myGotchiID: return a list of the gotchi seen
// PUST /myGotchiID/gotchiID: permit to save a new gotchi
func GotchiStatus(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var buf bytes.Buffer
	body, err := json.Marshal(map[string]interface{}{
		"message": fmt.Sprintf("Oh hello there, I just recive your request with method %s and with this body: %s", request.HTTPMethod, request.Body),
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
