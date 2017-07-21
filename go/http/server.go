package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	schema "github.com/WindomZ/graphql-examples/go/graphql"
	"github.com/graphql-go/graphql"
)

func executeQuery(query string, schema graphql.Schema) (*graphql.Result, error) {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
		RootObject: map[string]interface{}{
			"hello": func() interface{} { return "Should not be displayed!" },
			"bye":   func() interface{} { return "Goodbye!" },
		},
	})
	if len(result.Errors) > 0 {
		return nil, errors.New(fmt.Sprintf("wrong result, unexpected errors: %v", result.Errors))
	}
	return result, nil
}

func handleQuery(w http.ResponseWriter, r *http.Request, schema graphql.Schema) {
	var query string
	switch r.Method {
	case http.MethodPost:
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		obj := &struct {
			Query string `json:"query"`
		}{}
		if err = json.Unmarshal(data, obj); err != nil {
			panic(err)
		}
		query = obj.Query
	default:
		query = r.URL.Query().Get("query")
	}

	result, err := executeQuery(query, schema)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		handleQuery(w, r, schema.ExampleSchema)
	})

	fmt.Println("Now server is running on port 8080")
	panic(http.ListenAndServe(":8080", nil))
}
