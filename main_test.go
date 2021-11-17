package main

import (
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestFullFlowWithMock(t *testing.T) {
	os.Setenv("mode_mock", "true")
	processing(events.APIGatewayProxyRequest{})
}
