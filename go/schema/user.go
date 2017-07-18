package graphql_examples

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var data map[string]User = map[string]User{
	"1": {
		ID:   "1",
		Name: "Name-1",
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

var UserSchema graphql.Schema

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

	// Fields
	fields := graphql.Fields{
		"user": &graphql.Field{
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
		},
	}

	/*
	   Create Query object type with fields "user" has type [userType] by using GraphQLObjectTypeConfig:
	       - Name: name of object type
	       - Fields: a map of fields by using GraphQLFields
	   Setup type of field use GraphQLFieldConfig to define:
	       - Type: type of field
	       - Args: arguments to query with current field
	       - Resolve: function to query data using params from [Args] and return value with current type
	*/
	var queryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name:   "Query",
			Fields: fields,
		})

	// Schema
	var err error
	UserSchema, err = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
		},
	)
	if err != nil {
		panic(fmt.Sprintf("failed to create new schema, error: %v", err))
	}
}
