package construct

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	dynamodb "github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/jsii-runtime-go"
)

type DynamoDBConstruct struct {
	Table dynamodb.Table
	Scope awscdk.Stack
}

func NewDynamoDBConstruct(scope awscdk.Stack, name string) *DynamoDBConstruct {
	return &DynamoDBConstruct{
		Table: dynamodb.NewTable(
			scope,
			jsii.String(name),
			&dynamodb.TableProps{
				PartitionKey: &dynamodb.Attribute{
					Name: jsii.String("graphql"),
					Type: dynamodb.AttributeType_STRING,
				},
				SortKey:                    nil,
				BillingMode:                dynamodb.BillingMode_PAY_PER_REQUEST,
				ContributorInsightsEnabled: nil,
				Encryption:                 dynamodb.TableEncryption_AWS_MANAGED,
				ReadCapacity:               nil,
				RemovalPolicy:              awscdk.RemovalPolicy_DESTROY,
				ReplicationRegions: &[]*string{
					jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
				},
				ReplicationTimeout: nil,
				TableName:          jsii.String("graphql"),
			},
		),
		Scope: scope,
	}
}
