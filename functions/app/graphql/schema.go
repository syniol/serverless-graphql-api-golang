package graphql

import (
	"github.com/graphql-go/graphql/examples/todo/schema"
)

//var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{})

var Schema = schema.TodoSchema
