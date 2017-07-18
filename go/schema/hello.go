package graphql_examples

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

var HelloSchema graphql.Schema

func init() {
	// Fields
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}

	// Schema
	var err error
	HelloSchema, err = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   "Query",
					Fields: fields,
				},
			),
		},
	)
	if err != nil {
		panic(fmt.Sprintf("failed to create new schema, error: %v", err))
	}
}
