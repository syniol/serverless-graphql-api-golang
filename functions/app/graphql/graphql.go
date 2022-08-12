package graphql

import (
	"context"
	"encoding/json"
	"github.com/graphql-go/graphql"
)

type PostData struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

func NewGraphQLExecution(
	ctx context.Context,
	schema graphql.Schema,
	body string,
) (*graphql.Result, error) {
	var p PostData

	if err := json.Unmarshal([]byte(body), &p); err != nil {
		return nil, err
	}

	result := graphql.Do(graphql.Params{
		Context:        ctx,
		Schema:         schema,
		RequestString:  p.Query,
		VariableValues: p.Variables,
	})

	if len(result.Errors) > 0 {
		return nil, result.Errors[0]
	}

	return result, nil
}
