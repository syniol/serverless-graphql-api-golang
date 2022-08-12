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
	fmt.Printf(
		"Processing request for ID: %s.\n",
		request.RequestContext.RequestID,
	)

	if request.HTTPMethod != http.MethodPost {
		return events.APIGatewayProxyResponse{
			Body:       "only POST method is accepted",
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	res, err := graphql.NewGraphQLExecution(ctx, graphql.Schema, request.Body)
	if err != nil {
		fmt.Println("error at NewGraphQLExecution method", err.Error())

		return events.APIGatewayProxyResponse{
			Body:       "malformed request",
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	body, err := json.Marshal(res.Data)
	if err != nil {
		fmt.Println("error marshal graphql response.Data", err.Error())

		return events.APIGatewayProxyResponse{
			Body:       "error transforming result back to client",
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: http.StatusOK,
	}, nil
}
