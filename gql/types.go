package gql

import (
	"github.com/graphql-go/graphql"
)

// Duser describes a graphql object containing a Duser
var Duser = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Duser",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"age": &graphql.Field{
				Type: graphql.Int,
			},
			"profession": &graphql.Field{
				Type: graphql.String,
			},
			"friendly": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)
