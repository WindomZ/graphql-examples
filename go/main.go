package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/WindomZ/graphql-examples/go/schema"
	"github.com/graphql-go/graphql"
)

func executeQuery(query string, schema graphql.Schema) (*graphql.Result, error) {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		return nil, errors.New(fmt.Sprintf("wrong result, unexpected errors: %v", result.Errors))
	}
	return result, nil
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		result, err := executeQuery(r.URL.Query().Get("query"), graphql_examples.HelloSchema)
		if err != nil {
			panic(err)
		}

		json.NewEncoder(w).Encode(result)
	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		result, err := executeQuery(r.URL.Query().Get("query"), graphql_examples.UserSchema)
		if err != nil {
			panic(err)
		}

		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Now server is running on port 8080")
	panic(http.ListenAndServe(":8080", nil))
}
