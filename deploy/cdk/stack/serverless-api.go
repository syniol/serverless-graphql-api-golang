package stack

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"

	"cdk/construct"
)

func NewServerlessAPIStack(
	scope awscdk.App,
	id string,
	props *awscdk.StackProps,
) awscdk.Stack {
	stack := awscdk.NewStack(scope, &id, props)

	construct.NewAPIGateWayConstruct(
		stack,
		"GraphQLExecutionAPI",
		construct.NewLambdaConstruct(
			stack,
			"GraphQLExecution",
			nil,
			//construct.NewDynamoDBConstruct(stack, "GraphQLNoSQLDatabase"),
			nil,
			construct.NewNeptuneConstruct(stack, "GraphQLNoSQLDatabase", nil, nil),
		),
		nil,
	)

	return stack
}
