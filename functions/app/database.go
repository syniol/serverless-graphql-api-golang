package app

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	//"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/neptune"
)

type Database struct {
	*neptune.Neptune
}

func NewDatabase() *Database {
	sessionDatabase := session.Must(session.NewSession())
	dbClient := neptune.New(
		sessionDatabase,
		aws.NewConfig().WithRegion(
			os.Getenv("CDK_DEFAULT_REGION"),
		),
	)

	//_ = dbClient.NewRequest(
	//	&request.Operation{},
	//	"",
	//	"",
	//).Send()

	return &Database{
		dbClient,
	}
}
