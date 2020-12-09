package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestGotchiHandler(t *testing.T) {
	ctx := context.TODO()
	request := events.APIGatewayProxyRequest{
		HTTPMethod: "post",
		Body:       "example_body",
	}
	event, err := GotchiStatus(ctx, request)
	if err != nil {
		t.Fatalf("Error: %s\n", err.Error())
	}
	expectedResponse := "{\"message\":\"Oh hello there, I just recive your request with method post and with this body: example_body\"}"
	if event.Body != expectedResponse {
		t.Fatalf("Wrong body returned, expected %s but having %s\n", expectedResponse, event.Body)
	}
}
