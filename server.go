package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/japiirainen/go-oluet-api/db"
	"github.com/japiirainen/go-oluet-api/gql/generated"
	"github.com/japiirainen/go-oluet-api/gql/resolvers"
	log "github.com/sirupsen/logrus"
)

const defaultPort = "5000"

func main() {
	log.SetFormatter(&log.TextFormatter{})
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	c := db.Connect()
	defer c.CloseConnection()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{
		DB: c,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Infof("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
