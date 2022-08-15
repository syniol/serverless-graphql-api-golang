package construct

import (
	"fmt"
	"path/filepath"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	apigateway "github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	iam "github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	lambda "github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	s3assets "github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/jsii-runtime-go"
)

const AWSLambdaExecutionPath = "/asset-output/"

type LambdaConstruct struct {
	LambdaFunction  lambda.Function
	Scope           awscdk.Stack
	APIGatewayProps *apigateway.LambdaRestApiProps
}

func newLambdaProps(props *lambda.FunctionProps) *lambda.FunctionProps {
	if props != nil {
		return props
	}

	return &lambda.FunctionProps{
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
	}
}

func NewLambdaConstruct(
	scope awscdk.Stack,
	name string,
	props *lambda.FunctionProps,
	dynamoDB *DynamoDBConstruct,
	neptune *NeptuneConstruct,
) *LambdaConstruct {
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

	lambdaProps := newLambdaProps(props)
	lambdaProps.Role = iamRole

	graphQLAPIExecutorFunction := lambda.NewFunction(
		scope,
		jsii.String(fmt.Sprintf("%sLambdaConstruct", name)),
		lambdaProps,
	)

	iamRole.AddManagedPolicy(iam.ManagedPolicy_FromAwsManagedPolicyName(
		jsii.String("service-role/AWSLambdaBasicExecutionRole"),
	))
	iamRole.AddManagedPolicy(iam.ManagedPolicy_FromAwsManagedPolicyName(
		jsii.String("service-role/AWSLambdaVPCAccessExecutionRole"),
	))

	if dynamoDB != nil {
		iamRole.AddToPolicy(iam.NewPolicyStatement(
			&iam.PolicyStatementProps{
				Effect: iam.Effect_ALLOW,
				Actions: &[]*string{
					jsii.String("dynamodb:*"),
				},
				Resources: &[]*string{
					dynamoDB.Table.TableArn(),
				},
			},
		))
	}

	if neptune != nil {
		accountNumber := *scope.Account()
		region := *scope.Region()

		iamRole.AddToPolicy(iam.NewPolicyStatement(
			&iam.PolicyStatementProps{
				Effect: iam.Effect_ALLOW,
				Actions: &[]*string{
					jsii.String("neptune-db:*"),
				},
				Resources: &[]*string{
					jsii.String(fmt.Sprintf("arn:aws:rds:%s:%s:db:%s", region, accountNumber, *neptune.DatabaseInstance.DbInstanceIdentifier())),
					jsii.String(fmt.Sprintf("arn:aws:rds:%s:%s:cluster:%s", region, accountNumber, *neptune.DatabaseInstance.DbClusterIdentifier())),
				},
			},
		))
	}

	return &LambdaConstruct{
		LambdaFunction: graphQLAPIExecutorFunction,
		Scope:          scope,
	}
}
