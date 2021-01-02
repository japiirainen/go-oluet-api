package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/japiirainen/go-oluet-api/db"
	"github.com/japiirainen/go-oluet-api/graph"
	"github.com/japiirainen/go-oluet-api/graph/generated"
)

const defaultPort = "5000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	c := db.Connect()
	defer c.CloseConnection()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: c,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
