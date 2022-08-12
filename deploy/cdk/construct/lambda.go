package construct

import (
	"fmt"
	"path/filepath"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	iam "github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	lambda "github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	s3assets "github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/jsii-runtime-go"
)

const AWSLambdaExecutionPath = "/asset-output/"

type LambdaConstruct struct {
	LambdaFunction lambda.Function
	Scope          awscdk.Stack
}

func NewLambdaConstruct(
	scope awscdk.Stack,
	name string,
) LambdaConstruct {
	iamRole := iam.NewRole(
		scope,
		jsii.String(fmt.Sprintf("%sLambdaRole", name)),
		&iam.RoleProps{
			AssumedBy: iam.NewServicePrincipal(
				jsii.String("lambda.amazonaws.com"),
				nil,
			),
		},
	)

	graphQLAPIExecutorFunction := lambda.NewFunction(
		scope,
		jsii.String(fmt.Sprintf("%sLambdaConstruct", name)),
		&lambda.FunctionProps{
			FunctionName: jsii.String("GraphQLAPIExecutorFunction"),
			Runtime:      lambda.Runtime_GO_1_X(),
			Handler:      jsii.String("main"),
			Code: lambda.AssetCode_FromAsset(
				jsii.String(filepath.Join(".", "..", "..", "functions")),
				&s3assets.AssetOptions{
					Bundling: &awscdk.BundlingOptions{
						Image: lambda.Runtime_GO_1_X().BundlingImage(),
						User:  jsii.String("root"),
						Command: &[]*string{
							jsii.String("bash"),
							jsii.String("-c"),
							jsii.String(fmt.Sprintf(
								"go mod vendor && go build -o %smain ./cmd/lambda/main.go",
								AWSLambdaExecutionPath,
							)),
						},
					},
				},
			),
			Role: iamRole,
		},
	)

	iamRole.AddManagedPolicy(iam.ManagedPolicy_FromAwsManagedPolicyName(
		jsii.String("service-role/AWSLambdaBasicExecutionRole"),
	))
	iamRole.AddManagedPolicy(iam.ManagedPolicy_FromAwsManagedPolicyName(
		jsii.String("service-role/AWSLambdaVPCAccessExecutionRole"),
	))

	return LambdaConstruct{
		LambdaFunction: graphQLAPIExecutorFunction,
		Scope:          scope,
	}
}
