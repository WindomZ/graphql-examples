package graphqlexamples

import "github.com/graphql-go/graphql"

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var data map[string]User = map[string]User{
	"1": {
		ID:   "1",
		Name: "+86-13888888888",
	},
	"id": {
		ID:   "id",
		Name: "Name",
	},
	"编号": {
		ID:   "编号",
		Name: "名字",
	},
}

var UserField *graphql.Field

func init() {
	/*
		Create User object type with fields "id" and "name" by using GraphQLObjectTypeConfig:
		    - Name: name of object type
			- Fields: a map of fields by using GraphQLFields
		Setup type of field use GraphQLFieldConfig
	*/
	var userType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "User",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.String,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	/*
	   Setup type of field use GraphQLFieldConfig to define:
	       - Type: type of field
	       - Args: arguments to query with current field
	       - Resolve: function to query data using params from [Args] and return value with current type
	*/
	UserField = &graphql.Field{
		Type: userType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, ok := p.Args["id"].(string)
			if ok {
				return data[id], nil
			}
			return nil, nil
		},
	}
}
