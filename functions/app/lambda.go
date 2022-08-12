package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"

	"homadrone/app/graphql"
)

func NewLambda(
	ctx context.Context,
	request events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod != http.MethodPost {
		return events.APIGatewayProxyResponse{
			Body:       "only POST method is accepted",
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	fmt.Printf(
		"Processing request for ID: %s.\n",
		request.RequestContext.RequestID,
	)

	res, err := graphql.NewGraphQLExecution(ctx, graphql.Schema, request.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	body, _ := json.Marshal(res.Data)
	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: http.StatusOK,
	}, nil
}
