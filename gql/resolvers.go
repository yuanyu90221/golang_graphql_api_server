package gql

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/yuanyu90221/golang_graphql_api_server/postgres"
)

// Resolver struct holds a connection to database
type Resolver struct {
	db *postgres.Db
}

// DuserResolver resolves duser query through a db call to GetUserByName
func (r *Resolver) DuserResolver(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string
	name, ok := p.Args["name"].(string)
	fmt.Println("name:", name)
	if ok {
		dusers := r.db.GetUsersByName(name)
		return dusers, nil
	}
	return nil, nil
}
