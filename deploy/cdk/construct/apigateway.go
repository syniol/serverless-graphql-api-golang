package construct

import (
	"github.com/aws/aws-cdk-go/awscdk"
	apigateway "github.com/aws/aws-cdk-go/awscdk/awsapigateway"
	"github.com/aws/jsii-runtime-go"
)

func NewAPIGateWayConstruct(
	scope awscdk.Stack,
	name string,
	lambdaConstruct LambdaConstruct,
) awscdk.Stack {
	api := apigateway.NewLambdaRestApi(scope, jsii.String(name), &apigateway.LambdaRestApiProps{
		RestApiName: jsii.String("GraphQLAPIGateWayTrigger"),
		Handler:     lambdaConstruct.LambdaFunction,
		Proxy:       jsii.Bool(false),
	})

	api.Root().AddMethod(jsii.String("ANY"), nil, nil)

	resource := api.Root().AddResource(jsii.String("graphql"), nil)
	resource.AddMethod(jsii.String("POST"), nil, nil)

	return scope
}
