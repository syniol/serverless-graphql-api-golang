package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	lambda2 "homadrone/app"
)

func main() {
	lambda.Start(lambda2.NewLambda)
}
