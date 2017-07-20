package graphql_examples

import (
	"fmt"

	field "github.com/WindomZ/graphql-examples/go/graphql/field"
	"github.com/graphql-go/graphql"
)

var ExampleSchema graphql.Schema

func init() {
	var err error
	ExampleSchema, err = graphql.NewSchema(
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
}
