package construct

import (
	"fmt"
	"path/filepath"

	iam "github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	lambda "github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewLambdaConstruct(
	scope constructs.Construct,
	name string,
) constructs.Construct {
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

	lambda.NewFunction(
		scope,
		jsii.String(fmt.Sprintf("%sLambdaConstruct", name)),
		&lambda.FunctionProps{
			Runtime: lambda.Runtime_GO_1_X(),
			Handler: jsii.String("main"),
			Code: lambda.AssetCode_FromAsset(
				jsii.String(filepath.Join(".", "..", "..", "cmd", "lambda")),
				nil,
				//&s3assets.AssetOptions{
				//	Bundling: &awscdk.BundlingOptions{
				//		Image: lambda.Runtime_GO_1_X().BundlingImage(),
				//		User:  jsii.String("root"),
				//		Command: &[]*string{
				//			jsii.String("bash"),
				//			jsii.String("-c"),
				//			jsii.String("go version && go build -o /asset-output/main"),
				//		},
				//	},
				//},
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

	return scope
}
