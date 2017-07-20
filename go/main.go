package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	field "github.com/WindomZ/graphql-examples/go/field"
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

func HandleQuery(w http.ResponseWriter, r *http.Request, schema graphql.Schema) {
	result, err := executeQuery(r.URL.Query().Get("query"), schema)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(result)
}

func main() {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graphql.NewObject(
				graphql.ObjectConfig{
					Name: "Query",
					Fields: graphql.Fields{
						"hello": field.HelloField,
						"user":  field.UserField,
					},
				},
			),
		},
	)
	if err != nil {
		panic(fmt.Sprintf("failed to create new schema, error: %v", err))
	}

	http.HandleFunc("/example", func(w http.ResponseWriter, r *http.Request) {
		HandleQuery(w, r, schema)
	})

	fmt.Println("Now server is running on port 8080")
	panic(http.ListenAndServe(":8080", nil))
}
