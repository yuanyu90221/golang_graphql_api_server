package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/yuanyu90221/golang_graphql_api_server/postgres"
)

// Root holds a pointer to a graphql object
type Root struct {
	Query *graphql.Object
}

// NewRoot returns base query type. This is where we add all the base queries
func NewRoot(db *postgres.Db) *Root {
	// Create a resolver holding our database. Resolver can be found in resolvers.go
	resolver := Resolver{db: db}

	// Create a new Root that describes base query set up
	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"dusers": &graphql.Field{
						// Slice of User type which can be found in types.go
						Type: graphql.NewList(Duser),
						Args: graphql.FieldConfigArgument{
							"name": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: resolver.DuserResolver,
					},
				},
			},
		),
	}
	return &root
}
