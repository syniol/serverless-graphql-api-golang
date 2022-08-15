package construct

import (
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	stepfunctions "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	tasks "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctionstasks"
	"github.com/aws/jsii-runtime-go"
)

type StepFunctionsConstruct struct {
	Scope           awscdk.Stack
	StateMachine    stepfunctions.StateMachine
	APIGatewayProps *awsapigateway.StepFunctionsRestApiProps
}

func NewStepFunctionsConstruct(
	scope awscdk.Stack,
	name string,
	lambdas ...awslambda.IFunction,
) *StepFunctionsConstruct {
	smDefinition := tasks.NewLambdaInvoke(
		scope,
		jsii.String(fmt.Sprintf("%slientIdentifier", name)),
		&tasks.LambdaInvokeProps{LambdaFunction: lambdas[0]},
	).Next(
		stepfunctions.NewChoice(scope, jsii.String(fmt.Sprintf("%sSFChoice", name)), nil).
			When(stepfunctions.Condition_StringEquals(jsii.String("$.type"), jsii.String("mutation")), tasks.NewLambdaInvoke(
				scope,
				jsii.String(fmt.Sprintf("%slientIdentifier", name)),
				&tasks.LambdaInvokeProps{LambdaFunction: lambdas[1]},
			)).
			Otherwise(tasks.NewLambdaInvoke(
				scope,
				jsii.String(fmt.Sprintf("%slientIdentifier", name)),
				&tasks.LambdaInvokeProps{LambdaFunction: lambdas[3]},
			)),
	)

	return &StepFunctionsConstruct{
		Scope: scope,
		StateMachine: stepfunctions.NewStateMachine(scope, jsii.String(name), &stepfunctions.StateMachineProps{
			Definition:       smDefinition,
			Logs:             nil,
			Role:             nil,
			StateMachineName: jsii.String(fmt.Sprintf("%sStateMachine", name)),
			StateMachineType: stepfunctions.StateMachineType_STANDARD,
			Timeout:          awscdk.Duration_Minutes(jsii.Number(5)),
			TracingEnabled:   jsii.Bool(false),
		}),
	}
}
