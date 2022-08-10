package lambda

import (
	"github.com/aws/aws-lambda-go/lambda"

	"homadrone/app/graphql"
)

func main() {
	lambda.Start(graphql.NewLambda)
}
