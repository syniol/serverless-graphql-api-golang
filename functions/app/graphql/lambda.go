package graphql

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func NewLambda(
	_ context.Context,
	request events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	fmt.Printf(
		"Processing request data for request %s.\n",
		request.RequestContext.RequestID,
	)
	fmt.Printf("Body size = %d.\n", len(request.Body))

	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("%s: %s\n", key, value)
	}

	return events.APIGatewayProxyResponse{
		Body:       request.Body,
		StatusCode: http.StatusOK,
	}, nil
}
