package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	appLambda "homadrone/app"
)

func main() {
	lambda.Start(appLambda.NewLambda)
}
