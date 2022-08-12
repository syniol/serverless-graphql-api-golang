package graphql

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

type PostData struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

func PostRequest(w http.ResponseWriter, req *http.Request, schema graphql.Schema) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		errorMsg, _ := json.Marshal(`{ "error": "only post method is accepted"}`)
		_, _ = w.Write(errorMsg)

		return
	}

	var p PostData
	if err := json.NewDecoder(req.Body).Decode(&p); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		errorMsg, _ := json.Marshal(fmt.Sprintf(`{ "error": "%s" }`, err))
		_, _ = w.Write(errorMsg)
		return
	}

	result := graphql.Do(graphql.Params{
		Context:        req.Context(),
		Schema:         schema,
		RequestString:  p.Query,
		VariableValues: p.Variables,
	})

	if err := json.NewEncoder(w).Encode(result); err != nil {
		errorMsg, _ := json.Marshal(fmt.Sprintf(`{ "error": "%s" }`, err))
		_, _ = w.Write(errorMsg)

		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
