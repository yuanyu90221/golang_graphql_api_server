package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
	_ "github.com/joho/godotenv/autoload"
	"github.com/yuanyu90221/golang_graphql_api_server/gql"
	"github.com/yuanyu90221/golang_graphql_api_server/postgres"
	"github.com/yuanyu90221/golang_graphql_api_server/server"
)

func main() {
	// Initialze api and return a pointer to router for http.ListenAndServe
	router, db := initializeAPI()
	defer db.Close()
	//
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router))
}

func initializeAPI() (*chi.Mux, *postgres.Db) {
	// Create a new router
	router := chi.NewRouter()
	// current port
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	fmt.Println("db port", port)
	if err != nil {
		log.Fatal(err)
	}
	// Create a new connection to our pg database
	db, err := postgres.New(
		postgres.ConnString(os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_NAME")),
	)
	if err != nil {
		log.Fatal(err)
	}
	// Create root query for graphql
	rootQuery := gql.NewRoot(db)
	// Create a new graphql schema, passing in the root query
	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query},
	)
	if err != nil {
		fmt.Println("Error creating schema: ", err)
	}

	// Create a server struct that holds a pointer to our database as well as the address of graphql schema
	s := server.Server{
		GqlSchema: &sc,
	}
	// Add some middleware to router
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // set content-type headers as application/json
		middleware.Logger, // log api request calls
		// , // compress results, mostly gzipping assets and json
		middleware.StripSlashes, // match paths with a trailing slash, strip it
		middleware.Recoverer,
	)
	// Create the graphql route with a Server method to handle it
	router.Post("/graphql", s.GraphQL())
	return router, db
}
