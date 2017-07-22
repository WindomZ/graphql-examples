package graphqlexamples

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

// HelloField define a hello query field
var HelloField *graphql.Field

// ByeField define a bye query field
var ByeField *graphql.Field

func init() {
	HelloField = &graphql.Field{
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"message": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			message, ok := p.Args["message"].(string)
			if !ok {
				return "Can not read 'message'!", nil
			}
			return fmt.Sprintf("Hello %s!", message), nil
		},
	}
	ByeField = &graphql.Field{
		Type: graphql.String,
	}
}
