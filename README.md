# golang_graphql_api_server
## introduction

  This is a repository for use graphql api with golang on postgresql database project

## reference document
  [building-an-api-with-graphql-and-go](https://medium.com/@bradford_hamilton/building-an-api-with-graphql-and-go-9350df5c9356)
## reference package
  [github.com/go-chi/chi](github.com/go-chi/chi)
  [github.com/go-chi/render](github.com/go-chi/render)
  [github.com/graphql-go/graphql](github.com/graphql-go/graphql)
## key point for restful to graphql
  use resolver to handle the query to mapping handler

1 resolver part
```golang
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
```
2 query part
```golang
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
```
3 type part
```golang
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

```