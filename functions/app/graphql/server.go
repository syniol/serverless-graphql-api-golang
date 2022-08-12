package graphql

import (
	"fmt"
	"net/http"
	"os"

	"github.com/graphql-go/graphql"
)

var PortNumber = func() string {
	if len(os.Getenv("GRAPHQL_HTTP_PORT")) > 0 {
		return os.Getenv("GRAPHQL_HTTP_PORT")
	}

	return "8080"
}

func NewHttpServer(schema graphql.Schema) {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, req *http.Request) {
		PostRequest(w, req, schema)
	})

	fmt.Println(fmt.Sprintf(
		"🚀 GraphQL Server is running on port: %s", PortNumber(),
	))

	_ = http.ListenAndServe(
		fmt.Sprintf(":%s", PortNumber()),
		nil,
	)
}
