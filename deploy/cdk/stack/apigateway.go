package stack

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"

	"cdk/construct"
)

func NewAPIGateWayStack(
	scope constructs.Construct,
	id string,
	props *awscdk.StackProps,
) awscdk.Stack {
	stack := awscdk.NewStack(scope, &id, props)

	construct.NewLambdaConstruct(stack, "GraphQLExecution")

	return stack
}
