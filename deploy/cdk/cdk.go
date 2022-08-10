package main

import (
	"cdk/stack"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func main() {
	app := awscdk.NewApp(nil)

	stack.NewAPIGateWayStack(
		app,
		"ServerlessGraphQLAPIStack",
		nil,
	)

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	//return nil
	return &awscdk.Environment{
		Account: jsii.String("267519217152"),
		Region:  jsii.String("eu-west-2"),
	}
}
