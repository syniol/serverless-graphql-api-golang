package app

import (
	"fmt"
	"homadrone/app/graphql"
	"net/http"
	"os"

	graphqlLib "github.com/graphql-go/graphql"
)

var PortNumber = func() string {
	if len(os.Getenv("GRAPHQL_HTTP_PORT")) > 0 {
		return os.Getenv("GRAPHQL_HTTP_PORT")
	}

	return "8080"
}

func NewHttpServer(schema graphqlLib.Schema) {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, req *http.Request) {
		graphql.PostRequest(w, req, schema)
	})

	fmt.Println(fmt.Sprintf(
		"🚀 GraphQL Server is running on port: %s", PortNumber(),
	))

	_ = http.ListenAndServe(
		fmt.Sprintf(":%s", PortNumber()),
		nil,
	)
}
