package graphqlexamples

import (
	"errors"

	"github.com/graphql-go/graphql"
)

// SumField define a sum mutation field
var SumField *graphql.Field

func init() {
	SumField = &graphql.Field{
		Type: graphql.Int,
		Args: graphql.FieldConfigArgument{
			"x": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"y": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			x, ok := p.Args["x"].(int)
			if !ok {
				return "Can not read 'x'!", nil
			}
			y, ok := p.Args["y"].(int)
			if !ok {
				return nil, errors.New("Can not read 'y'!")
			}
			return x + y, nil
		},
	}
}
