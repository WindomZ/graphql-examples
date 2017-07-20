package graphql_examples

import "github.com/graphql-go/graphql"

var HelloField *graphql.Field

func init() {
	HelloField = &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "world", nil
		},
	}
}
