package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestGotchiHandler(t *testing.T) {
	ctx := context.TODO()
	request := events.APIGatewayProxyRequest{
		HTTPMethod: "POST",
		Body:       "example_body",
	}
	event, err := GotchiStatus(ctx, request)
	if err != nil {
		t.Fatalf("Error: %s\n", err.Error())
	}
	expectedResponse := "{\"message\":\"Oh hello there, I just recive your request with method POST and with this body: 'example_body' path: ''\"}"
	if event.Body != expectedResponse {
		t.Fatalf("Wrong body returned, expected %s but having %s\n", expectedResponse, event.Body)
	}
}

func TestGotchiHandlerGET(t *testing.T) {
	ctx := context.TODO()
	request := events.APIGatewayProxyRequest{
		HTTPMethod: "GET",
		Path:       "/gotchi1",
	}
	event, err := GotchiStatus(ctx, request)
	if err != nil {
		t.Fatalf("Error: %s\n", err.Error())
	}
	expectedResponse := "{\"know_gotchi\":[\"id1\",\"id2\",\"id3\"]}"
	if event.Body != expectedResponse {
		t.Fatalf("Wrong body returned, expected %s but having %s\n", expectedResponse, event.Body)
	}
}
