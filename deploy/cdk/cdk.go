package main

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/jsii-runtime-go"

	"cdk/stack"
)

func main() {
	app := awscdk.NewApp(&awscdk.AppProps{
		Context: &map[string]interface{}{},
	})

	stack.NewServerlessAPIStack(
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
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
